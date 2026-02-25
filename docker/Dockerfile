FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o pi-agent .

FROM alpine:3.21
COPY --from=builder /app/pi-agent /pi-agent
ENTRYPOINT ["/pi-agent"]