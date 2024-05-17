package chainconfig

import "github.com/gobuffalo/plush/v4"

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
