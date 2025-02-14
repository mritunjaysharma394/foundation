// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package source

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"namespacelabs.dev/foundation/internal/cli/fncobra"
	"namespacelabs.dev/foundation/internal/console"
	"namespacelabs.dev/foundation/internal/runtime/rtypes"
	"namespacelabs.dev/foundation/internal/sdk/deno"
	"namespacelabs.dev/foundation/internal/sdk/host"
	"namespacelabs.dev/foundation/std/cfg"
)

func newDenoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deno",
		Short: "Run Namespace-configured deno.",
	}

	return fncobra.CmdWithEnv(cmd, func(ctx context.Context, env cfg.Context, args []string) error {
		d, err := deno.EnsureSDK(ctx, host.HostPlatform())
		if err != nil {
			return err
		}

		if err := d.CacheImports(ctx, env.Workspace().LoadedFrom().AbsPath); err != nil {
			return err
		}

		x := console.EnterInputMode(ctx)
		defer x()

		return d.Run(ctx, env.Workspace().LoadedFrom().AbsPath, rtypes.IO{Stdin: os.Stdin, Stdout: os.Stdout, Stderr: os.Stderr}, "repl", "--cached-only")
	})
}
