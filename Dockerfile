FROM golang:1.26-alpine3.24 AS builder

WORKDIR /app

COPY . /app

RUN go build -o /app/hopscotch cmd/hopscotch/main.go

FROM alpine:3.24

WORKDIR /app

COPY --from=builder /app/hopscotch /app/hopscotch

ENV HOPSCOTCH_HOST="0.0.0.0"
ENV HOPSCOTCH_PORT="80"
EXPOSE 80

ENTRYPOINT ["/app/hopscotch"]
