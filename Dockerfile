FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY app/go.mod ./
COPY app/*.go ./

RUN go mod tidy
RUN go test ./...
RUN go build -o server .

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]