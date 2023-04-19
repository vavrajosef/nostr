# TODO: add nodejs build step

# TODO: change to golang build step
FROM golang

WORKDIR /go/src

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -v -o /go/bin/nostr ${PWD}/cmd

CMD ["/go/bin/nostr"]

# TODO: add run step