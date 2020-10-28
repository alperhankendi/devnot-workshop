FROM golang:1.14-alpine AS builder
RUN mkdir -p /go/src/github.com/alperhankendi/ws01
RUN CGO_ENABLED=0
RUN GOOS=linux

ENV GOPATH /go
WORKDIR /go/src/github.com/alperhankendi/ws01

ADD go.mod /go/github.com/alperhankendi/ws01
ADD go.sum /go/github.com/alperhankendi/ws01
RUN go mod download
ADD . /go/src/github.com/alperhankendi/ws01

#RUN go get
RUN go build

FROM alpine

RUN mkdir -p /app

COPY --from=builder /go/src/github.com/alperhankendi/ws01 /app/
RUN chmod +x /app/demo
WORKDIR /app
ENTRYPOINT ["/app/demo"]