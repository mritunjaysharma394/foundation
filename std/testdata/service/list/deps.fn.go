// This file was automatically generated by Foundation.
// DO NOT EDIT. To update, re-run `fn generate`.

package list

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"namespacelabs.dev/foundation/std/go/core"
	"namespacelabs.dev/foundation/std/go/server"
	"namespacelabs.dev/foundation/universe/db/postgres/incluster"
)

// Dependencies that are instantiated once for the lifetime of the service.
type ServiceDeps struct {
	Db *pgxpool.Pool
}

// Verify that WireService is present and has the appropriate type.
type checkWireService func(context.Context, server.Registrar, ServiceDeps)

var _ checkWireService = WireService

var (
	Package__ffkppv = &core.Package{
		PackageName: "namespacelabs.dev/foundation/std/testdata/service/list",
	}

	Provider__ffkppv = core.Provider{
		Package:     Package__ffkppv,
		Instantiate: makeDeps__ffkppv,
	}
)

func makeDeps__ffkppv(ctx context.Context, di core.Dependencies) (_ interface{}, err error) {
	var deps ServiceDeps

	if err := di.Instantiate(ctx, incluster.Provider__udoubi, func(ctx context.Context, v interface{}) (err error) {
		// name: "list"
		if deps.Db, err = incluster.ProvideDatabase(ctx, core.MustUnwrapProto("CgRsaXN0", &incluster.Database{}).(*incluster.Database), v.(incluster.ExtensionDeps)); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return deps, nil
}
