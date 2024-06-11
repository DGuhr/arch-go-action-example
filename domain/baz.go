package domain

import (
	"github.com/spf13/cobra"
	"os"
)

func Wrong() {
	var rootCmd = &cobra.Command{
		Use:   "authz --config <config.yaml>",
		Short: "authz service, alpha.",
		Long:  `authz service based on zanzibar access systems. alpha`,
	}

	if err := rootCmd.Execute(); err != nil {
		rootCmd.PrintErr(err)
		os.Exit(-1)
	}
}
