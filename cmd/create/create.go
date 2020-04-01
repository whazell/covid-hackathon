package create

import (
	"covid/cmd/create/company"
	"covid/cmd/create/fact"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "create",
		Short: "Use the admin API to create resources",
	}
	cmd.AddCommand(fact.NewCommand())
	cmd.AddCommand(company.NewCommand())
	return cmd
}
