# @go-nostr/nostr

⚠️ **WARNING: This project is actively under development and no stable releases have been made yet.** ⚠️

Please be aware that the code and documentation may contain unfinished features or inconsistencies. We appreciate your interest in the project and encourage you to check back for updates on progress towards a stable release.

## Overview

An all-in-one Go/Hugo/Angular monorepo for the nostr protocol, featuring comprehensive documentation, a nostr relay executable, reusable Go packages, and an Angular web client for relay administration.

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

### API

Multiple protocols are used to support communication between clients and relays.

#### gRPC

TBD

#### HTTP

TBD

#### Websocket

TBD

### Docs

TBD

### Web

TBD

## Contributing

For information on contributing to this project, please see the [CODE_OF_CONDUCT](./CODE_OF_CONDUCT.md).

### Snippets

#### Build Docker image

```shell
docker build -t nostr .
```

#### Build Docker services

```shell
docker-compose build
```

#### Build NPM packages

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

#### Run Docker services

```shell
docker-compose up
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

- [@klakegg](https://github.com/klakegg)
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
- [@nostr-protocol/nips](https://github.com/nostr-protocol/nips)
- [@nhooyr/websocket](https://github.com/nhooyr/websocket)
