package cmd

import (
	"github.com/ibelikov/ghosctl/pkg/config"
	"github.com/ibelikov/ghosctl/pkg/secrets"
	"github.com/spf13/cobra"
)

var (
	path string
)

var applyCmd = &cobra.Command{
	Use:   "apply -f filename.yaml",
	Short: "apply secrets configuration from yaml manifest",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		config := config.New()
		secrets.Apply(config, path)
	},
}

func init() {
	applyCmd.Flags().StringVarP(&path, "file", "f", path, "Path to yaml manifest with secrets definition")
	rootCmd.AddCommand(applyCmd)
}
