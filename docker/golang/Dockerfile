FROM golang:latest

ENV SRC_DIR=/go/src/github.com/ono5/book-list

ENV GOBIN=/go/bin

WORKDIR $GOBIN

ADD ./api $SRC_DIR

RUN cd /go/src/;

# mysql のドライバ
RUN go get github.com/lib/pq && \
    go get github.com/gorilla/mux && \
    go get github.com/subosito/gotenv

RUN go install github.com/ono5/book-list/

ENTRYPOINT ["./book-list"]
