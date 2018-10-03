FROM golang:1.11.1

ADD . /go/

WORKDIR /go

RUN go build squeal.go

CMD ["squeal"]