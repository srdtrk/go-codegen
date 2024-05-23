---
title: Set Up Your Work Environment
sidebar_label: Set Up Your Work Environment
sidebar_position: 1
slug: /setup-env
---

import HighlightBox from '@site/src/components/HighlightBox';

# Set up your work environment

On this page, you can find helpful links to set up your work environment.

<HighlightBox type="info" title="Dependencies">

In this section, you can find all you need to install:

- [Git](https://git-scm.com/)
- [Go v1.22.3](https://go.dev/)
- [go-codegen v0.2.2](https://github.com/srdtrk/go-codegen)
- [Docker](https://www.docker.com/)
- [Rust and Cargo](https://www.rust-lang.org/)
- [just](https://just.systems/) (optional)

</HighlightBox>

<HighlightBox type="note" title="Note">

On a general note, it is advisable to prepare a separate project folder to keep all your Cosmos exercises.

</HighlightBox>

## Git

Install Git following the instructions on the [Git website](https://git-scm.com/). Test if Git is installed by running the following command:

```bash
git --version
```

## Go

Install the latest version of Go following the instructions on the [Go website](https://go.dev/). Test if Go is installed by running the following command:

```bash
go version
```

## go-codegen

Install go-codegen by following the instructions on the [README.md](https://github.com/srdtrk/go-codegen/blob/main/README.md) or by using the command below.

```bash
go install github.com/srdtrk/go-codegen@v0.2.2
```

Test if go-codegen is installed by running the following command:

```bash
go-codegen version
```

## Docker

Install Docker following the instructions on the [Docker documentation](https://docs.docker.com/get-docker/). Test if Docker is installed by running the following command:

```bash
docker --version
```

## Rust and Cargo

Install Rust and Cargo following the instructions on the [Rust book](https://doc.rust-lang.org/book/ch01-01-installation.html). Test if Rust is installed by running the following commands:

```bash
cargo --version
```

## just

<HighlightBox type="info" title="Just">

[Just](https://just.systems/) is a handy `make` alternative written in Rust. It is optional but recommended for this tutorial. Instead of writing a `Makefile`, you can write a `justfile` to automate your tasks.

</HighlightBox>

Install just following the instructions on the [just manual](https://just.systems/man/en/chapter_4.html) or by using the command below.

```bash
cargo install just
```

Test if just is installed by running the following command:

```bash
just --version
```

<HighlightBox type="tip" title="Just">

If you don't want to install `just`, you can simply run the commands in the `justfile` manually.

</HighlightBox>
