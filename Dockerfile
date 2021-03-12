FROM golang:1.16-alpine
WORKDIR /go/src/app
ADD . .
ENV HTTP_PORT=80
ENV MESSAGE="New Merge Request Created"
ENV MERGE_REQUEST_LABEL="just-testing"
ENV SLACK_WEBHOOK_URL="http://turboenigma.localhost"
RUN go build
CMD ["./turboenigma"]
