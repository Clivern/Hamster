FROM golang:1.11.1

RUN curl https://raw.githubusercontent.com/golang/dep/v0.5.0/install.sh | sh

RUN mkdir -p /go/src/github.com/clivern/hamster/

ADD . /go/src/github.com/clivern/hamster/

WORKDIR /go/src/github.com/clivern/hamster

RUN dep ensure

RUN go build -o hamster hamster.go

CMD ["./hamster"]