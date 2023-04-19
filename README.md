# @go-nostr/go-nostr

[![Angular](https://github.com/go-nostr/go-nostr/actions/workflows/angular.yml/badge.svg)](https://github.com/go-nostr/go-nostr/actions/workflows/angular.yml)

[![Docker](https://github.com/go-nostr/go-nostr/actions/workflows/docker.yml/badge.svg)](https://github.com/go-nostr/go-nostr/actions/workflows/docker.yml)

[![Go](https://github.com/go-nostr/go-nostr/actions/workflows/go.yml/badge.svg)](https://github.com/go-nostr/go-nostr/actions/workflows/go.yml)

[![Hugo](https://github.com/go-nostr/go-nostr/actions/workflows/hugo.yml/badge.svg)](https://github.com/go-nostr/go-nostr/actions/workflows/hugo.yml)

An all-in-one Go/Hugo/Angular monorepo for the nostr protocol, featuring comprehensive documentation, a nostr relay executable, reusable Go packages, and an Angular web client for relay administration.

## Overview

This monorepo contains various subprojects related to the nostr protocol, including:

- `docs`: Documentation pages built using GitHub Pages and Hugo
- `cmd`: Executable used for running a nostr relay
- `pkg`: Various nostr/Go packages for use by external applications
- `web`: Angular web client for administering the nostr relay (maintained in `cmd/`)

## Getting Started

### Prerequisites

Before starting development, make sure to install the following dependencies:

- [Install Angular](https://angular.io/guide/setup-local)
- [Install Docker](https://docs.docker.com/engine/install/)
- [Install Go](https://go.dev/doc/install)
- [Install Hugo](https://gohugo.io/installation/)
- [Install NPM and NodeJS](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm)
- [Install Wire](https://github.com/google/wire#installing)

### Installation

> Installation _should_ include a list of step-by-step instructions for installing the application.

TBD

## Usage

> Usage _should_ include sections related to running and testing the application.

TBD

## Contributing

For information on contributing to this project, please see the [CODE_OF_CONDUCT](./CODE_OF_CONDUCT.md).

### Snippets

#### Build NPM dependencies

```shell
npm run build -ws
```

#### Format Go files

```shell
go fmt ./...
```

#### Generate Go dependencies

```shell
go generate ./...
```

#### Install NPM dependencies

```shell
npm i -ws
```

#### Run Go tests

```shell
go test ./...
```

#### Run NPM tests

```shell
npm test -ws
```

## Deployment

Deployment information will be added as the project matures.

## Built-with

This project is built with the following core technologies:

- [Angular](https://angular.io/)
- [Docker](https://docker.com/)
- [Go](https://go.dev/)
- [Hugo](https://gohugo.io/)
- [NPM](https://www.npmjs.com/)
- [NodeJS](https://nodejs.org/en)

## Acknowledgements

- [@peaceiris](https://github.com/peaceiris)

## Authors

- [@chantzlarge](https://github.com/chantzlarge)

## Versioning

Versioning information will be added as the project matures.

## License

This project is licensed under the terms of the [LICENSE](./LICENSE).

## References

For additional materials helpful for contributors or users, please see:

- [@decred/dcrd](https://github.com/decred/dcrd)
- [@golang-standards/project-layout](https://github.com/golang-standards/project-layout)
- [@nostr-protocol/nips](https://github.com/nostr-protocol/nips)
- [@nhooyr/websocket](https://github.com/nhooyr/websocket)
