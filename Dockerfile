FROM golang:1.23.2
COPY . /app
WORKDIR /app
RUN go build -o eg .
CMD ["/app/eg"]