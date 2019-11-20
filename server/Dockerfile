FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

ENV CAS_SERVER_PORT 5000

EXPOSE 5000

CMD ["./main"]