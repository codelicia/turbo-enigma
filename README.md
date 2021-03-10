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
helm upgrade my-enigma . --set slack.webhookUrl=$SLACK_WEBHOOK_URL
```

Build
-----

```
docker build -t turbo-enigma .
```

Run
---

```
docker run -it --rm -p 8000:80 -e SLACK_WEBHOOK_URL=$SLACK_WEBHOOK_URL turbo-enigma
```

Testing
-------

```
curl localhost:8000 -d @payload.yaml
```
