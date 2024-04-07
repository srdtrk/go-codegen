package interchaintestv8

import "golang.org/x/mod/module"

type Options struct {
	ModulePath string
	Github     bool
}

// Validate that options are usable.
func (opts *Options) Validate() error {
	return module.CheckPath(opts.ModulePath)
}
