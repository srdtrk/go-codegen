package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func RootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "go-codegen",
		Short: "Generate Golang code for your CosmWasm smart contracts.",
	}

	rootCmd.AddCommand(
		generateCmd(),
	)

	return rootCmd
}

func generateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate [SCHEMA_FILE]",
		Short: "Generate Golang code from a JSON schema file.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("expected 1 argument, got %d", len(args))
			}

			// TODO: Implement the code generation logic here.
			return nil
		},
	}

	return cmd
}
