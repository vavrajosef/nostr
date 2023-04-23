# Dockerfile for building a NodeJS, Hugo, and Golang application

# Builder step for Angular client
# - Set the base image to node:latest
# - Set the working directory to the root
# - Copy package configuration files
# - Install NodeJS dependencies
# - Install Angular CLI globally
# - Copy all source files
# - Build the internal client using Angular
FROM node:latest as client_builder
WORKDIR /
COPY package.json package-lock.json ./
RUN npm ci
RUN npm install -g @angular/cli @angular-devkit/build-angular
COPY . .
RUN npm run build -w internal/client

# Builder step for Hugo documentation
# - Set the base image to node:alpine
# - Set the working directory to the root
# - Install Hugo
# - Copy package configuration files for NodeJS dependencies
# - Install NodeJS dependencies
# - Copy all source files
# - Build the NPM docs package
# - Build the Hugo documentation
FROM node:latest as docs_builder
WORKDIR /
RUN apk add --no-cache hugo
COPY package.json package-lock.json ./
RUN npm ci
COPY . .
RUN npm run build -w internal/docs
RUN hugo -s internal/docs

# Builder step for Golang application
# - Set the base image to golang
# - Set the working directory to /go/src
# - Copy the built assets from previous steps (client_builder and docs_builder)
# - Copy all source files
# - Download Golang dependencies
# - Install Golang dependencies
# - Build the Golang application binary
FROM golang as cmd_builder
WORKDIR /go/src
COPY --from=client_builder /internal/client/dist /go/src/internal/client/dist
COPY --from=docs_builder /internal/docs/public /go/src/internal/docs/public
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -v -o /go/bin/nostr ${PWD}/internal/cmd

# Final step for creating a minimalist and secure runtime environment
# - Set the base image to gcr.io/distroless/base
# - Copy the built Golang binary from the previous step
# - Set the command to run the application
FROM gcr.io/distroless/base
COPY --from=cmd_builder /go/bin/nostr /nostr
CMD ["/nostr"]
