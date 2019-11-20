FROM golang:latest

WORKDIR /go/src/app

ADD go.mod go.mod
ADD go.sum go.sum

ADD . .

RUN go build -mod=vendor -o main .

CMD ["./main"]