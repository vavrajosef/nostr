# NodeJS dependency installer step for workspace
FROM node:latest as node_builder
WORKDIR /
COPY package.json package-lock.json ./
RUN npm ci

# NodeJS builder step for workspace
COPY . .
RUN npm run build -ws

# Hugo builder step for docs/
FROM klakegg/hugo:latest as hugo_builder
WORKDIR /docs
COPY docs/ .
RUN hugo

# Golang builder step
FROM golang as go_builder
WORKDIR /go/src
COPY --from=node_builder /web/dist /web/dist
COPY --from=hugo_builder /docs/public /docs/public
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -v -o /go/bin/nostr ${PWD}/cmd

# Minimalist, secure Go run step
FROM gcr.io/distroless/base
COPY --from=go_builder /go/bin/nostr /nostr
CMD ["/nostr"]
