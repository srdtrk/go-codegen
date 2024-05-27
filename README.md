# go-codegen

Generate Go code for your CosmWasm smart contracts.

<p align="center">
  <img width="230" src="go-codegen.svg">
</p>

<div align="center">

  [![E2E Status](https://github.com/srdtrk/go-codegen/actions/workflows/integration.yml/badge.svg?branch=main)](https://github.com/srdtrk/go-codegen/tree/main/integration_test)
  [![Tag](https://img.shields.io/github/tag/srdtrk/go-codegen.svg)](https://github.com/srdtrk/go-codegen/releases/latest)
  [![License: Apache-2.0](https://img.shields.io/github/license/srdtrk/go-codegen.svg)](https://github.com/srdtrk/go-codegen/blob/main/LICENSE)
  [![Lines of Code](https://tokei.rs/b1/github/srdtrk/go-codegen)](https://github.com/srdtrk/go-codegen)
</div>

```
go install github.com/srdtrk/go-codegen@latest
```

The quickest way to generate Go code for your CosmWasm smart contracts. Currently, this tool supports generating:

- Message types
- A gRPC query client
- An [`interchaintest`](https://github.com/strangelove-ventures/interchaintest) based test suite

**You can find a tutorial on how to use this tool's interchaintest feature [here](https://srdtrk.github.io/go-codegen/).**

## Usage

In order to generate Go code for your CosmWasm smart contract, you need to have its full API exported as a JSON file
using [`cosmwasm-schema`](https://crates.io/crates/cosmwasm-schema).
Once you have the JSON file, you can use it to generate the Go code.

Note that you can learn more about how to use this tool by running `go-codegen help` or `-h` flag with any command.

### Generate messages

```sh
go-codegen generate messages /path/to/contract-api.json --output /path/to/output.go --package-name mypackage
```

This will generate the Go code in the specified optional output directory, if not specified, it will generate the code in `msgs.go` in the current directory.
Package name is also optional, if not specified, it will use the name of the contract.

### Generate gRPC query client

```sh
go-codegen generate query-client /path/to/contract-api.json --output /path/to/output.go --package-name mypackage
```

This will generate the Go code in the specified optional output directory, if not specified, it will generate the code in `query.go` in the current directory. Package name is also optional, if not specified, it will use the name of the contract. The generated code depends on the generated messages, the [wasmd package](https://pkg.go.dev/github.com/CosmWasm/wasmd), and the [grpc package](https://pkg.go.dev/google.golang.org/grpc). You can install them by running `go get github.com/CosmWasm/wasmd@latest` (or `go get github.com/CosmWasm/wasmd@v0.50.0` for a specific version) and `go get google.golang.org/grpc@latest` (or `go get google.golang.org/grpc@v1.63.3` for a specific version).

### Generate interchaintest test suite

```sh
go-codegen interchaintest scaffold
```

This will launch an interactive prompt to guide you through the process of generating the test suite.
The scaffolded test suite will include a basic test and instructions on how to run it. This test suite will not contain any contract specific code, or tests, you will need to add them using `add-contract` command.

### Adding a contract to the interchaintest test suite

```sh
go-codegen interchaintest add-contract /path/to/contract-api.json --suite-dir /path/to/suite
```

This will add a contract to the test suite. The suite directory is the directory where the test suite is located. If not specified, it will use the current directory. The contract API JSON file is the same file that was used to generate the messages and query client.

Note that you can find a tutorial on how to use this tool's interchaintest feature [here](https://srdtrk.github.io/go-codegen/).

## Acknowledgements

The Go Gopher mascot was created by [Renee French](https://reneefrench.blogspot.com/) and is licensed under the [Creative Commons 4.0 Attribution License](https://creativecommons.org/licenses/by/4.0/).
I've modified the original image to include the logo of [`ts-codegen`](https://github.com/CosmWasm/ts-codegen).
