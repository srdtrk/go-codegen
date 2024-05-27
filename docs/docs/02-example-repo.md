---
title: Setup the Example Contract Repository
sidebar_label: Setup the Example Contract Repository
sidebar_position: 2
slug: /example-repo
---

import HighlightBox from '@site/src/components/HighlightBox';

# Clone the Example Contract Repository

In this tutorial, we will be creating a test suite for an example CosmWasm contract using go-codegen. We first need to clone the repository containing the example contract.

You may also fork the repository to your GitHub account and clone it from there. This will allow you to push changes to your forked repository and test github actions in your forked repository.

```sh
git clone https://github.com/srdtrk/awesomwasm-2024-workshop
cd awesomwasm-2024-workshop
```

Now, create and checkout a new branch for the tutorial:

```sh
git checkout -b tutorial
```

You may take a look at the repository to familiarize yourself, but we will not be making any changes to the contract code in this tutorial.

<HighlightBox type="note" title="Note">

This contract contains one main contract and two testing contracts in the `testing` directory. We will be using all three contracts in our test suite.

</HighlightBox>

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

<HighlightBox type="tip" title="Exploration (Optional)">

At this point, you may explore go-codegen by running generate commands. For example:

```sh
go-codegen generate messages schema/cw-ica-controller.json
```

And you can explore the generated `msgs.go` file in the root directory. Then remove it:

```sh
rm msgs.go
```

</HighlightBox>

<HighlightBox type="note" title="Note">

Build and schema artifacts are not committed to the repository due to `.gitignore` settings.

</HighlightBox>
