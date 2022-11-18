// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package opaque

import "namespacelabs.dev/foundation/schema"

const (
	forceProd = false
)

func UseDevBuild(env *schema.Environment) bool {
	return !forceProd && env.Purpose == schema.Environment_DEVELOPMENT
}