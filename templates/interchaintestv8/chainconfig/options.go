package chainconfig

import (
	"errors"

	"github.com/gobuffalo/plush/v4"
)

type Options struct {
	ChainNum uint8
}

// Validate that options are usable.
func (opts *Options) Validate() error {
	if opts.ChainNum == 0 {
		return errors.New("ChainNum must be greater than 0")
	}
	return nil
}

func (opts *Options) plushContext() *plush.Context {
	ctx := plush.NewContext()
	ctx.Set("ChainNum", int(opts.ChainNum))
	ctx.Set("between", betweenHelper)
	return ctx
}
