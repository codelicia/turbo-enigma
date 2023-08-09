# syntax=docker/dockerfile:1.2
FROM golang:1.21-alpine AS builder
WORKDIR /turbo-enigma

COPY . ./
RUN go mod download \
    && CGO_ENABLED=0 GOOS=linux go build -o bin/turbo-enigma

FROM scratch

COPY --from=builder /turbo-enigma/bin/. /.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENV HTTP_PORT=80
ENV NOTIFICATION_RULES="[{\"channel\":\"#random\", \"labels\":[\"just-testing\"]}]"
ENV REACTION_RULES="[{\"action\":\"approved\", \"reaction\":\"thumbsup\"}]"
ENV MESSAGE="New Merge Request Created"
ENV SLACK_AVATAR_URL="https://avatars.githubusercontent.com/u/46966179?s=200&v=4"
ENV SLACK_USERNAME="codelicia/turbo-enigma"
ENV SLACK_TOKEN="slack-token"

ENTRYPOINT ["/turbo-enigma"]
