package interchaintest

import (
	"github.com/gobuffalo/genny/v2"

	"github.com/srdtrk/go-codegen/templates/interchaintestv8"
	"github.com/srdtrk/go-codegen/templates/interchaintestv8/chainconfig"
	"github.com/srdtrk/go-codegen/templates/interchaintestv8/e2esuite"
)

func GetGenerators(moduleName string) ([]*genny.Generator, error) {
	var generators []*genny.Generator

	// main generator
	g, err := interchaintestv8.NewGenerator(&interchaintestv8.Options{
		ModulePath: moduleName,
	})
	if err != nil {
		return nil, err
	}
	generators = append(generators, g)

	// suite generator
	sg, err := e2esuite.NewGenerator()
	if err != nil {
		return nil, err
	}
	generators = append(generators, sg)

	// chain config generator
	cg, err := chainconfig.NewGenerator(&chainconfig.Options{
		ModulePath: moduleName,
	})
	if err != nil {
		return nil, err
	}
	generators = append(generators, cg)

	return generators, nil
}
