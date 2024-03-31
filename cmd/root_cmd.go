package cmd

import (
	"fmt"
	"strings"

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

			packageName, err := cmd.Flags().GetString(types.PackageNameFlag)
			if err != nil {
				return err
			}

			outputFilePath, err := cmd.Flags().GetString(types.OutputFlag)
			if err != nil {
				return err
			}

			if !strings.HasSuffix(outputFilePath, ".go") {
				return fmt.Errorf("%s file must have a .go extension", outputFilePath)
			}
			isInCheckMode, err := cmd.Flags().GetBool(types.CheckFlag)
			if err != nil {
				return err
			}

			idlSchema, err := schemas.IDLSchemaFromFile(args[0])
			if err != nil {
				return err
			}

			if isInCheckMode {
				types.DefaultLogger().Info().Msgf("Checking %s", outputFilePath)
			} else {
				types.DefaultLogger().Info().Msgf("Generating code to %s", outputFilePath)
			}

			jenFile, err := codegen.GenerateCodeFromIDLSchema(idlSchema, packageName)
			if err != nil {
				return err
			}

			if isInCheckMode {
				err = codegen.ValidateOutputFile(jenFile, outputFilePath)
				if err != nil {
					return err
				}

				types.DefaultLogger().Info().Msgf("✨ All good! ✨")
				return nil
			}

			err = jenFile.Save(outputFilePath)
			if err != nil {
				return err
			}

			types.DefaultLogger().Info().Msgf("✨ All done! ✨")

			return nil
		},
	}

	cmd.Flags().StringP(types.OutputFlag, "o", "output.go", "Path to the output file. If not provided, the output will be written to 'output.go'.")
	cmd.Flags().String(types.PackageNameFlag, "", "Name of the package to be used in the generated go code. If not provided, the package name will be inferred from the contract name in the schema file.")
	cmd.Flags().Bool(types.CheckFlag, false, "Check whether or not the generated code matches the existing code in the output file.")

	return cmd
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
