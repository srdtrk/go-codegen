package e2esuite

import (
	"errors"
	"path/filepath"

	"golang.org/x/mod/module"
)

type Options struct {
	ModulePath string
	ChainNum   uint8
	TestDir    string
}

// Validate that options are usable.
func (opts *Options) Validate() error {
	err := module.CheckPath(opts.ModulePath)
	if err != nil {
		return err
	}

	if !filepath.IsLocal(opts.TestDir) {
		return errors.New("TestDir must be a local path")
	}

	if opts.ChainNum == 0 {
		return errors.New("ChainNum must be greater than 0")
	}

	return nil
}
