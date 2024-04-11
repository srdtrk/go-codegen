//nolint:gosec
package gocmd

import (
	"os"
	"os/exec"
)

const (
	// CommandInstall represents go "install" command.
	CommandInstall = "install"

	// CommandGet represents go "get" command.
	CommandGet = "get"

	// CommandBuild represents go "build" command.
	CommandBuild = "build"

	// CommandMod represents go "mod" command.
	CommandMod = "mod"

	// CommandModTidy represents go mod "tidy" command.
	CommandModTidy = "tidy"

	// CommandModVerify represents go mod "verify" command.
	CommandModVerify = "verify"

	// CommandModDownload represents go mod "download" command.
	CommandModDownload = "download"

	// CommandFmt represents go "fmt" command.
	CommandFmt = "fmt"

	// CommandEnv represents go "env" command.
	CommandEnv = "env"

	// CommandList represents go "list" command.
	CommandList = "list"

	// CommandTest represents go "test" command.
	CommandTest = "test"

	// EnvGOARCH represents GOARCH variable.
	EnvGOARCH = "GOARCH"
	// EnvGOMOD represents GOMOD variable.
	EnvGOMOD = "GOMOD"
	// EnvGOOS represents GOOS variable.
	EnvGOOS = "GOOS"

	// FlagGcflags represents gcflags go flag.
	FlagGcflags = "-gcflags"
	// FlagGcflagsValueDebug represents debug go flags.
	FlagGcflagsValueDebug = "all=-N -l"
	// FlagLdflags represents ldflags go flag.
	FlagLdflags = "-ldflags"
	// FlagTags represents tags go flag.
	FlagTags = "-tags"
	// FlagMod represents mod go flag.
	FlagMod = "-mod"
	// FlagModValueReadOnly represents readonly go flag.
	FlagModValueReadOnly = "readonly"
	// FlagOut represents out go flag.
	FlagOut = "-o"
)

// Env returns the value of `go env name`.
func Env(name string) (string, error) {
	bytes, err := exec.Command(Name(), CommandEnv, name).Output()
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// Name returns the name of Go binary to use.
func Name() string {
	custom := os.Getenv("GONAME")
	if custom != "" {
		return custom
	}
	return "go"
}

// Fmt runs go fmt on path.
func Fmt(path string) error {
	cmd := exec.Command(Name(), CommandFmt, "./...")
	cmd.Dir = path

	_, err := cmd.Output()
	return err
}

// ModTidy runs go mod tidy on path with options.
func ModTidy(path string) error {
	cmd := exec.Command(Name(), CommandMod, CommandModTidy)
	cmd.Dir = path

	_, err := cmd.Output()
	return err
}

// ModVerify runs go mod verify on path with options.
func ModVerify(path string) error {
	cmd := exec.Command(Name(), CommandMod, CommandModVerify)
	cmd.Dir = path

	_, err := cmd.Output()
	return err
}

// ModDownload runs go mod download on a path with options.
func ModDownload(path string, json bool) error {
	command := []string{CommandMod, CommandModDownload}
	if json {
		command = append(command, "-json")
	}

	cmd := exec.Command(Name(), command...)
	cmd.Dir = path

	_, err := cmd.Output()
	return err
}

// Get runs go get pkgs on path with options.
func Get(path string, pkgs []string) error {
	command := []string{CommandGet}
	command = append(command, pkgs...)

	cmd := exec.Command(Name(), command...)
	cmd.Dir = path

	_, err := cmd.Output()
	return err
}
