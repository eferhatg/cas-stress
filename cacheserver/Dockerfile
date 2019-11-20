FROM golang:latest

WORKDIR /go/src/app

ADD go.mod go.mod
ADD go.sum go.sum

ADD . .

RUN go build -mod=vendor -o main .

ENV CAS_SERVER_ADDR=localhost:5000
ENV PEER_PORT=6001
ENV LISTEN_PORT=16001

EXPOSE 16001

CMD ["./main"]