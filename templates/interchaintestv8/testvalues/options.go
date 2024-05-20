package testvalues

import "github.com/gobuffalo/plush/v4"

type Options struct{}

// Validate that options are usable.
func (opts *Options) Validate() error {
	return nil
}

func (opts *Options) plushContext() *plush.Context {
	ctx := plush.NewContext()
	return ctx
}
