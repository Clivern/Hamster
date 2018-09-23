FROM golang:latest

ADD . /go/

WORKDIR /go

RUN go build squeal.go

CMD ["squeal.go"]