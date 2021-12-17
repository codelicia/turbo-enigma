.PHONY: *

.SILENT:

include $(shell test -f .env.local || cp .env.local.dist .env.local && echo .env.local)

export TURBO_ENIGMA_IMAGE ?= codelicia/turbo-enigma:latest

image/build:
	docker build -t ${TURBO_ENIGMA_IMAGE} .

app/run:
	docker run -it --rm -p 8000:80 \
        -e SLACK_WEBHOOK_URL="${SLACK_WEBHOOK_URL}" \
        -e NOTIFICATION_RULES="${NOTIFICATION_RULES}" \
        ${TURBO_ENIGMA_IMAGE}

app/rerun: image/build app/run

test/unit:
	go test ./...

coverage/generate:
	go test ./... -coverprofile=coverage.out

coverage/view:
	go tool cover -html=coverage.out

echoes:
	env