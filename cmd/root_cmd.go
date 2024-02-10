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

			outputFilePath, err := cmd.Flags().GetString("output")
			if err != nil {
				return err
			}
			if outputFilePath == "" {
				outputFilePath = "output.go"
			}

			idlSchema, err := schemas.IDLSchemaFromFile(args[0])
			if err != nil {
				return err
			}

			err = codegen.GenerateCodeFromIDLSchema(idlSchema, outputFilePath)
			if err != nil {
				return err
			}

			// TODO: Implement the code generation logic here.
			return nil
		},
	}

	cmd.Flags().String("output", "", "Path to the output file. If not provided, the output will be written to 'output.go'.")

	return cmd
}
