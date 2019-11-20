FROM golang:latest

WORKDIR /go/src/app

ADD go.mod go.mod
ADD go.sum go.sum

ADD . .

RUN go build -mod=vendor -o main .

ENV CAS_SERVER_ADDR=localhost:5000
ENV PEER_PORT=8001
ENV LISTEN_PORT=18001

EXPOSE 18001
EXPOSE 18555

CMD ["./main"]