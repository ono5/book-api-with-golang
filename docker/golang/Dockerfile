FROM golang:1.14

ENV SRC_DIR=/go/src/github.com/ono5/book-list
ENV GOBIN=/go/bin
ENV GO111MODULE=on
ENV GOPATH=

COPY go.mod go.sum ./
RUN go mod download

ADD ./api $SRC_DIR
WORKDIR $SRC_DIR

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o $GOBIN

RUN ls /go/bin

ENTRYPOINT ["/go/bin/book-list"]
