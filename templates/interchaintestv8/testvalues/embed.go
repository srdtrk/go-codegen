package testvalues

import (
	"embed"
	"fmt"
	"io/fs"

	"github.com/gobuffalo/genny/v2"

	"github.com/srdtrk/go-codegen/pkg/xgenny"
)

//go:embed files/* files/**/*
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

	err = g.SelectiveFS(subfs, nil, nil, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("generator selective fs: %w", err)
	}

	g.Transformer(xgenny.Transformer(opts.plushContext()))

	return g, nil
}
