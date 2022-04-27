// This file was automatically generated by Foundation.
// DO NOT EDIT. To update, re-run `fn generate`.

package multidb

import (
	"context"
	"database/sql"
	"github.com/jackc/pgx/v4/pgxpool"
	"namespacelabs.dev/foundation/std/go/core"
	"namespacelabs.dev/foundation/std/go/server"
	"namespacelabs.dev/foundation/universe/db/maria/incluster"
	incluster1 "namespacelabs.dev/foundation/universe/db/postgres/incluster"
)

// Dependencies that are instantiated once for the lifetime of the service.
type ServiceDeps struct {
	Maria    *sql.DB
	Postgres *pgxpool.Pool
}

// Verify that WireService is present and has the appropriate type.
type checkWireService func(context.Context, server.Registrar, ServiceDeps)

var _ checkWireService = WireService

var (
	Package__7cco3b = &core.Package{
		PackageName: "namespacelabs.dev/foundation/std/testdata/service/multidb",
	}

	Provider__7cco3b = core.Provider{
		Package:     Package__7cco3b,
		Instantiate: makeDeps__7cco3b,
	}
)

func makeDeps__7cco3b(ctx context.Context, di core.Dependencies) (_ interface{}, err error) {
	var deps ServiceDeps

	if err := di.Instantiate(ctx, incluster.Provider__r7qsle, func(ctx context.Context, v interface{}) (err error) {
		// name: "mariadblist"
		if deps.Maria, err = incluster.ProvideDatabase(ctx, core.MustUnwrapProto("CgttYXJpYWRibGlzdA==", &incluster.Database{}).(*incluster.Database), v.(incluster.ExtensionDeps)); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	if err := di.Instantiate(ctx, incluster1.Provider__udoubi, func(ctx context.Context, v interface{}) (err error) {
		// name: "postgreslist"
		if deps.Postgres, err = incluster1.ProvideDatabase(ctx, core.MustUnwrapProto("Cgxwb3N0Z3Jlc2xpc3Q=", &incluster1.Database{}).(*incluster1.Database), v.(incluster1.ExtensionDeps)); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return deps, nil
}
