package asset

import (
	"github.com/sensu/sensu-go/cli"
	"github.com/spf13/cobra"
)

// HelpCommand defines new parent
func HelpCommand(cli *cli.SensuCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:    "asset",
		Short:  "Manage assets",
		Hidden: true,
	}

	// Add sub-commands
	cmd.AddCommand(
		CreateCommand(cli),
		ListCommand(cli),
	)

	return cmd
}
