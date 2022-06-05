package main

import (
	"github.com/KiPIDesTAN/porter-databricks/pkg/databricks"
	"github.com/spf13/cobra"
)

func buildUninstallCommand(m *databricks.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "uninstall",
		Short: "Execute the uninstall functionality of this mixin",
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.Execute()
		},
	}
	return cmd
}
