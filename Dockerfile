FROM golang:1.19-alpine AS builder
WORKDIR /src
RUN go build -o /build/keeper ./cmd/server/main.go

FROM alpine:3.16
COPY --from=builder /build/keeper /
USER nobody
CMD ["/keeper"]