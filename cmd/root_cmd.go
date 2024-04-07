package cmd

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"golang.org/x/mod/module"

	"github.com/srdtrk/go-codegen/pkg/codegen"
	"github.com/srdtrk/go-codegen/pkg/interchaintest"
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
		Long:  "Scaffold an end-to-end test suite for CosmWasm contracts using strangelove's interchaintest library. Currently supports v8 of the interchaintest library, which corresponds to wasmd v0.50.0",
	}

	cmd.AddCommand(
		interchaintestScaffold(),
	)

	return cmd
}

func interchaintestScaffold() *cobra.Command {
	const (
		defaultDir      = "e2e/interchaintestv8"
		defaultGoModule = "github.com/srdtrk/go-codegen/e2esuite/v8"
	)

	cmd := &cobra.Command{
		Use:   "scaffold",
		Short: "Scaffold an end-to-end test suite with any number of chains. Safe to use without any flags.",
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				moduleName    string
				outDir        string
				chainNum      uint8
				githubActions bool
			)

			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().
						Title("Go module name?").
						Description("Enter the go module name to be used in the generated test suite.").
						Value(&moduleName).
						Placeholder(defaultGoModule).
						Validate(func(s string) error {
							if s == "" {
								moduleName = defaultGoModule
								return nil
							}
							return module.CheckPath(s)
						}),

					huh.NewInput().
						Title("Output directory?").
						Description("Enter the output directory to scaffold the testsuite in. Relative to the current working directory.").
						Value(&outDir).
						Placeholder(defaultDir).
						Validate(func(s string) error {
							if s == "" {
								outDir = defaultDir
								return nil
							}

							if !filepath.IsLocal(s) {
								return errors.New("output directory must be a relative path")
							}

							return nil
						}),

					huh.NewSelect[uint8]().
						Title("Number of chains to scaffold?").
						Value(&chainNum).
						Options(
							huh.NewOption[uint8]("2", 2),
						),

					huh.NewConfirm().
						Title("Would you like to generate github actions?").
						Value(&githubActions),
				),
			)

			err := form.Run()
			if err != nil {
				return err
			}

			err = interchaintest.GenerateTestSuite(moduleName, outDir, chainNum, githubActions)
			if err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}
