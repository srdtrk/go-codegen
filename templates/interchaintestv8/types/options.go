package types

import (
	"errors"

	"github.com/gobuffalo/plush/v4"
	"golang.org/x/mod/module"
)

type AddContractOptions struct {
	ModulePath         string
	ContractName       string
	PackageName        string
	InstantiateMsgName string
	ExecuteMsgName     string
	QueryMsgName       string
}

// Validate that options are usable.
func (opts *AddContractOptions) Validate() error {
	err := module.CheckPath(opts.ModulePath)
	if err != nil {
		return err
	}

	if opts.PackageName == "" {
		return errors.New("PackageName must not be empty")
	}

	if opts.ContractName == "" {
		return errors.New("ContractName must not be empty")
	}

	if opts.InstantiateMsgName == "" {
		return errors.New("InstantiateMsgName must not be empty")
	}

	if opts.ExecuteMsgName == "" {
		return errors.New("ExecuteMsgName must not be empty")
	}

	if opts.QueryMsgName == "" {
		return errors.New("QueryMsgName must not be empty")
	}

	return nil
}

// plushContext returns a plush context for the options.
func (opts *AddContractOptions) plushContext() *plush.Context {
	ctx := plush.NewContext()
	ctx.Set("ModulePath", opts.ModulePath)
	ctx.Set("PackageName", opts.PackageName)
	ctx.Set("ContractName", opts.ContractName)
	ctx.Set("InstantiateMsgName", opts.InstantiateMsgName)
	ctx.Set("ExecuteMsgName", opts.ExecuteMsgName)
	ctx.Set("QueryMsgName", opts.QueryMsgName)
	return ctx
}
