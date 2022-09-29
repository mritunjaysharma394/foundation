// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package docker

import (
	"namespacelabs.dev/foundation/internal/frontend/cuefrontend/entity"
	"namespacelabs.dev/foundation/internal/frontend/cuefrontend/integration/helpers"
	"namespacelabs.dev/foundation/schema"
)

func NewParser() entity.EntityParser {
	return &helpers.SimpleJsonParser[*schema.DockerIntegration]{
		SyntaxUrl:      "namespace.so/from-dockerfile",
		SyntaxShortcut: "docker",
	}
}