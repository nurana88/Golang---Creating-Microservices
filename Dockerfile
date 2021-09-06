FROM golang:latest

WORKDIR /app

COPY go.mod .

COPY go.sum .

COPY . .

ENV PORT 8000

RUN go build

CMD ["./microservices"]
