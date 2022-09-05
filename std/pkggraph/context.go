// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package pkggraph

import "namespacelabs.dev/foundation/std/planning"

type Context interface {
	planning.Context
	PackageLoader
}

type SealedContext interface {
	planning.Context
	SealedPackageLoader
}

type ContextWithMutableModule interface {
	Context
	MutableModule
}

type sealedCtx struct {
	planning.Context
	SealedPackageLoader
}

var _ SealedContext = sealedCtx{}

func MakeSealedContext(env planning.Context, pr SealedPackageLoader) SealedContext {
	return sealedCtx{Context: env, SealedPackageLoader: pr}
}