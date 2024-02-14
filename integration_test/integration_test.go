package main

import (
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

	// Clean up
	s.T().Cleanup(func() {
		// nolint:gosec
		_, err := exec.Command("rm", "-rf", "output.go").Output()
		s.Require().NoError(err)
	})
}

func (s *MySuite) GenerateGoCode(schemaDir string) {
	// Generate Go code
	// nolint:gosec
	_, err := exec.Command(s.goCodegenDir, "generate", schemaDir).Output()
	s.Require().NoError(err)
}

func (s *MySuite) GenerateGoCodeTestWithSchema(schemaDir string) {
	s.GenerateGoCode(schemaDir)

	// Run tests
	// nolint:gosec
	_, err := exec.Command("golangci-lint", "run", "output.go").Output()
	s.Require().NoError(err)
}

func (s *MySuite) TestMessageComposer() {
	s.GenerateGoCodeTestWithSchema("testdata/cw-ica-controller.json")
	s.GenerateGoCodeTestWithSchema("testdata/cw3-fixed-multisig.json")
}
