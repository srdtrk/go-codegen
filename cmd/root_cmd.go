package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/spf13/cobra"
	"golang.org/x/mod/module"

	"github.com/srdtrk/go-codegen/pkg/codegen"
	gocmd "github.com/srdtrk/go-codegen/pkg/go/cmd"
	"github.com/srdtrk/go-codegen/pkg/interchaintest"
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

func generateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate Golang code for your CosmWasm smart contracts.",
		Long:  "Generate Golang code for your CosmWasm smart contracts from a JSON schema file.",
	}

	cmd.AddCommand(
		genMessagesCmd(),
		genQueryCliCmd(),
	)

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

			packageName, err := cmd.Flags().GetString(PackageNameFlag)
			if err != nil {
				return err
			}

			outputFilePath, err := cmd.Flags().GetString(OutputFlag)
			if err != nil {
				return err
			}

			return codegen.GenerateCodeFromIDLSchema(args[0], outputFilePath, packageName)
		},
	}

	cmd.Flags().StringP(OutputFlag, "o", "output.go", "Path to the output file.")
	cmd.Flags().String(PackageNameFlag, "p", "Package name of the generated go code. If not provided, the package name will be inferred from the contract name in the schema file.")

	return cmd
}

func genQueryCliCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query-client schema_file [flags]",
		Short: "Generate a gRPC query client for a CosmWasm contract from a JSON schema file. Requires the message types to be generated first.",
		Long:  "Generate a gRPC query client for a CosmWasm contract from a JSON schema file. Requires the message types to be generated first. This has a dependency on wasmd, so make sure to have them installed in your go project.",
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

			return codegen.GenerateQueryClientFromIDLSchema(args[0], outputFilePath, packageName)
		},
	}

	cmd.Flags().StringP(OutputFlag, "o", "query.go", "Path to the output file.")
	cmd.Flags().String(PackageNameFlag, "p", "Package name of the generated go code. Should be the same as the package name of the generated message types.")

	return cmd
}

func genInterchaintest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "interchaintest",
		Short: "Scaffold an end-to-end test suite for CosmWasm contracts. Ideal for testing IBC functionality.",
		Long:  "Scaffold an end-to-end test suite for CosmWasm contracts using strangelove's interchaintest library.\nCurrently supports v8 of the interchaintest library, which corresponds to wasmd v0.50.0.",
	}

	cmd.AddCommand(
		interchaintestScaffold(),
		ictestAddContract(),
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
				chainNum      uint8 = 2
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
							huh.NewOption[uint8]("1", 1),
							huh.NewOption[uint8]("2", 2),
							huh.NewOption[uint8]("3", 3),
						),

					huh.NewConfirm().
						Title("Would you like to generate github actions?").
						Value(&githubActions),
				),
			)

			yes, err := cmd.Flags().GetBool(YesFlag)
			if err != nil {
				return err
			}

			if yes {
				moduleName = defaultGoModule
				outDir = defaultDir
				chainNum = 2
				githubActions = false
			} else {
				err = form.Run()
				if err != nil {
					return err
				}

			}

			err = interchaintest.GenerateTestSuite(moduleName, outDir, chainNum, githubActions)
			if err != nil {
				return err
			}

			p := tea.NewProgram(spinner.New().Title("Downloading go modules..."), tea.WithContext(context.Background()), tea.WithOutput(os.Stdout))

			go func() {
				_, err := p.Run()
				if err != nil {
					panic(err)
				}
			}()

			err = gocmd.ModDownload(outDir, false)
			if err != nil {
				return err
			}

			p.Quit()
			err = p.ReleaseTerminal()
			if err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().BoolP(YesFlag, "y", false, "Skip the interactive form and use the default values.")

	return cmd
}

func ictestAddContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-contract schema_file [flags]",
		Short: "Add a contract to the interchaintest suite.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("expected 1 argument, got %d", len(args))
			}

			suiteDir, err := cmd.Flags().GetString(SuiteDirFlag)
			if err != nil {
				return err
			}

			packageName, err := cmd.Flags().GetString(ContractNameFlag)
			if err != nil {
				return err
			}

			msgsOnly, err := cmd.Flags().GetBool(MessagesOnlyFlag)
			if err != nil {
				return err
			}

			return interchaintest.AddContract(args[0], suiteDir, packageName, msgsOnly)
		},
	}

	cmd.Flags().String(SuiteDirFlag, ".", "Path to the test suite directory. If not provided, the current working directory will be used.")
	cmd.Flags().String(ContractNameFlag, "", "Name of the contract to be added to the test suite. If not provided, the contract name will be inferred from the schema file. Recommend leaving this empty.")
	cmd.Flags().Bool(MessagesOnlyFlag, false, "If set, the contract will not be added to the test suite but its messages will be added.")

	return cmd
}
