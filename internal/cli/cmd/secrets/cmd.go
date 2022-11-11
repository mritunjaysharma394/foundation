// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package secrets

import (
	"context"
	"io"
	"io/fs"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"namespacelabs.dev/foundation/internal/cli/fncobra"
	"namespacelabs.dev/foundation/internal/console"
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/internal/fnfs"
	"namespacelabs.dev/foundation/internal/keys"
	"namespacelabs.dev/foundation/internal/parsing"
	"namespacelabs.dev/foundation/internal/secrets"
	"namespacelabs.dev/foundation/schema"
	"namespacelabs.dev/foundation/std/cfg"
)

func NewSecretsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "secrets",
		Short: "Manage secrets for a given server.",
	}

	cmd.AddCommand(newInfoCmd())
	cmd.AddCommand(newSetCmd())
	cmd.AddCommand(newDeleteCmd())
	cmd.AddCommand(newRevealCmd())
	cmd.AddCommand(newAddReaderCmd())

	return cmd
}

type createFunc func(context.Context) (*secrets.Bundle, error)

type location struct {
	workspaceFS fnfs.ReadWriteFS
	packageName schema.PackageName
	sourceFile  string
}

func loadBundleFromArgs(ctx context.Context, env cfg.Context, locs fncobra.Locations, createIfMissing createFunc) (*location, *secrets.Bundle, error) {
	var loc fnfs.Location
	switch len(locs.Locs) {
	case 0:
		// Workspace
		loc = locs.Root.RelPackage(".")
	case 1:
		loc = locs.Locs[0]
	default:
		return nil, nil, fnerrors.New("expected a single package to be selected, saw %d", len(locs.Locs))
	}

	pkg, err := parsing.NewPackageLoader(env).LoadByName(ctx, loc.AsPackageName())
	if err != nil {
		return nil, nil, err
	}

	if !isModuleRoot(loc) && pkg.Server == nil {
		return nil, nil, fnerrors.BadInputError("%s: expected a server or a workspace root", loc.AsPackageName())
	}

	if env.Workspace().LoadedFrom() == nil {
		return nil, nil, fnerrors.InternalError("%s: missing workspace's source", loc.AsPackageName())
	}

	workspaceFS := fnfs.ReadWriteLocalFS(env.Workspace().LoadedFrom().AbsPath)
	result := &location{workspaceFS, loc.AsPackageName(), secretBundleFilename(loc)}

	contents, err := fs.ReadFile(workspaceFS, result.sourceFile)
	if err != nil {
		if os.IsNotExist(err) && createIfMissing != nil {
			bundle, err := createIfMissing(ctx)
			return result, bundle, err
		}

		return nil, nil, err
	}

	keyDir, err := keys.KeysDir()
	if err != nil {
		return nil, nil, err
	}

	bundle, err := secrets.LoadBundle(ctx, keyDir, contents)
	return result, bundle, err
}

func parseKey(v string, defaultPkgName string) (*secrets.ValueKey, error) {
	parts := strings.SplitN(v, ":", 2)
	if len(parts) < 2 {
		parts = []string{defaultPkgName, parts[0]}
	}

	return &secrets.ValueKey{PackageName: parts[0], Key: parts[1]}, nil
}

func writeBundle(ctx context.Context, loc *location, bundle *secrets.Bundle, encrypt bool) error {
	return fnfs.WriteWorkspaceFile(ctx, console.Stdout(ctx), loc.workspaceFS, loc.sourceFile, func(w io.Writer) error {
		return bundle.SerializeTo(ctx, w, encrypt)
	})
}

func secretBundleFilename(loc fnfs.Location) string {
	if isModuleRoot(loc) {
		return secrets.WorkspaceBundleName
	}

	return loc.Rel(secrets.ServerBundleName)
}

func isModuleRoot(loc fnfs.Location) bool {
	return loc.RelPath == "."
}
