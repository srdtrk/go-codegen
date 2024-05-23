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

- Basic knowledge of [golang](https://go.dev/)
- Basic knowledge of [CosmWasm](https://cosmwasm.com/)
- Basic knowledge of [github actions](https://github.com/features/actions)

</HighlightBox>

## Context

### What is go-codegen?

go-codegen is a tool that generates go code for CosmWasm applications. It can currently generate the following:

- Message definitions for a contract.
- gRPC query client for a contract.
- An end-to-end test suite for a contract based on interchaintest.

<HighlightBox type="info" title="What is interchaintest?">

`interchaintest` is a framework for testing blockchain functionality and interoperability between chains, primarily with the Inter-Blockchain Communication (IBC) protocol.

It can quickly spin up custom testnets and dev environments to test IBC, Relayer setup, chain infrastructure, smart contracts, etc. Interchaintest orchestrates Go tests that utilize Docker containers for multiple IBC-compatible blockchains.

In order to ship production-grade software for the Interchain, we needed sophisticated developer tooling...but IBC and Web3 have a lot of moving parts, which can lead to a steep learning curve and all sorts of pain.

- repeatedly building repo-specific, Docker- and shell-based testing solutions,
- duplication of effort, and
- difficulty in repurposing existing testing harnesses for new problem domains.

Read more about interchaintest in its [README.md](https://github.com/strangelove-ventures/interchaintest/blob/main/README.md).

</HighlightBox>

### Why use this over other scripting solutions?

- **Go as a Scripting Language**
- **Local Testing Environment:** The testing framework operates entirely locally, mirroring a production environment closely. This ensures that tests are reliable and can be seamlessly integrated into continuous integration (CI) pipelines.
- **Industry Adoption:** Interchaintest, the underlying framework, is widely adopted by leading projects and core teams in the CosmWasm ecosystem, including Strangelove, the IBC team, Noble, DAODAO, and others.
- **Funding and Support:** The development of Interchaintest is supported by the Interchain Foundation (ICF), ensuring ongoing improvements and stability.

By using go-codegen, your CosmWasm application can benefit from a well-supported, reliable, widely-adopted testing framework.

## Scope

This tutorial will cover spinning up a local testing environment with interchaintest. We will be testing [cw-ica-controller](https://github.com/srdtrk/cw-ica-controller) to control a native interchain account in the host chain.

<HighlightBox type="learning" title="Learning Goals">

In this tutorial, you will:

- Compile a CosmWasm contract into wasm and generate its JSON schema
- Scaffold an end-to-end testsuite with github actions consisting of two chains using go-codegen.
- Add our contract to the testsuite using go-codegen.
- Create a transfer channel between the two chains.
- Create an IBC channel between the contract and the counterparty go module.
- Send an ICA transaction and verify it.

</HighlightBox>
