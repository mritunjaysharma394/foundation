// This file was automatically generated by Foundation.
// DO NOT EDIT. To update, re-run `fn generate`.

package post

import (
	"context"
	"google.golang.org/grpc"
	"namespacelabs.dev/foundation/std/go/core"
	"namespacelabs.dev/foundation/std/go/server"
	fngrpc "namespacelabs.dev/foundation/std/grpc"
	"namespacelabs.dev/foundation/std/grpc/deadlines"
	"namespacelabs.dev/foundation/std/testdata/datastore"
	"namespacelabs.dev/foundation/std/testdata/service/simple"
)

// Dependencies that are instantiated once for the lifetime of the service.
type ServiceDeps struct {
	Dl         *deadlines.DeadlineRegistration
	Main       *datastore.DB
	Simple     simple.EmptyServiceClient
	SimpleConn *grpc.ClientConn
}

// Verify that WireService is present and has the appropriate type.
type checkWireService func(context.Context, server.Registrar, ServiceDeps)

var _ checkWireService = WireService

var (
	Package__j7h7h5 = &core.Package{
		PackageName: "namespacelabs.dev/foundation/std/testdata/service/post",
	}

	Provider__j7h7h5 = core.Provider{
		Package:     Package__j7h7h5,
		Instantiate: makeDeps__j7h7h5,
	}
)

func makeDeps__j7h7h5(ctx context.Context, di core.Dependencies) (_ interface{}, err error) {
	var deps ServiceDeps

	if err := di.Instantiate(ctx, deadlines.Provider__vbko45, func(ctx context.Context, v interface{}) (err error) {
		// configuration: {
		//   service_name: "PostService"
		//   method_name: "*"
		//   maximum_deadline: 5
		// }
		if deps.Dl, err = deadlines.ProvideDeadlines(ctx, core.MustUnwrapProto("ChUKC1Bvc3RTZXJ2aWNlEgEqHQAAoEA=", &deadlines.Deadline{}).(*deadlines.Deadline), v.(deadlines.ExtensionDeps)); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	if err := di.Instantiate(ctx, datastore.Provider__38f4mh, func(ctx context.Context, v interface{}) (err error) {
		// name: "main"
		// schema_file: {
		//   path: "schema.txt"
		//   contents: "just a test file"
		// }
		if deps.Main, err = datastore.ProvideDatabase(ctx, core.MustUnwrapProto("CgRtYWluEh4KCnNjaGVtYS50eHQSEGp1c3QgYSB0ZXN0IGZpbGU=", &datastore.Database{}).(*datastore.Database), v.(datastore.ExtensionDeps)); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	// package_name: "namespacelabs.dev/foundation/std/testdata/service/simple"
	if deps.SimpleConn, err = fngrpc.ProvideConn(ctx, core.MustUnwrapProto("CjhuYW1lc3BhY2VsYWJzLmRldi9mb3VuZGF0aW9uL3N0ZC90ZXN0ZGF0YS9zZXJ2aWNlL3NpbXBsZQ==", &fngrpc.Backend{}).(*fngrpc.Backend)); err != nil {
		return nil, err
	}

	deps.Simple = simple.NewEmptyServiceClient(deps.SimpleConn)

	return deps, nil
}
