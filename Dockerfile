# Builder step for Angular client
# - Use the official Node.js Alpine image as the base image
# - Set the working directory to the root directory
# - Copy package.json and package-lock.json to the container root
# - Copy internal/client/package.json to the container /internal/client directory
# - Install Angular CLI globally
# - Install dependencies defined in package.json using npm ci
# - Copy the rest of the application code to the container
# - Build the Angular client using the "build" script in package.json
FROM node:alpine as client_builder
WORKDIR /
COPY package.json package-lock.json ./
COPY internal/client/package.json /internal/client/
RUN npm i -g @angular/cli
RUN npm ci
COPY . .
RUN npm run build -w internal/client

# Builder step for Hugo documentation
# - Use the official Node.js Alpine image as the base image
# - Set the working directory to the root directory
# - Install Hugo via apk package manager
# - Copy package.json and package-lock.json to the container root
# - Copy internal/docs/package.json to the container /internal/docs directory
# - Install dependencies defined in package.json using npm ci
# - Copy the rest of the application code to the container
# - Build the documentation using the "build" script in package.json
# - Run Hugo to generate the static site for the documentation
FROM node:alpine as docs_builder
WORKDIR /
RUN apk add --no-cache hugo
COPY package.json package-lock.json ./
COPY internal/docs/package.json /internal/docs/
RUN npm ci
COPY . .
RUN npm run build -w internal/docs
RUN hugo -s internal/docs

# Builder step for Golang application
# - Use the latest official Golang image as the base image
# - Set the working directory to /go/src
# - Copy the built Angular client and Hugo documentation to the container
# - Copy the contents of the current directory to the container
# - Download and install the dependencies defined in go.mod
# - Install the Go application to /go/bin
# - Build the Go application binary named nostr
FROM golang:latest as cmd_builder
WORKDIR /go/src
COPY --from=client_builder /internal/client/dist /go/src/internal/client/dist
COPY --from=docs_builder /internal/docs/public /go/src/internal/docs/public
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -v -o /go/bin/nostr ${PWD}/internal/cmd

# Final step for creating a minimalist and secure runtime environment
# - Use the official Distroless base image as the base image
# - Copy the built nostr binary from the cmd_builder stage to the root directory
# - Set the command to run the nostr binary
FROM gcr.io/distroless/base
COPY --from=cmd_builder /go/bin/nostr /nostr
CMD ["/nostr"]
