FROM golang:1.23.2 AS Builder
COPY ../.. /app
WORKDIR /app
RUN go build -o eg .

FROM alpine:latest
COPY --from=Builder /app/eg /app/eg
CMD ["/app/eg"]
