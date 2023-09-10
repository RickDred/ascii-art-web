FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main ./cmd/

FROM alpine
WORKDIR /app
COPY --from=builder /app/main ./
COPY /internal/ /app/internal/
COPY /pkg/ /app/pkg/
COPY /static/ /app/static/
COPY /templates/ /app/templates/

CMD ["./main"]