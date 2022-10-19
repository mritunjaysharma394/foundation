// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package kubeops

import (
	"context"
	"encoding/json"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	applycorev1 "k8s.io/client-go/applyconfigurations/core/v1"
	"namespacelabs.dev/foundation/framework/rpcerrors/multierr"
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/internal/planning/deploy"
	"namespacelabs.dev/foundation/runtime/kubernetes/kubedef"
	"namespacelabs.dev/foundation/schema"
	"namespacelabs.dev/foundation/std/execution"
	"namespacelabs.dev/foundation/std/tasks"
)

const runtimeConfigVersion = 0

func registerEnsureRuntimeConfig() {
	execution.RegisterFuncs(execution.Funcs[*kubedef.OpEnsureRuntimeConfig]{
		Handle: func(ctx context.Context, inv *schema.SerializedInvocation, ensure *kubedef.OpEnsureRuntimeConfig) (*execution.HandleResult, error) {
			action := tasks.Action("kubernetes.ensure-runtime-config").
				Scope(schema.PackageName(ensure.Deployable.PackageName)).
				Arg("deployable", ensure.Deployable.PackageName).
				HumanReadablef(inv.Description)

			return tasks.Return(ctx, action, func(ctx context.Context) (*execution.HandleResult, error) {
				data := map[string]string{}

				output := &kubedef.EnsureRuntimeConfigOutput{}

				if ensure.RuntimeConfig != nil {
					serializedConfig, err := json.Marshal(ensure.RuntimeConfig)
					if err != nil {
						return nil, fnerrors.InternalError("failed to serialize runtime configuration: %w", err)
					}
					data["runtime.json"] = string(serializedConfig)
					output.SerializedRuntimeJson = string(serializedConfig)
				}

				if ensure.BuildVcs != nil {
					serializedConfig, err := json.Marshal(ensure.BuildVcs)
					if err != nil {
						return nil, fnerrors.InternalError("failed to serialize runtime configuration: %w", err)
					}
					data["buildvcs.json"] = string(serializedConfig)
				}

				resourceData, err := deploy.BuildResourceMap(ctx, ensure.Dependency)
				if err != nil {
					return nil, err
				}

				if len(ensure.InjectResource) > 0 {
					if resourceData == nil {
						resourceData = map[string]deploy.RawJSONObject{}
					}

					var errs []error
					for _, injected := range ensure.InjectResource {
						var m deploy.RawJSONObject
						if err := json.Unmarshal(injected.SerializedJson, &m); err != nil {
							errs = append(errs, err)
						} else {
							resourceData[injected.GetResourceRef().Canonical()] = m
						}
					}

					if err := multierr.New(errs...); err != nil {
						return nil, fnerrors.InternalError("failed to handle injected resources: %w", err)
					}
				}

				if len(resourceData) > 0 {
					serializedConfig, err := json.Marshal(resourceData)
					if err != nil {
						return nil, fnerrors.InternalError("failed to serialize resource configuration: %w", err)
					}
					data["resources.json"] = string(serializedConfig)

					output.SerializedResourceJson = string(serializedConfig)
				}

				if len(data) > 0 && ensure.PersistConfiguration {
					keys := maps.Keys(data)
					slices.Sort(keys)
					hashInput := []any{runtimeConfigVersion}
					for _, key := range keys {
						hashInput = append(hashInput, data[key])
					}

					configDigest, err := schema.DigestOf(hashInput...)
					if err != nil {
						return nil, fnerrors.InternalError("failed to digest runtime configuration: %w", err)
					}

					deploymentId := kubedef.MakeDeploymentId(ensure.Deployable)
					configId := kubedef.MakeVolumeName(deploymentId, "rtconfig-"+configDigest.Hex[:8])

					cluster, err := kubedef.InjectedKubeClusterNamespace(ctx)
					if err != nil {
						return nil, err
					}

					annotations := kubedef.MakeAnnotations(cluster.KubeConfig().Environment, schema.PackageName(ensure.Deployable.PackageName))
					labels := kubedef.MakeLabels(cluster.KubeConfig().Environment, ensure.Deployable)

					if _, err := cluster.Cluster().(kubedef.KubeCluster).PreparedClient().Clientset.CoreV1().
						ConfigMaps(cluster.KubeConfig().Namespace).
						Apply(ctx,
							applycorev1.ConfigMap(configId, cluster.KubeConfig().Namespace).
								WithAnnotations(annotations).
								WithLabels(labels).
								WithLabels(map[string]string{
									kubedef.K8sKind: kubedef.K8sRuntimeConfigKind,
								}).
								WithImmutable(true).
								WithData(data), kubedef.Ego()); err != nil {
						return nil, err
					}

					output.ConfigId = configId
				}

				return &execution.HandleResult{
					Outputs: []execution.Output{
						{InstanceID: kubedef.RuntimeConfigOutput(ensure.Deployable), Message: output},
					},
				}, nil
			})
		},

		PlanOrder: func(ensure *kubedef.OpEnsureRuntimeConfig) (*schema.ScheduleOrder, error) {
			return nil, nil
		},
	})
}
