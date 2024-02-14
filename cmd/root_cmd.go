package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/srdtrk/go-codegen/pkg/codegen"
	"github.com/srdtrk/go-codegen/pkg/schemas"
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
		Use:   "generate [schema_file]",
		Short: "Generate Golang code from a JSON schema file.",
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

			// TODO: Implement the code generation logic here.
			return nil
		},
	}

	cmd.Flags().StringP("output", "o", "output.go", "Path to the output file. If not provided, the output will be written to 'output.go'.")
	cmd.Flags().String("package-name", "", "Name of the package to be used in the generated go code. If not provided, the package name will be inferred from the contract name in the schema file.")

	return cmd
}
