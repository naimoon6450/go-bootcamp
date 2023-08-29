package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var lineCount bool

var rootCmd = &cobra.Command{
	Use:   "wc",
	Short: "wc will provide the word count of a passed in file",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1)),
	Run:   handleWCCmd,
}

func main() {
	rootCmd.Flags().BoolVarP(&lineCount, "line count", "l", false, "do a line count")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// how do we take in a file
