FROM golang:1.11

WORKDIR /go/src/app

COPY . .

RUN go get github.com/julienschmidt/httprouter && \
    go get github.com/ma-null/NetInterface


CMD ["go",  "run", "server.go"]