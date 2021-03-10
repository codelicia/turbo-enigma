FROM golang:1.16-alpine
WORKDIR /go/src/app
ADD main.go .
ENV HTTP_PORT=80
ENV MESSAGE="New Merge Request Created"
ENV SLACK_WEBHOOK_URL="http://turboenigma.localhost"
RUN go build main.go
CMD ["./main"]
