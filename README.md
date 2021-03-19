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
