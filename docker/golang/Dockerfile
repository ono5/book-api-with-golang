FROM golang:1.14

ENV SRC_DIR=/go/src/github.com/ono5/book-list

ADD ./api $SRC_DIR
WORKDIR $SRC_DIR

RUN go get github.com/lib/pq && \
    go get github.com/gorilla/mux && \
    go get github.com/gorilla/handlers

ENTRYPOINT ["go", "run", "main.go"]
