FROM golang:latest

WORKDIR WORKDIR /go/src/github.com/MicBun/github.com/MicBun/ggithub.com/MicBun/go-file-manager

COPY . .

RUN go get -d -v ./...

RUN go build -o go-file-manager .

EXPOSE 8080

ENTRYPOINT ["./go-file-manager"]