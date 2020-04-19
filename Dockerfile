FROM golang:1.14

WORKDIR /app

COPY . .

RUN go build -o /go/bin/prettyjsonbot

ENTRYPOINT ["/go/bin/prettyjsonbot"]