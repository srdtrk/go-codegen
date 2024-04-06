package interchaintest

import (
	"context"

	"github.com/srdtrk/go-codegen/pkg/xgenny"
	"github.com/srdtrk/go-codegen/templates/interchaintestv8"
	"github.com/srdtrk/go-codegen/templates/interchaintestv8/e2esuite"
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

	// suite generator
	sg, err := e2esuite.NewGenerator()
	if err != nil {
		return err
	}

	// new runner
	r := xgenny.NewRunner(ctx, outDir)

	err = r.RunAndApply(g, sg)
	if err != nil {
		return err
	}

	return nil
}
