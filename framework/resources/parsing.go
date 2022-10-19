// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package resources

import (
	"encoding/json"

	"namespacelabs.dev/foundation/internal/fnerrors"
)

type Parser struct {
	data []byte
}

func NewParser(data []byte) *Parser {
	return &Parser{data: data}
}

func (p *Parser) Decode(resource string, out any) error {
	resources := make(map[string]any)
	if err := json.Unmarshal(p.data, &resources); err != nil {
		return err
	}

	val, ok := resources[resource]
	if !ok {
		return fnerrors.InternalError("no resource config found for resource %q", resource)
	}

	// TODO use json decoder to avoid this marshal
	data, err := json.Marshal(val)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, out)
}