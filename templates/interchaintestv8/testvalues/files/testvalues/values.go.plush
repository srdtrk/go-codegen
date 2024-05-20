package testvalues

import "time"

const (
	// StartingTokenAmount is the amount of tokens to give to each user at the start of the test.
	StartingTokenAmount int64 = 10_000_000_000
)

var (
	// Maximum period to deposit on a proposal.
	// This value overrides the default value in the gov module using the `modifyGovV1AppState` function.
	MaxDepositPeriod = time.Second * 10
	// Duration of the voting period.
	// This value overrides the default value in the gov module using the `modifyGovV1AppState` function.
	VotingPeriod = time.Second * 30
)
