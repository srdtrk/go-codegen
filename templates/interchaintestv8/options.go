package interchaintestv8

import (
	"errors"
	"path/filepath"

	"github.com/gobuffalo/plush/v4"
	"golang.org/x/mod/module"
)

type Options struct {
	ModulePath string
	Github     bool
	ChainNum   uint8
}

type ContractTestOptions struct {
	ModulePath    string
	RelPackageDir string
}

// Validate that Options is usable.
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

// Validate that ContractTestOptions is usable.
func (opts *ContractTestOptions) Validate() error {
	err := module.CheckPath(opts.ModulePath)
	if err != nil {
		return err
	}

	if !filepath.IsLocal(opts.RelPackageDir) {
		return errors.New("RelPackageDir must be a local path")
	}

	return nil
}

// plushContext returns a plush context from Options.
func (opts *Options) plushContext() *plush.Context {
	ctx := plush.NewContext()
	ctx.Set("ModulePath", opts.ModulePath)
	ctx.Set("ChainNum", int(opts.ChainNum))
	return ctx
}

// plushContext returns a plush context from ContractTestOptions.
func (opts *ContractTestOptions) plushContext() *plush.Context {
	ctx := plush.NewContext()
	ctx.Set("ModulePath", opts.ModulePath)
	ctx.Set("RelPackageDir", opts.RelPackageDir)
	return ctx
}
