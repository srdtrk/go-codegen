---
title: Scaffold a Test Suite with go-codegen
sidebar_label: Scaffold a Test Suite with go-codegen
sidebar_position: 2
slug: /scaffold-suite
---

import CodeBlock from '@theme/CodeBlock';

# Scaffold a Test Suite with go-codegen

In this tutorial, we will be creating a test suite for an example CosmWasm contract using go-codegen. We first need to clone the repository containing the example contract.

## Clone the Example Contract Repository

You may also fork the repository to your GitHub account and clone it from there. This will allow you to push changes to your forked repository and test github actions in your forked repository.

```sh
git clone https://github.com/srdtrk/awesomwasm-2024-workshop
cd awesomwasm-2024-workshop
```

You may take a look at the repository to familiarize yourself, but we will not be making any changes to the contract code in this tutorial.

Note that this contract contains one main contract and two testing contracts in the `testing` directory. We will be using all three contracts in our test suite.

### Build the Contracts

To build the contracts, we use the [`cosmwasm/optimizer`](https://github.com/CosmWasm/optimizer) docker image version `0.15.1`.

```sh
just build-optimize
```

```sh
just build-test-contracts
```

### Generate CosmWasm Schemas

To use go-codegen, we need to generate the CosmWasm schema for the contracts. We can do this by running the following command:

```sh
just generate-schemas
```
