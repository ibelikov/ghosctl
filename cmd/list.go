package cmd

import (
	"github.com/ibelikov/ghosctl/pkg/config"
	"github.com/ibelikov/ghosctl/pkg/secrets"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists org secrets",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		config := config.New()
		secrets.List(config)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
