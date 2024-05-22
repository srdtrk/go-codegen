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
- [Docker](TODO)

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

Install go-codegen by following the instructions on the [README.md](TODO) or by using the command below.

```bash
go install github.com/srdtrk/go-codegen@v0.2.2
```

Test if go-codegen is installed by running the following command:

```bash
go-codegen version
```

## Docker

TODO
