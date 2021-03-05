FROM golang:1.16-alpine
WORKDIR /go/src/app
ADD main.go .
RUN go build main.go
CMD ["./main"]
