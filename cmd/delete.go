package cmd

import (
	"fmt"

	"github.com/ibelikov/ghosctl/pkg/config"
	"github.com/ibelikov/ghosctl/pkg/secrets"
	"github.com/spf13/cobra"
)

var (
	deleteName string
)

var deleteCmd = &cobra.Command{
	Use:   "delete -n [name]",
	Short: "delete org secret with given name",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		config := config.New()
		response := secrets.Delete(config, name)
		fmt.Printf("%v", response.Response.Body)
	},
}

func init() {
	deleteCmd.Flags().StringVarP(&deleteName, "name", "n", deleteName, "Name of the Org secret")
	deleteCmd.MarkFlagRequired("name")
	rootCmd.AddCommand(createCmd)
}