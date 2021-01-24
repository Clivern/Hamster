FROM golang:1.15.7

RUN mkdir -p /go/src/github.com/clivern/hamster/

ADD . /go/src/github.com/clivern/hamster/

WORKDIR /go/src/github.com/clivern/hamster

RUN go build -o hamster hamster.go

EXPOSE 8080

CMD ["./hamster"]