package cmd

import (
	"encoding/json"
	"fmt"

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
		secrets := secrets.List(config)

		for _, secret := range secrets.Secrets {
			prettyOutput, _ := json.MarshalIndent(secret, "", "  ")
			fmt.Printf("%s\n", string(prettyOutput))
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
