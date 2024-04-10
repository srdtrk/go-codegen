package interchaintest

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/srdtrk/go-codegen/templates/interchaintestv8"
)

func validateSuitePath(suitePath string) error {
	goModPath := filepath.Join(suitePath, "go.mod")

	contains, err := fileContainsLine(goModPath, interchaintestv8.PlaceholderSuiteModule)
	if err != nil {
		return err
	}

	if !contains {
		return errors.New("go.mod does not contain the placeholder for suite module")
	}

	return nil
}

func findLine(targetDir, line string) (string, bool, error) {
	var (
		found bool
		resPath string
	)
	err := filepath.WalkDir(targetDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}

		contains, err := fileContainsLine(path, line)
		if err != nil {
			return nil
		}

		if contains {
			found = true
			resPath = path
			return fs.SkipAll
		}

		return nil
	})
	if err != nil {
		return "", false, err
	}

	return resPath, found, nil
}

func fileContainsLine(filePath, targetLine string) (bool, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return false, err
	}

	return strings.Contains(string(file), targetLine), nil
}
