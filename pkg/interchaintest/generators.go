package interchaintest

import (
	"github.com/gobuffalo/genny/v2"

	"github.com/srdtrk/go-codegen/templates/interchaintestv8"
	"github.com/srdtrk/go-codegen/templates/interchaintestv8/chainconfig"
	"github.com/srdtrk/go-codegen/templates/interchaintestv8/e2esuite"
	"github.com/srdtrk/go-codegen/templates/interchaintestv8/github"
	"github.com/srdtrk/go-codegen/templates/interchaintestv8/types"
)

func getWorkflowGenerators(outDir string) ([]*genny.Generator, error) {
	var generators []*genny.Generator

	// main generator
	g, err := github.NewGenerator(&github.Options{
		TestDir: outDir,
	})
	if err != nil {
		return nil, err
	}
	generators = append(generators, g)

	return generators, nil
}

func getInitGenerators(moduleName string, chainNum uint8, outDir string, githubActions bool) ([]*genny.Generator, error) {
	var generators []*genny.Generator

	// main generator
	g, err := interchaintestv8.NewGenerator(&interchaintestv8.Options{
		ModulePath: moduleName,
		Github:     githubActions,
		ChainNum:   chainNum,
	})
	if err != nil {
		return nil, err
	}
	generators = append(generators, g)

	// suite generator
	sg, err := e2esuite.NewGenerator(&e2esuite.Options{
		ModulePath: moduleName,
		ChainNum:   chainNum,
		TestDir:    outDir,
	})
	if err != nil {
		return nil, err
	}
	generators = append(generators, sg)

	// chain config generator
	cg, err := chainconfig.NewGenerator(&chainconfig.Options{
		ChainNum: chainNum,
	})
	if err != nil {
		return nil, err
	}
	generators = append(generators, cg)

	// types generator
	tg, err := types.NewGenerator()
	if err != nil {
		return nil, err
	}
	generators = append(generators, tg)

	return generators, nil
}
