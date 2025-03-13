FROM golang:alpine AS builder

ENV GOPROXY=https://goproxy.cn,direct

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . /app

RUN go mod tidy && go build -o app

FROM alpine:3.12

WORKDIR /app

COPY --from=builder /app .

EXPOSE 8080

CMD ["./app"]
