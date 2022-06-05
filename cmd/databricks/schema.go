package main

import (
	"github.com/KiPIDesTAN/porter-databricks/pkg/databricks"
	"github.com/spf13/cobra"
)

func buildSchemaCommand(m *databricks.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "schema",
		Short: "Print the json schema for the mixin",
		Run: func(cmd *cobra.Command, args []string) {
			m.PrintSchema()
		},
	}
	return cmd
}
