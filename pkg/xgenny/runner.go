package xgenny

import (
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/gobuffalo/genny/v2"
)

type Runner struct {
	*genny.Runner
	ctx     context.Context
	results []genny.File
	tmpPath string
}

// NewRunner is a xgenny Runner with a logger.
func NewRunner(ctx context.Context, root string) *Runner {
	var (
		runner  = genny.WetRunner(ctx)
		tmpPath = filepath.Join(os.TempDir(), Runes(5))
	)
	runner.Root = root
	r := &Runner{
		ctx:     ctx,
		Runner:  runner,
		tmpPath: tmpPath,
		results: make([]genny.File, 0),
	}
	runner.FileFn = func(f genny.File) (genny.File, error) {
		return wetFileFn(r, f)
	}
	return r
}

// ApplyModifications copy all modifications from the temporary folder to the target path.
func (r *Runner) ApplyModifications() error {
	if _, err := os.Stat(r.tmpPath); os.IsNotExist(err) {
		return err
	}

	// Create the target path and copy the content from the temporary folder.
	if err := os.MkdirAll(r.Root, os.ModePerm); err != nil {
		return err
	}
	err := CopyFolder(r.tmpPath, r.Root)
	if err != nil {
		return err
	}

	return os.RemoveAll(r.tmpPath)
}

// RunAndApply run the generators and apply the modifications to the target path.
func (r *Runner) RunAndApply(gens ...*genny.Generator) error {
	if err := r.Run(gens...); err != nil {
		return err
	}
	return r.ApplyModifications()
}

// Run all generators into a temp folder for we can apply the modifications later.
func (r *Runner) Run(gens ...*genny.Generator) error {
	// execute the modification with a wet runner
	for _, gen := range gens {
		if err := r.With(gen); err != nil {
			return err
		}
		if err := r.Runner.Run(); err != nil {
			return err
		}
	}
	r.results = append(r.results, r.Results().Files...)
	return nil
}

func wetFileFn(runner *Runner, f genny.File) (genny.File, error) {
	if d, ok := f.(genny.Dir); ok {
		if err := os.MkdirAll(d.Name(), d.Perm); err != nil {
			return f, err
		}
		return d, nil
	}

	if filepath.IsAbs(runner.Root) {
		return nil, errors.New("root path must be relative")
	}

	name := f.Name()
	if !filepath.IsAbs(name) {
		name = filepath.Join(runner.Root, name)
	}
	relPath, err := filepath.Rel(runner.Root, name)
	if err != nil {
		return f, err
	}

	dstPath := filepath.Join(runner.tmpPath, relPath)
	dir := filepath.Dir(dstPath)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return f, err
	}
	ff, err := os.Create(dstPath)
	if err != nil {
		return f, err
	}
	defer ff.Close()
	if _, err := io.Copy(ff, f); err != nil {
		return f, err
	}
	return f, nil
}
