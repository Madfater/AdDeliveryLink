FROM golang:1.23.3-alpine3.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./

# 先 download dependies 以防 build 時每次都要 download
RUN go mod download

COPY . .

ENV GOCACHE=/root/.cache/go-build

RUN --mount=type=cache,target="/root/.cache/go-build" go build -o main

FROM alpine:3.20

WORKDIR /root

COPY --from=builder /app/main .

CMD ["/root/main"]