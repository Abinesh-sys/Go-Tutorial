FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLE=0 GOOS=linux go build -o go-tutorial .
FROM scratch
COPY --from=builder /app/go-tutorial .
CMD ["./go-tutorial"]