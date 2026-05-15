FROM golang:1.24 AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o api .

FROM scratch

COPY --from=builder /app/api .

EXPOSE 8080

CMD ["./api"]