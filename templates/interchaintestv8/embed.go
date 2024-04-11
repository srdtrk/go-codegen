package interchaintestv8

import (
	"embed"
	"fmt"
	"io/fs"

	"github.com/gobuffalo/genny/v2"

	"github.com/srdtrk/go-codegen/pkg/xgenny"
)

//go:embed files/*
var files embed.FS

func NewGenerator(opts *Options) (*genny.Generator, error) {
	if err := opts.Validate(); err != nil {
		return nil, fmt.Errorf("invalid options: %w", err)
	}

	// Remove "files/" prefix
	subfs, err := fs.Sub(files, "files")
	if err != nil {
		return nil, fmt.Errorf("generator sub: %w", err)
	}

	g := genny.New()

	excludeList := []string{"contract_test.go"}
	err = g.SelectiveFS(subfs, nil, nil, excludeList, nil)
	if err != nil {
		return nil, fmt.Errorf("generator selective fs: %w", err)
	}

	g.Transformer(xgenny.Transformer(opts.plushContext()))

	return g, nil
}

func NewGeneratorForContractTest(opts *ContractTestOptions) (*genny.Generator, error) {
	if err := opts.Validate(); err != nil {
		return nil, fmt.Errorf("invalid options: %w", err)
	}

	// Remove "files/" prefix
	subfs, err := fs.Sub(files, "files")
	if err != nil {
		return nil, fmt.Errorf("generator sub: %w", err)
	}

	g := genny.New()

	includeList := []string{"contract_test.go"}
	err = g.SelectiveFS(subfs, includeList, nil, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("generator selective fs: %w", err)
	}

	g.Transformer(xgenny.Transformer(opts.plushContext()))

	return g, nil
}
