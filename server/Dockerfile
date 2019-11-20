FROM golang:latest

WORKDIR /go/src/app

ADD go.mod go.mod
ADD go.sum go.sum

ADD . .

RUN go build -mod=vendor -o main .

ENV CAS_SERVER_PORT 5000

EXPOSE 5000

CMD ["./main"]