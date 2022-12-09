// This file was automatically generated by Namespace.
// DO NOT EDIT. To update, re-run `ns generate`.

package datastore

import (
	"context"
	fncore "namespacelabs.dev/foundation/std/core"
	"namespacelabs.dev/foundation/std/go/core"
)

// Dependencies that are instantiated once for the lifetime of the extension.
type ExtensionDeps struct {
	ReadinessCheck core.Check
}

type _checkProvideDatabase func(context.Context, *Database, ExtensionDeps) (*DB, error)

var _ _checkProvideDatabase = ProvideDatabase

var (
	Package__iij69l = &core.Package{
		PackageName: "namespacelabs.dev/foundation/internal/testdata/datastore",
	}

	Provider__iij69l = core.Provider{
		Package:     Package__iij69l,
		Instantiate: makeDeps__iij69l,
	}
)

func makeDeps__iij69l(ctx context.Context, di core.Dependencies) (_ interface{}, err error) {
	var deps ExtensionDeps

	if deps.ReadinessCheck, err = fncore.ProvideReadinessCheck(ctx, nil); err != nil {
		return nil, err
	}

	return deps, nil
}
