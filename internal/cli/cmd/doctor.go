// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package cmd

import (
	"context"
	"fmt"
	"io"
	"time"

	dockertypes "github.com/docker/docker/api/types"
	"github.com/kr/text"
	"github.com/moby/buildkit/client/llb"
	"github.com/spf13/cobra"
	"namespacelabs.dev/foundation/build"
	"namespacelabs.dev/foundation/build/buildkit"
	"namespacelabs.dev/foundation/internal/artifacts/oci"
	"namespacelabs.dev/foundation/internal/cli/fncobra"
	"namespacelabs.dev/foundation/internal/compute"
	"namespacelabs.dev/foundation/internal/console"
	"namespacelabs.dev/foundation/internal/console/colors"
	"namespacelabs.dev/foundation/internal/console/common"
	"namespacelabs.dev/foundation/internal/dependencies/pins"
	"namespacelabs.dev/foundation/internal/fnapi"
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/internal/parsing"
	"namespacelabs.dev/foundation/internal/parsing/module"
	"namespacelabs.dev/foundation/runtime"
	"namespacelabs.dev/foundation/runtime/docker"
	"namespacelabs.dev/foundation/runtime/kubernetes"
	"namespacelabs.dev/foundation/runtime/rtypes"
	"namespacelabs.dev/foundation/std/cfg"
	"namespacelabs.dev/foundation/std/tasks"
	"namespacelabs.dev/go-ids"
)

const (
	checkTimeLimit = 30 * time.Second
)

func NewDoctorCmd() *cobra.Command {
	return fncobra.Cmd(&cobra.Command{
		Use:   "doctor",
		Short: "Collect diagnostic information about the system.",
		Args:  cobra.NoArgs,
	}).Do(func(ctx context.Context) error {
		//

		versionI := runDiagnostic(ctx, "doctor.version-info", func(ctx context.Context) (*VersionInfo, error) {
			return CollectVersionInfo()
		})
		printDiagnostic(ctx, "Namespace version", versionI, FormatVersionInfo)

		//

		dockerI := runDiagnostic(ctx, "doctor.docker-info", func(ctx context.Context) (dockerInfo, error) {
			client, err := docker.NewClient()
			if err != nil {
				return dockerInfo{}, err
			}
			version, err := client.ServerVersion(ctx)
			if err != nil {
				return dockerInfo{}, err
			}
			info, err := client.Info(ctx)
			if err != nil {
				return dockerInfo{}, err
			}
			return dockerInfo{version, info}, nil
		})
		printDiagnostic(ctx, "Docker", dockerI, func(w io.Writer, info dockerInfo) {
			fmt.Fprintf(w, "version=%s (commit=%s) for %s-%s\n", info.Version.Version, info.Version.GitCommit, info.Version.Os, info.Version.Arch)
			fmt.Fprintf(w, "api version=%s (min=%s)\n", info.Version.APIVersion, info.Version.MinAPIVersion)
			fmt.Fprintf(w, "kernel version=%s\n", info.Version.KernelVersion)
			fmt.Fprintf(w, "ncpu=%d mem_total=%d\n", info.Info.NCPU, info.Info.MemTotal)
			fmt.Fprintf(w, "containers total=%d running=%d paused=%d stopped=%d images=%d\n", info.Info.Containers, info.Info.ContainersRunning,
				info.Info.ContainersPaused, info.Info.ContainersStopped, info.Info.Images)
			fmt.Fprintf(w, "driver=%s logging_driver=%s cgroup_driver=%s cgroup_version=%s\n", info.Info.Driver, info.Info.LoggingDriver,
				info.Info.CgroupDriver, info.Info.CgroupVersion)
			fmt.Fprintf(w, "containerd expected=%s present=%s\n", info.Info.ContainerdCommit.Expected, info.Info.ContainerdCommit.ID)
			fmt.Fprintf(w, "runc expected=%s present=%s\n", info.Info.RuncCommit.Expected, info.Info.RuncCommit.ID)
			fmt.Fprintf(w, "init expected=%s present=%s\n", info.Info.InitCommit.Expected, info.Info.InitCommit.ID)
			if len(info.Info.Warnings) > 0 {
				fmt.Fprintln(w, "Warnings:")
				for _, wn := range info.Info.Warnings {
					fmt.Fprintf(w, "  %s", wn)
				}
			}
		})

		//

		loginI := runDiagnostic(ctx, "doctor.userauth", func(ctx context.Context) (*fnapi.UserAuth, error) {
			return fnapi.LoadUser()
		})
		printDiagnostic(ctx, "Authenticated User", loginI, func(w io.Writer, info *fnapi.UserAuth) {
			fmt.Fprintf(w, "user=%s\n", info.Username)
		})

		//
		type dockerResults struct {
			ImageLatency time.Duration
			RunLatency   time.Duration
		}

		dockerRunI := runDiagnostic(ctx, "doctor.docker-run", func(ctx context.Context) (dockerResults, error) {
			t := time.Now()
			image, err := compute.GetValue(ctx, oci.ResolveImage(pins.Image("hello-world"), docker.HostPlatform()).Image())
			if err != nil {
				return dockerResults{}, err
			}
			var r dockerResults
			r.ImageLatency = time.Since(t)
			t = time.Now()
			err = docker.Impl().Run(ctx, rtypes.RunToolOpts{
				RunBinaryOpts: rtypes.RunBinaryOpts{
					Image: image,
				},
			})
			r.RunLatency = time.Since(t)
			return r, err
		})
		printDiagnostic(ctx, "Docker Run Check", dockerRunI, func(w io.Writer, info dockerResults) {
			fmt.Fprintf(w, "image_pull_latency=%v docker_run_latency=%v\n", info.ImageLatency, info.RunLatency)
		})

		//

		workspaceI := runDiagnostic(ctx, "doctor.workspace", func(ctx context.Context) (*parsing.Root, error) {
			return module.FindRoot(ctx, ".")
		})
		printDiagnostic(ctx, "Workspace", workspaceI, func(w io.Writer, ws *parsing.Root) {
			fmt.Fprintf(w, "module=%s\n", ws.ModuleName())
		})

		//

		type buildkitResults struct {
			BuildLatency time.Duration
			Image        oci.Image
		}

		buildkitI := errorOr[buildkitResults]{err: fnerrors.New("no workspace")}
		if workspaceI.err == nil {
			buildkitI = runDiagnostic(ctx, "doctor.build", func(ctx context.Context) (buildkitResults, error) {
				env, err := cfg.LoadContext(workspaceI.v, "dev")
				if err != nil {
					return buildkitResults{}, err
				}
				hostPlatform := buildkit.HostPlatform()
				state := llb.Scratch().File(llb.Mkfile("/hello", 0644, []byte("cachehit")))
				conf := build.NewBuildTarget(&hostPlatform).WithSourceLabel("doctor.hello-world")
				var r buildkitResults
				t := time.Now()
				imageC, err := buildkit.BuildImage(ctx, env, conf, state)
				if err != nil {
					return buildkitResults{}, err
				}
				image, err := compute.Get(ctx, imageC)
				if err != nil {
					return buildkitResults{}, err
				}
				r.BuildLatency = time.Since(t)
				r.Image = image.Value
				return r, nil
			})
		}
		printDiagnostic(ctx, "Buildkit", buildkitI, func(w io.Writer, r buildkitResults) {
			digest, _ := r.Image.Digest()
			fmt.Fprintf(w, "success build_latency=%v image_digest=%v\n", r.BuildLatency, digest)
		})

		//

		type kubeResults struct {
			ConnectLatency time.Duration
			RunLatency     time.Duration
		}
		kubernetesI := errorOr[kubeResults]{err: fnerrors.New("no workspace")}
		if workspaceI.err == nil {
			kubernetesI = runDiagnostic(ctx, "doctor.kube", func(ctx context.Context) (kubeResults, error) {
				var r kubeResults
				env, err := cfg.LoadContext(workspaceI.v, "dev")
				if err != nil {
					return r, err
				}
				t := time.Now()
				k, err := kubernetes.ConnectToCluster(ctx, env.Configuration())
				if err != nil {
					return r, err
				}
				r.ConnectLatency = time.Since(t)
				helloID, err := oci.ParseImageID(pins.Image("hello-world"))
				if err != nil {
					return r, err
				}
				t = time.Now()
				err = k.RunAttachedOpts(ctx, "default", "doctor-"+ids.NewRandomBase32ID(8),
					runtime.ContainerRunOpts{Image: helloID}, runtime.TerminalIO{}, func() {})
				r.RunLatency = time.Since(t)
				return r, err
			})
		}

		printDiagnostic(ctx, "Kubernetes Run Check", kubernetesI, func(w io.Writer, info kubeResults) {
			fmt.Fprintf(w, "connect_latency=%v run_latency=%v\n", info.ConnectLatency, info.RunLatency)
		})
		return nil
	})
}

func runDiagnostic[V any](ctx context.Context, title string, f func(ctx context.Context) (V, error)) errorOr[V] {
	res := errorOr[V]{}
	// We have to run in a separate orchestrator so that failures in one diagnostic
	// do not prevent other diagnostics from proceeding.
	res.err = compute.Do(ctx, func(ctx context.Context) error {
		v, err := tasks.Return(ctx, tasks.Action(title), func(ctx context.Context) (V, error) {
			timedCtx, cancel := context.WithTimeout(ctx, checkTimeLimit)
			defer cancel()
			return f(timedCtx)
		})
		res.v = v
		return err
	})
	return res
}

func printDiagnostic[V any](ctx context.Context, title string, res errorOr[V], print func(io.Writer, V)) {
	style := colors.Ctx(ctx)

	w := console.TypedOutput(ctx, title, common.CatOutputUs)

	fmt.Fprintln(w, style.Header.Apply(fmt.Sprintf("* %s", title)))
	x := text.NewIndentWriter(w, []byte("  "))
	if res.err != nil {
		// Not using format.Format since it's too verbose.
		fmt.Fprintln(x, style.ErrorHeader.Apply("Failed:"), res.err)
	} else {
		print(x, res.v)
	}
}

type dockerInfo struct {
	dockertypes.Version
	dockertypes.Info
}

type errorOr[V any] struct {
	v   V
	err error
}