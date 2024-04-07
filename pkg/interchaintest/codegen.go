package interchaintest

import (
	"context"

	"github.com/srdtrk/go-codegen/pkg/xgenny"
)

// GenerateTestSuite generates the interchaintest test suite
func GenerateTestSuite(moduleName, outDir string, chainNum uint8, githubActions bool) error {
	ctx := context.Background()

	generators, err := getInitGenerators(moduleName, chainNum, githubActions)
	if err != nil {
		return err
	}

	// create runners
	globalRunner := xgenny.NewRunner(ctx, ".github/workflows")
	testRunner := xgenny.NewRunner(ctx, outDir)

	err = testRunner.RunAndApply(generators...)
	if err != nil {
		return err
	}

	if githubActions {
		workflowGenerators, err := getWorkflowGenerators(outDir)
		if err != nil {
			return err
		}

		err = globalRunner.RunAndApply(workflowGenerators...)
		if err != nil {
			return err
		}
	}

	return nil
}
