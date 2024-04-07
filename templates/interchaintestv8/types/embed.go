package types

import (
	"embed"
	"fmt"
	"io/fs"

	"github.com/gobuffalo/genny/v2"
	"github.com/gobuffalo/plush/v4"

	"github.com/srdtrk/go-codegen/pkg/xgenny"
)

var (
	//go:embed files/**/*.go.plush
	files embed.FS

	//nolint:unused
	//go:embed files/types/{{contractName}}/*
	contractFiles embed.FS
)

func NewGenerator() (*genny.Generator, error) {
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

	g.Transformer(xgenny.Transformer(plush.NewContext()))

	return g, nil
}
