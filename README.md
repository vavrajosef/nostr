# @go-nostr/nostr

⚠️ **WARNING: This project is actively under development and no stable releases have been made yet.** ⚠️

Please be aware that the code and documentation may contain unfinished features or inconsistencies. We appreciate your interest in the project and encourage you to check back for updates on progress towards a stable release.

## Overview

All-in-one Angular/Go/Hugo monorepo for the Nostr protocol, featuring comprehensive documentation, an all-in-one Go executable with an Angular web client and reusable packages.

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

### Build Docker image

Create a Docker image for the application by using the Dockerfile in the repository. This command will build the image and tag it with the name 'nostr'.

```shell
docker build -t nostr .
```

### Build Docker services

Build all the services defined in the docker-compose.yml file. This command will create Docker images for each service and store them locally.

```shell
docker-compose build
```

### Build NPM packages

Compile and build all NPM packages in the monorepo, considering their interdependencies. The -ws flag ensures that the build is performed across all workspaces.

```shell
npm run build -ws
```

### Format Go files

Format all Go source files in the repository by running the Go formatter. This ensures consistent coding style across the project.

```shell
go fmt ./...
```

### Generate Go dependencies

Automatically generate code for Go dependencies, such as mocks, based on the source files. This command should be executed whenever there are changes to the Go source files.

```shell
go generate ./...
```

### Install NPM dependencies

Install all NPM dependencies for the monorepo workspaces. The -ws flag ensures that the installation is performed across all workspaces.

```shell
npm i -ws
```

### Run Docker services

Start all the Docker services defined in the docker-compose.yml file. This command will run the containers and display the logs in the console.

```shell
docker-compose up
```

### Run Go tests

Execute all Go tests in the repository, including unit and integration tests. This command will also display a summary of the test results.

```shell
go test ./...
```

### Run NPM tests

Run all NPM tests in the monorepo, considering their interdependencies. The -ws flag ensures that the tests are executed across all workspaces.

```shell
npm test -ws
```

## Contributing

For information on contributing to this project, please see the [CODE_OF_CONDUCT](./CODE_OF_CONDUCT.md).

## Deployment

The deployment process is handled by GitHub Actions, with separate workflows for Angular, Docker, Go, and Hugo. When a push is made, the Angular workflow runs, building the application, running tests, and linting the code. The built Angular application is then archived into a dist.tar file, which is uploaded as a build artifact.

On pushes to the main branch, the Docker workflow builds and pushes a Docker image to the GitHub Container Registry (GHCR), supporting both amd64 and arm64 platforms. The Go workflow runs in parallel, checking Go code formatting, building NPM packages, and running Go tests. The Hugo workflow, also triggered on pushes to the main branch, builds NPM packages and generates the Hugo site, uploading the artifacts and deploying them to GitHub Pages. This process ensures the codebase remains up-to-date and deployable, making it easy for current and future contributors to understand the project's deployment process.

## Built-with

This project is built with the following core technologies:

- [Angular](https://angular.io/)
- [Docker](https://docker.com/)
- [Go](https://go.dev/)
- [Hugo](https://gohugo.io/)
- [NPM](https://www.npmjs.com/)
- [NodeJS](https://nodejs.org/en)

## Acknowledgements

- [@peaceiris](https://github.com/peaceiris) for the [Hugo GitHub Action](https://github.com/peaceiris/actions-hugo)

## Authors

- [@chantzlarge](https://github.com/chantzlarge)

## Versioning

This project uses [semantic versioning](https://semver.org) and [GitHub](https://docs.github.com/en/repositories/releasing-projects-on-github/managing-releases-in-a-repository) to publish new versions and manage releases.

## License

This project is licensed under the terms of the [LICENSE](./LICENSE).

## References

For additional materials helpful for contributors or users, please see:

- [@decred/dcrd](https://github.com/decred/dcrd)
- [@nostr-protocol/nips](https://github.com/nostr-protocol/nips)
- [@nhooyr/websocket](https://github.com/nhooyr/websocket)
