package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "GHOScTl",
	Short: "GHOScTl is a small CLI to manage GitHub Org Secrets",
}

// Execute is the main entrance point for cobra CLI
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
