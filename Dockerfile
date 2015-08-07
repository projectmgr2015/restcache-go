FROM golang:1.5

ADD . /go
WORKDIR /go
RUN mkdir -p /go/pkg


RUN go-wrapper download

CMD ["go", "run", "main.go"]