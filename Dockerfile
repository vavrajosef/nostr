# Builder step for Angular client
# - TBD
FROM node:alpine as client_builder
WORKDIR /
COPY . .
RUN npm i -g @angular/cli
RUN npm ci
RUN npm run build -w internal/client

# Builder step for Hugo documentation
# - TBD
FROM node:alpine as docs_builder
WORKDIR /
RUN apk add --no-cache hugo
COPY . .
RUN npm ci
RUN npm run build -w internal/docs
RUN hugo -s internal/docs

# Builder step for Golang application
# - TBD
FROM golang:latest as cmd_builder
WORKDIR /go/src
COPY --from=client_builder /internal/client/dist /go/src/internal/client/dist
COPY --from=docs_builder /internal/docs/public /go/src/internal/docs/public
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -v -o /go/bin/nostr ${PWD}/internal/cmd

# Final step for creating a minimalist and secure runtime environment
# - TBD
FROM gcr.io/distroless/base
COPY --from=cmd_builder /go/bin/nostr /nostr
CMD ["/nostr"]
