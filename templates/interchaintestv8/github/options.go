package github

import (
	"errors"
	"path/filepath"
)

type Options struct {
	TestDir string
}

// Validate that options are usable.
func (opts *Options) Validate() error {
	if !filepath.IsLocal(opts.TestDir) {
		return errors.New("TestDir must be a local path")
	}
	return nil
}
