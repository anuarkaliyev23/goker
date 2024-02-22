package cmd

import "github.com/spf13/cobra"

var texasCmd = &cobra.Command{
	Use: "texas",
	Short: "utilities for texas hold'em",
}

func init() {
	rootCmd.AddCommand(texasCmd)
}

