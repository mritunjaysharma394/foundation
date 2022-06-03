// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go/aws/session"
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/schema"
	"namespacelabs.dev/foundation/workspace/devhost"
)

func ConfiguredSessionV1(ctx context.Context, devHost *schema.DevHost, env *schema.Environment) (*session.Session, string, error) {
	conf := &Conf{}
	if !devhost.ConfigurationForEnvParts(devHost, env).Get(conf) {
		return nil, "", fnerrors.UsageError(fmt.Sprintf("Run `fn prepare --env=%s`.", env.GetName()), "Foundation has not been configured to access AWS.")
	}

	profile := conf.Profile
	if profile == "" {
		profile = "default"
	}

	sess, err := session.NewSessionWithOptions(session.Options{Profile: profile})
	return sess, profile, err
}

func ConfiguredSession(ctx context.Context, devHost *schema.DevHost, env *schema.Environment) (aws.Config, string, error) {
	conf := &Conf{}
	if !devhost.ConfigurationForEnvParts(devHost, env).Get(conf) {
		return aws.Config{}, "", fnerrors.UsageError(fmt.Sprintf("Run `fn prepare --env=%s`.", env.GetName()), "Foundation has not been configured to access AWS.")
	}

	profile := conf.Profile
	if profile == "" {
		profile = "default"
	}

	cfg, err := config.LoadDefaultConfig(ctx, config.WithSharedConfigProfile(profile))
	return cfg, profile, err
}
