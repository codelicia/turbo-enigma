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
cp helm/values.yaml.dist helm/values.yaml
helm upgrade --install my-enigma helm --set slack.webhookUrl=$SLACK_WEBHOOK_URL
```

Redeploying
---------

```
vim helm/values.yaml # Adding new notification rules for instance
helm upgrade --install my-enigma helm --set slack.webhookUrl=$SLACK_WEBHOOK_URL
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
    -e NOTIFICATION_RULES='[{"channel":"#codelicia-team", "labels": ["Codelicia"]}]' turbo-enigma
```

Testing
-------

```
curl localhost:8000 -d @payload/merge_request-open.json
```

Unit tests
----------

To run the tests locally, run the following command:

```sh
$ go test ./... 
```

If you want to see how covered the project is, you can run the following command to get coverage report
```sh
$ go test ./... -coverprofile=coverage.out
```

Once the above has been run, it's time to see it in your browser. The following command will open a new tab in your browser with the code coverage.

```sh
$ go tool cover -html=coverage.out
```
