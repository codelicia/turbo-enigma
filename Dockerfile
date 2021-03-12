FROM golang:1.16-alpine
WORKDIR /go/src/app
ADD . .
ENV HTTP_PORT=80
ENV MERGE_REQUEST_LABEL="just-testing"
ENV MESSAGE="New Merge Request Created"
ENV SLACK_AVATAR_URL="https://avatars.githubusercontent.com/u/46966179?s=200&v=4"
ENV SLACK_USERNAME="codelicia/turbo-enigma"
ENV SLACK_WEBHOOK_URL="http://turboenigma.localhost"
RUN go build
CMD ["./turboenigma"]
