package testvalues

import "github.com/gobuffalo/plush/v4"

type Options struct{}

// Validate that options are usable.
func (*Options) Validate() error {
	return nil
}

func (*Options) plushContext() *plush.Context {
	ctx := plush.NewContext()
	return ctx
}
