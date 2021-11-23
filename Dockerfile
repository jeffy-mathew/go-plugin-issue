
FROM golang:1.15.8-stretch
LABEL description="driver for plugin issue reproduction"

WORKDIR /driver

COPY src/driver .

RUN go build -o main main.go

ENTRYPOINT ["/driver/main"]
