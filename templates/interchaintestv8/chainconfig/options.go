package chainconfig

import "golang.org/x/mod/module"

type Options struct {
	ModulePath string
}

// Validate that options are usable.
func (opts *Options) Validate() error {
	return module.CheckPath(opts.ModulePath)
}
