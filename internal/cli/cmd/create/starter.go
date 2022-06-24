// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package create

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"namespacelabs.dev/foundation/internal/cli/fncobra"
	"namespacelabs.dev/foundation/internal/console"
	"namespacelabs.dev/foundation/internal/console/colors"
)

func newStarterCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "starter",
		Short: "Creates a new workspace in a new directory from a template.",
	}

	cmd.RunE = fncobra.RunE(func(ctx context.Context, args []string) error {
		workspaceName, err := workspaceNameFromArgs(ctx, args)
		if err != nil || workspaceName == "" {
			return err
		}

		nameParts := strings.Split(workspaceName, "/")
		dirName := nameParts[len(nameParts)-1]

		if err := os.MkdirAll(dirName, 0755); err != nil {
			return err
		}

		if err := os.Chdir(dirName); err != nil {
			return err
		}

		commands := [][]string{
			{"create", "workspace", workspaceName},
			{"prepare", "local"},
		}

		stdout := console.Stdout(ctx)
		rootCmd := cmd.Root()
		for _, args := range commands {
			if err := runAndPrintCommand(ctx, stdout, rootCmd, args); err != nil {
				return err
			}
		}

		return nil
	})

	return cmd
}

func runAndPrintCommand(ctx context.Context, out io.Writer, cmd *cobra.Command, args []string) error {
	fmt.Fprintf(out, "\n > %s\n\n", colors.Ctx(ctx).Highlight.Apply(fmt.Sprintf("ns %s", strings.Join(args, " "))))
	return runCommand(ctx, cmd, args)
}

func runCommand(ctx context.Context, cmd *cobra.Command, args []string) error {
	cmdCopy := *cmd
	cmdCopy.SetArgs(args)
	return cmdCopy.ExecuteContext(ctx)
}
