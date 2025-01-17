package cmd

import (
	"log"

	"github.com/ibelikov/ghosctl/pkg/config"
	"github.com/ibelikov/ghosctl/pkg/secrets"
	"github.com/spf13/cobra"
)

var (
	name  string
	value string
	repos []string
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create org secret with given name and value",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		config := config.New()
		secrets.Create(config, name, value, repos)
		log.Printf("Created Organization Secret %s", name)
	},
}

func init() {
	createCmd.Flags().StringVarP(&name, "name", "n", name, "Name of the Org secret")
	createCmd.Flags().StringVarP(&value, "value", "v", value, "Value of the Org secret")
	createCmd.Flags().StringSliceVarP(&repos, "repos", "r", repos, "List of repository names that could use this secret")
	createCmd.MarkFlagRequired("name")
	createCmd.MarkFlagRequired("value")
	rootCmd.AddCommand(createCmd)
}
