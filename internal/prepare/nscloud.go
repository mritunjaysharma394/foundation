// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package prepare

import (
	"context"
	"encoding/json"

	"namespacelabs.dev/foundation/build/registry"
	"namespacelabs.dev/foundation/internal/engine/ops"
	"namespacelabs.dev/foundation/providers/nscloud"
	"namespacelabs.dev/foundation/runtime/kubernetes/client"
	"namespacelabs.dev/foundation/schema"
	"namespacelabs.dev/foundation/workspace/compute"
	"namespacelabs.dev/foundation/workspace/devhost"
	"namespacelabs.dev/foundation/workspace/tasks"
)

func PrepareNewNamespaceCluster(env ops.Environment) compute.Computable[[]*schema.DevHost_ConfigureEnvironment] {
	return compute.Map(
		tasks.Action("prepare.nscloud.new-cluster"),
		compute.Inputs().Proto("env", env.Proto()).Indigestible("foobar", "foobar"),
		compute.Output{NotCacheable: true},
		func(ctx context.Context, _ compute.Resolved) ([]*schema.DevHost_ConfigureEnvironment, error) {
			cfg, err := nscloud.CreateClusterForEnv(ctx, env.Proto(), false)
			if err != nil {
				return nil, err
			}

			serializedConfig, err := json.Marshal(cfg.KubeConfig)
			if err != nil {
				return nil, err
			}

			k8sHostEnv := &client.HostEnv{
				Provider: "nscloud",
			}

			registryProvider := &registry.Provider{
				Provider: "nscloud",
			}

			prebuilt := &nscloud.PrebuiltCluster{
				ClusterId:        cfg.ClusterId,
				SerializedConfig: serializedConfig,
			}

			c, err := devhost.MakeConfiguration(k8sHostEnv, prebuilt, registryProvider)
			if err != nil {
				return nil, err
			}

			c.Name = env.Proto().Name
			return []*schema.DevHost_ConfigureEnvironment{c}, nil
		})
}