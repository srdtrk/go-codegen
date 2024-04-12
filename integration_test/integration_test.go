package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/suite"
)

type MySuite struct {
	suite.Suite

	logger       *zerolog.Logger
	goCodegenDir string
}

func TestWithMySuite(t *testing.T) {
	suite.Run(t, new(MySuite))
}

func (s *MySuite) SetupSuite() {
	logger := zerolog.New(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339},
	).Level(zerolog.TraceLevel).With().Timestamp().Caller().Logger()

	s.logger = &logger

	s.goCodegenDir = "../build/go-codegen"
}

func (s *MySuite) GenerateGoCode(schemaDir string) {
	// Generate Go code
	// nolint:gosec
	_, err := exec.Command(s.goCodegenDir, "messages", schemaDir).Output()
	s.Require().NoError(err)
}

func (s *MySuite) GenerateGoCodeTestWithSchema(schemaDir string) {
	s.GenerateGoCode(schemaDir)

	// Run tests
	// nolint:gosec
	_, err := exec.Command("golangci-lint", "run", "output.go").Output()
	s.Require().NoError(err)

	defer func() {
		_, err := exec.Command("rm", "-rf", "output.go").Output()
		s.Require().NoError(err)
	}()
}

func (s *MySuite) TestMessageComposer() {
	s.GenerateGoCodeTestWithSchema("testdata/cw-ica-controller.json")
	s.GenerateGoCodeTestWithSchema("testdata/cw3-fixed-multisig.json")
	s.GenerateGoCodeTestWithSchema("testdata/account-nft.json")
	s.GenerateGoCodeTestWithSchema("testdata/cyberpunk.json")
	s.GenerateGoCodeTestWithSchema("testdata/hackatom.json")
}

func (s *MySuite) TestInterchaintestScaffold() {
	// nolint:gosec
	_, err := exec.Command(s.goCodegenDir, "interchaintest", "scaffold", "-y", "--debug").Output()
	s.Require().NoError(err)

	err = os.Chdir("e2e/interchaintestv8")
	s.Require().NoError(err)

	s.T().Cleanup(func() {
		err = os.Chdir("../..")
		s.Require().NoError(err)

		_, err := exec.Command("rm", "-rf", "e2e").Output()
		s.Require().NoError(err)
	})

	_, err = exec.Command("golangci-lint", "run").Output()
	s.Require().NoError(err)

	// nolint:gosec
	basicCmd := exec.Command("go", "test", "-v", "-run", "TestWithBasicTestSuite", "-testify.m", "TestBasic")

	stdout, err := basicCmd.StdoutPipe()
	s.Require().NoError(err)

	err = basicCmd.Start()
	s.Require().NoError(err)

	// output command stdout
	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}

	err = basicCmd.Wait()
	s.Require().NoError(err)
}
