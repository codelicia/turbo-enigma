Turbo Enigma ⚡️🔋
=================

Environment
-----------

```
export SLACK_WEBHOOK_URL="https://find-me-on.slack.com"
```

Deploying
---------

```
helm upgrade --install my-enigma helm --set slack.webhookUrl=$SLACK_WEBHOOK_URL --set gitlab.mergeRequestLabel=codelicia-team
```

Build
-----

```
docker build -t turbo-enigma .
```

Run
---

```
docker run -it --rm -p 8000:80 \
    -e SLACK_WEBHOOK_URL=$SLACK_WEBHOOK_URL \
    -e MERGE_REQUEST_LABEL="codelicia-team" turbo-enigma
```

Testing
-------

```
curl localhost:8000 -d @pkg/payload/merge_request-open.json
```
