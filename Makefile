.PHONY: *

.DEFAULT_GOAL := setup

.SILENT:

include $(shell test -f .env.local || cp .env.local.dist .env.local && echo .env.local)

export DOCKER_BUILDKIT=1
export TURBO_ENIGMA_IMAGE ?= codelicia/turbo-enigma:latest

setup: image/build test/unit analysis/revive app/run

image/build:
	docker build -t ${TURBO_ENIGMA_IMAGE} .

app/run:
	docker run -it --rm -p 8000:80 \
        -e SLACK_WEBHOOK_URL="${SLACK_WEBHOOK_URL}" \
        -e NOTIFICATION_RULES="${NOTIFICATION_RULES}" \
        ${TURBO_ENIGMA_IMAGE}

app/rerun: image/build app/run

style/fix:
	go fmt ./...

test/unit:
	go test ./...

coverage/generate:
	go test ./... -coverprofile=coverage.out

coverage/view:
	go tool cover -html=coverage.out

analysis/revive:
	docker run --rm -it \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v  "${PWD}":"${PWD}" \
		-w "${PWD}" \
		morphy/revive-action