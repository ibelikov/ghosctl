package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/ibelikov/ghosctl/pkg/config"
	"github.com/ibelikov/ghosctl/pkg/secrets"
	"github.com/spf13/cobra"
)

var (
	getName string
)

var getCmd = &cobra.Command{
	Use:   "get -n [name]",
	Short: "get org secret with given name",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		config := config.New()
		secret := secrets.Get(config, getName)
		prettyOutput, _ := json.MarshalIndent(secret, "", "  ")
		fmt.Printf("%s\n", string(prettyOutput))
	},
}

func init() {
	getCmd.Flags().StringVarP(&getName, "name", "n", getName, "Name of the Org secret")
	getCmd.MarkFlagRequired("name")
	rootCmd.AddCommand(getCmd)
}
