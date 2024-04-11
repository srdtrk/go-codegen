package interchaintest

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/srdtrk/go-codegen/pkg/codegen"
	gomodule "github.com/srdtrk/go-codegen/pkg/go/module"
	"github.com/srdtrk/go-codegen/pkg/schemas"
	"github.com/srdtrk/go-codegen/pkg/types"
	"github.com/srdtrk/go-codegen/pkg/xgenny"
	"github.com/srdtrk/go-codegen/templates/interchaintestv8"
	templatetypes "github.com/srdtrk/go-codegen/templates/interchaintestv8/types"
)

// GenerateTestSuite generates the interchaintest test suite
func GenerateTestSuite(moduleName, outDir string, chainNum uint8, githubActions bool) error {
	ctx := context.Background()

	types.DefaultLogger().Info().Msgf("Generating test suite in %s", outDir)

	generators, err := getInitGenerators(moduleName, chainNum, outDir, githubActions)
	if err != nil {
		return err
	}

	// create runners
	workflowsRunner := xgenny.NewRunner(ctx, ".github/workflows")
	testRunner := xgenny.NewRunner(ctx, outDir)

	err = testRunner.Run(generators...)
	if err != nil {
		return err
	}

	if githubActions {
		workflowGenerators, err := getWorkflowGenerators(outDir)
		if err != nil {
			return err
		}

		err = workflowsRunner.Run(workflowGenerators...)
		if err != nil {
			return err
		}

		err = workflowsRunner.ApplyModifications()
		if err != nil {
			return err
		}
	}

	err = testRunner.ApplyModifications()
	if err != nil {
		return err
	}

	types.DefaultLogger().Info().Msgf("✨ All done! ✨")

	return nil
}

func AddContract(schemaPath, suiteDir, contractName string, msgsOnly bool) error {
	ctx := context.Background()

	types.DefaultLogger().Info().Msgf("Adding contract %s", contractName)

	idlSchema, err := schemas.IDLSchemaFromFile(schemaPath)
	if err != nil {
		return err
	}

	if contractName == "" {
		if idlSchema.ContractName == "" {
			panic("no contract name")
		}

		contractName = idlSchema.ContractName
	}

	if err := validateSuitePath(suiteDir); err != nil {
		return err
	}

	goMod, err := gomodule.ParseAt(suiteDir)
	if err != nil {
		return err
	}

	contractDir, found, err := findLine(suiteDir, templatetypes.PlaceholderContractDir)
	if err != nil {
		return err
	}
	if !found {
		return fmt.Errorf("could not find placeholder '%s' in %s", templatetypes.PlaceholderContractDir, suiteDir)
	}

	_, foundContractTest, err := findLine(suiteDir, interchaintestv8.PlaceholderContractSuite)
	if err != nil {
		return err
	}

	nonAlphanumericRegex := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	packageName := nonAlphanumericRegex.ReplaceAllString(contractName, "")
	packageName = strings.ToLower(packageName)

	generators, err := getContractGenerators(idlSchema, packageName, goMod.Module.Mod.Path)
	if err != nil {
		return err
	}

	contractsDir, _ := filepath.Split(contractDir)
	contractRunner := xgenny.NewRunner(ctx, contractsDir)

	err = contractRunner.Run(generators...)
	if err != nil {
		return err
	}

	if !foundContractTest {
		contractTestRunner := xgenny.NewRunner(ctx, suiteDir)

		contractTestGenerators, err := getContractTestGenerators(packageName, goMod.Module.Mod.Path)
		if err != nil {
			return err
		}

		err = contractTestRunner.Run(contractTestGenerators...)
		if err != nil {
			return err
		}

		err = contractTestRunner.ApplyModifications()
		if err != nil {
			return err
		}
	}

	err = contractRunner.ApplyModifications()
	if err != nil {
		return err
	}

	msgsPath := filepath.Join(contractsDir, packageName, "msgs.go")
	err = codegen.GenerateCodeFromIDLSchema(schemaPath, msgsPath, packageName)
	if err != nil {
		_ = os.RemoveAll(filepath.Join(contractsDir, packageName))
		return err
	}

	return nil
}
