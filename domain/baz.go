package domain

import (
	"github.com/spf13/cobra"
	"os"
)

func Wrong() {
	var rootCmd = &cobra.Command{
		Use:   "test",
		Short: "test",
		Long:  `testtest`,
	}

	if err := rootCmd.Execute(); err != nil {
		rootCmd.PrintErr(err)
		os.Exit(-1)
	}
}
