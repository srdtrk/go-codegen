---
title: Introduction
sidebar_label: Introduction
sidebar_position: 0
slug: /
---

import HighlightTag from '@site/src/components/HighlightTag';
import HighlightBox from '@site/src/components/HighlightBox';

# `go-codegen` Tutorial

<HighlightTag type="tutorial"/><HighlightTag type="guided-coding"/><HighlightTag type="cosmwasm"/><HighlightTag type="developer"/><HighlightTag type="ibc-go"/><HighlightTag type="advanced"/>

This is a tutorial for writing end-to-end test suite for IBC enabled CosmWasm applications using go-codegen.
Powered by [interchaintest](https://github.com/strangelove-ventures/interchaintest).

<HighlightBox type="prerequisite" title="Prerequisites">

- Basic knowledge of [golang](TODO)
- Basic knowledge of [CosmWasm](TODO)
- Basic knowledge of [github actions](TODO)

</HighlightBox>

## Context

### What is interchaintest?

TODO

### Why use this over other scripting solutions?

- Go is a great scripting language.
- This runs completely locally while closely resembling production environment.
    - This allows you to run all the tests in CI.
    - Reliable.
- Interchaintest is used by many core teams, and top cosmwasm projects.
    - IBC team, Strangelove, Osmosis, DAODAO, TODO
- Intterchaintest's development is funded by the ICF.

## Scope

This tutorial will cover spinning up a local testing environment with interchaintest. We will be testing [cw-ica-controller](https://github.com/srdtrk/cw-ica-controller) to control a native interchain account in the host chain.

<HighlightBox type="learning" title="Learning Goals">

In this tutorial, you will:

- Compile a CosmWasm contract into wasm and generate its JSON schema
- Generate its message types using go-codegen. (Optional)
- Scaffold an end-to-end testsuite with github actions consisting of two chains using go-codegen.
- Add our contract to the testsuite using go-codegen.
- Create a transfer channel between the two chains.
- Create an IBC channel between the contract and the counterparty go module.
- Send an ICA transaction and verify it.

</HighlightBox>
