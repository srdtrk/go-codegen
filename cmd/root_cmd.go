package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/srdtrk/go-codegen/pkg/codegen"
	"github.com/srdtrk/go-codegen/pkg/schemas"
	"github.com/srdtrk/go-codegen/pkg/types"
)

func RootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "go-codegen",
		Short: "Generate Golang code for your CosmWasm smart contracts.",
	}

	rootCmd.AddCommand(
		versionCmd(),
		genMessagesCmd(),
		genInterchaintest(),
	)

	return rootCmd
}

func versionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version of go-codegen",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(types.Version)
			return nil
		},
	}

	return cmd
}

func genMessagesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "messages schema_file [flags]",
		Short: "Generate the Golang types for a CosmWasm contract from a JSON schema file.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("expected 1 argument, got %d", len(args))
			}

			packageName, err := cmd.Flags().GetString("package-name")
			if err != nil {
				return err
			}

			outputFilePath, err := cmd.Flags().GetString("output")
			if err != nil {
				return err
			}

			idlSchema, err := schemas.IDLSchemaFromFile(args[0])
			if err != nil {
				return err
			}

			err = codegen.GenerateCodeFromIDLSchema(idlSchema, outputFilePath, packageName)
			if err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringP("output", "o", "output.go", "Path to the output file. If not provided, the output will be written to 'output.go'.")
	cmd.Flags().String("package-name", "", "Name of the package to be used in the generated go code. If not provided, the package name will be inferred from the contract name in the schema file.")

	return cmd
}

func genInterchaintest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "interchaintest",
		Short: "Scaffold an end-to-end test suite for CosmWasm contracts. Ideal for testing IBC functionality.",
		Long: "Scaffold an end-to-end test suite for CosmWasm contracts using strangelove's interchaintest library. Currently supports v8 of the interchaintest library, which corresponds to wasmd v0.50.0",
	}

	cmd.AddCommand(
		interchaintestScaffold(),
	)

	return cmd
}

func interchaintestScaffold() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "scaffold",
		Short: "Scaffold an end-to-end test suite with any number of chains. Safe to use without any flags.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	cmd.Flags().String("dir", "e2e/interchaintestv8", "Directory to scaffold the test suite in, relative to the current working directory.")
	cmd.Flags().Uint8("chains", 2, "Number of chains to scaffold the test suite for.")
	cmd.Flags().String("go-module", "github.com/srdtrk/go-codegen/e2esuite/v8", "Go module name to be used in the generated test suite.")
	cmd.Flags().Bool("github-actions", false, "Generate a GitHub Actions workflow file for the test suite.")

	return cmd
}
