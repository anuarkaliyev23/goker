package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const version = "0.0.1"

var rootCmd = &cobra.Command{
	Use:   "goker",
	Short: "goker is cli utility for poker",
	Version: version,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
