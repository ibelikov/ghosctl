package cmd

import (
	"github.com/ibelikov/org-secrets-manager/pkg/config"
	"github.com/ibelikov/org-secrets-manager/pkg/secrets"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists org secrets",
	Run: func(cmd *cobra.Command, args []string) {
		config := config.New()
		secrets.List(config)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
