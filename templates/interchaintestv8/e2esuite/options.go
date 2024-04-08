package e2esuite

import "golang.org/x/mod/module"

type Options struct {
	ModulePath string
	ChainNum   uint8
}

// Validate that options are usable.
func (opts *Options) Validate() error {
	return module.CheckPath(opts.ModulePath)
}
