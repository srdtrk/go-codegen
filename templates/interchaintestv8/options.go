package interchaintestv8

import (
	"errors"

	"golang.org/x/mod/module"
)

type Options struct {
	ModulePath string
	Github     bool
	ChainNum   uint8
}

// Validate that options are usable.
func (opts *Options) Validate() error {
	err := module.CheckPath(opts.ModulePath)
	if err != nil {
		return err
	}

	if opts.ChainNum == 0 {
		return errors.New("ChainNum must be greater than 0")
	}

	return nil
}
