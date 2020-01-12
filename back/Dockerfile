FROM golang:1.13.5-alpine

WORKDIR /app

COPY . .

RUN go build -o main .

CMD ["./main"]
