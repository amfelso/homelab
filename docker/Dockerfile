FROM --platform=$BUILDPLATFORM golang:1.24-alpine AS builder
ARG TARGETARCH
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH} go build -o pi-agent .

FROM alpine:3.21
COPY --from=builder /app/pi-agent /pi-agent
ENTRYPOINT ["/pi-agent"]
