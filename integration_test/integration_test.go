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

func (s *MySuite) GenerateMessageTypes(schemaDir string) {
	// Generate Go code
	// nolint:gosec
	_, err := exec.Command(s.goCodegenDir, "generate", "messages", schemaDir).Output()
	s.Require().NoError(err)
}

func (s *MySuite) GenerateMessageTypesTest(schemaDir string) {
	s.Run(fmt.Sprintf("GenerateMessageTypesTest: %s", schemaDir), func() {
		s.GenerateMessageTypes(schemaDir)

		// Run tests
		// nolint:gosec
		_, err := exec.Command("golangci-lint", "run", "msgs.go").Output()
		s.Require().NoError(err)

		defer func() {
			_, err := exec.Command("rm", "-rf", "msgs.go").Output()
			s.Require().NoError(err)
		}()
	})
}

func (s *MySuite) GenerateQueryClient(schemaDir string) {
	// Generate Go code
	// nolint:gosec
	_, err := exec.Command(s.goCodegenDir, "generate", "query-client", schemaDir).Output()
	s.Require().NoError(err)
}

func (s *MySuite) GenerateQueryClientTest(schemaDir string) {
	s.Run(fmt.Sprintf("GenerateQueryClientTest: %s", schemaDir), func() {
		s.GenerateMessageTypes(schemaDir)
		s.GenerateQueryClient(schemaDir)

		defer func() {
			_, err := exec.Command("rm", "-rf", "msgs.go").Output()
			s.Require().NoError(err)
			_, err = exec.Command("rm", "-rf", "query.go").Output()
			s.Require().NoError(err)
		}()

		// Run tests
		// nolint:gosec
		_, err := exec.Command("golangci-lint", "run", "query.go", "msgs.go").Output()
		s.Require().NoError(err)
	})
}

func (s *MySuite) TestMessageComposer() {
	s.GenerateMessageTypesTest("testdata/cw-ica-controller.json")
	s.GenerateMessageTypesTest("testdata/cw3-fixed-multisig.json")
	s.GenerateMessageTypesTest("testdata/account-nft.json")
	s.GenerateMessageTypesTest("testdata/cyberpunk.json")
	s.GenerateMessageTypesTest("testdata/hackatom.json")
	s.GenerateMessageTypesTest("testdata/cw721-base.json")
	s.GenerateMessageTypesTest("testdata/cw2981-royalties.json")
	s.GenerateMessageTypesTest("testdata/ics721.json")
}

func (s *MySuite) TestQueryClient() {
	s.GenerateQueryClientTest("testdata/cw-ica-controller.json")
	s.GenerateQueryClientTest("testdata/cw3-fixed-multisig.json")
	s.GenerateQueryClientTest("testdata/account-nft.json")
	s.GenerateQueryClientTest("testdata/cyberpunk.json")
	s.GenerateQueryClientTest("testdata/hackatom.json")
	s.GenerateQueryClientTest("testdata/cw721-base.json")
	s.GenerateQueryClientTest("testdata/cw2981-royalties.json")
	s.GenerateQueryClientTest("testdata/ics721.json")
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
	basicCmd := exec.Command("go", "test", "-v", "-run", "TestWithBasicTestSuite/TestBasic")

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
