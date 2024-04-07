package chainconfig

import "errors"

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
