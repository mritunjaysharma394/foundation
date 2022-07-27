// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package debug

import (
	"github.com/spf13/cobra"
)

func NewDebugCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "debug",
		Short:   "Internal commands, for debugging.",
		Aliases: []string{"d", "dbg"},
	}

	cmd.AddCommand(newPrintComputedCmd())
	cmd.AddCommand(newComputeConfigCmd())
	cmd.AddCommand(newPrintSealedCmd())
	cmd.AddCommand(newImageIndexCmd())
	cmd.AddCommand(newImageCmd())
	cmd.AddCommand(newActionDemoCmd())
	cmd.AddCommand(newGoSourcesCmd())
	cmd.AddCommand(newDownloadCmd())
	cmd.AddCommand(newPrepareCmd())
	cmd.AddCommand(newDnsQuery())
	cmd.AddCommand(newObserveFsCmd())
	cmd.AddCommand(newDecodeProtoCmd())
	cmd.AddCommand(newUpdateLicenseCmd())
	cmd.AddCommand(newKubernetesCmd())
	cmd.AddCommand(newFindConfigCmd())

	return cmd
}
