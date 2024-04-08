package e2esuite

import "github.com/gobuffalo/plush/v4"

const numToLetter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func (opts *Options) plushContext() *plush.Context {
	ctx := plush.NewContext()
	ctx.Set("ModulePath", opts.ModulePath)
	ctx.Set("ChainNum", int(opts.ChainNum))
	ctx.Set("between", betweenHelper)
	ctx.Set("numToLetter", numToLetterHelper)
	return ctx
}

type ranger struct {
	pos int
	end int
}

func (r *ranger) Next() interface{} {
	if r.pos < r.end {
		r.pos++
		return r.pos
	}
	return nil
}

// returns a plush.Iterator that iterates from a (inclusive) to b (exclusive)
func betweenHelper(a, b int) plush.Iterator {
	return &ranger{pos: a - 1, end: b - 1}
}

func numToLetterHelper(num int) string {
	return string(numToLetter[num])
}
