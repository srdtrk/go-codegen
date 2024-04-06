package e2esuite

import (
	"embed"
	"fmt"
	"io/fs"

	"github.com/gobuffalo/genny/v2"
	"github.com/gobuffalo/plush/v4"

	"github.com/srdtrk/go-codegen/pkg/xgenny"
)

//go:embed files/* files/**/*
var files embed.FS

func NewGenerator() (*genny.Generator, error) {
	// Remove "files/" prefix
	subfs, err := fs.Sub(files, "files")
	if err != nil {
		return nil, fmt.Errorf("generator sub: %w", err)
	}

	g := genny.New()

	err = g.FS(subfs)
	if err != nil {
		return nil, fmt.Errorf("generator selective fs: %w", err)
	}

	g.Transformer(xgenny.Transformer(plush.NewContext()))

	return g, nil
}
