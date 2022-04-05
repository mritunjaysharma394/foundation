// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package creds

import (
	"context"
)

func ProvideCreds(ctx context.Context, caller string, _ *CredsRequest, deps ExtensionDeps) (*Creds, error) {
	creds := &Creds{
		Password: string(deps.Password.Value),
	}
	return creds, nil
}
