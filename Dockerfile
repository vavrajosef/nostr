# TODO: add nodejs build step

# TODO: change to golang build step
FROM golang

WORKDIR /go/src

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -v -o /go/bin ${PWD}/cmd

CMD ["/go/bin"]

# TODO: add run step