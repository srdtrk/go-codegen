package interchaintest

import (
	"context"

	"github.com/srdtrk/go-codegen/pkg/xgenny"
	"github.com/srdtrk/go-codegen/templates/interchaintestv8"
)

// GenerateTestSuite generates the interchaintest test suite
func GenerateTestSuite(moduleName, outDir string, chainNum uint8, githubActions bool) error {
	ctx := context.Background()

	// main generator
	g, err := interchaintestv8.NewGenerator(&interchaintestv8.Options{
		ModulePath: moduleName,
	})
	if err != nil {
		return err
	}

	// new runner
	r := xgenny.NewRunner(ctx, outDir)

	err = r.RunAndApply(g)
	if err != nil {
		return err
	}

	return nil
}
