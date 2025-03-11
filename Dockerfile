FROM golang:1.23.2 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o eg .

FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/eg /app/eg
EXPOSE 8080
RUN chmod +x /app/eg
CMD ["/app/eg"]