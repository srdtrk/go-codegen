package interchaintest

import (
	"context"

	"github.com/srdtrk/go-codegen/pkg/xgenny"
)

// GenerateTestSuite generates the interchaintest test suite
func GenerateTestSuite(moduleName, outDir string, chainNum uint8, githubActions bool) error {
	ctx := context.Background()

	generators, err := getInitGenerators(moduleName, githubActions)
	if err != nil {
		return err
	}

	// new runner
	r := xgenny.NewRunner(ctx, outDir)

	err = r.RunAndApply(generators...)
	if err != nil {
		return err
	}

	return nil
}
