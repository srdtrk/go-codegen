---
title: Add Contracts to the Test Suite
sidebar_label: Add Contracts to the Test Suite
sidebar_position: 4
slug: /add-contracts
---

import HighlightBox from '@site/src/components/HighlightBox';

# Add Contracts to the Test Suite

It's time to add the contracts to the test suite. It might be helpful to give a brief overview of the contracts before adding them to the test suite.

## Overview of `cw-ica-controller`

CosmWasm Interchain Account Controller is a pure CosmWasm implementation of the [ICS-27 controller specifications](https://github.com/cosmos/ibc/tree/main/spec/app/ics-027-interchain-accounts). It can therefore directly create and control accounts in the counterparty `icahost` module.

![cw-ica-controller](https://raw.githubusercontent.com/srdtrk/cw-ica-controller/main/docs/static/img/cw-ica-controller.svg)

This contract was designed to be instantiated by other contracts that wish to control accounts in a counterparty chain. Learn more about the contract in its repository: [srdtrk/cw-ica-controller](https://github.com/srdtrk/cw-ica-controller).

<HighlightBox type="note" title="Note">

There are a couple important things to note about this contract:

- Since this contract was designed to be instantiated by other contracts, it does have the capability to make callbacks to an external contract on channel and packet lifecycle events. This is what the `callback-counter` test contract is for.

<!-- TODO: replace link with workshop repo -->
```rust reference
https://github.com/srdtrk/cw-ica-controller/blob/d6b033092071e37f2dd015b58810a02257a92b6b/src/types/callbacks.rs#L15-L46
```

- This contract automatically initiates the channel opening handshake when it is instantiated and doesn't allow any relayers to initiate the handshake.

<!-- TODO: replace link with workshop repo -->
```rust reference
https://github.com/srdtrk/cw-ica-controller/blob/d6b033092071e37f2dd015b58810a02257a92b6b/src/types/msg.rs#L8-L21
```

</HighlightBox>

## Add the Contracts

In this test, we will only add the `cw-ica-controller` and `callback-counter` contracts to the test suite. To add the `cw-ica-controller` contract to the test suite, run the following command from the root of the workshop repository:

<!-- TODO: link commit -->
```sh
go-codegen interchaintest add-contract schema/cw-ica-controller.json --suite-dir e2e/interchaintestv8
```

This command will generate the following files:

```text
e2e/interchaintestv8/
├── contract_test.go
└── types
    └── cwicacontroller
        ├── contract.go
        ├── msgs.go
        └── query.go
```

`contract_test.go` contains an example test case for instantiating the generated contract. This file is only generated if it doesn't already exist in the test suite when the `add-contract` command is run.

You can also add the `callback-counter` contract to the test suite by running the following command:

<!-- TODO: link commit -->
```sh
go-codegen interchaintest add-contract testing/contracts/callback-counter/schema/callback-counter.json --suite-dir e2e/interchaintestv8
```

This command will generate the following files:

```text
e2e/interchaintestv8/types/callbackcounter
├── contract.go
├── msgs.go
└── query.go
```

This will not modify the `contract_test.go` and the 

<HighlightBox type="tip" title="Exploration">

Explore the generated files to understand how the contracts are added to the test suite.

Run the generated test and understand why it fails.

```sh
cd e2e/interchaintestv8
go test -v . -run=TestWithContractTestSuite/TestContract
```

</HighlightBox>
