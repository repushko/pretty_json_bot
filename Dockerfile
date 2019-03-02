FROM golang:1.11

WORKDIR $GOPATH/src/github.com/grisme/prettyjsonbot

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

CMD ["prettyjsonbot"]