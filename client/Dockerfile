FROM golang:latest

WORKDIR /go/src/app

ADD go.mod go.mod
ADD go.sum go.sum

ADD . .

ENV CAS_SERVER_PORT 5000
ENV CAS_SERVER_PORT localhost:5000
ENV CACHE_SERVER_ADDR localhost:18001
ENV HIT_MISS_RATIO 0.9
ENV ATTACK_FREQ 3
ENV ATTACK_DURATION 10
ENV CONTENT_SIZE 25000



RUN go build -mod=vendor -o main .

CMD ["./main"]