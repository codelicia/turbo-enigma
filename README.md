Turbo Enigma ⚡️🔋
=================

Environment
-----------

The application will need the Slack webhook URL. you can place it in the `.env.local`. 
if the file does not exist, you can duplicate it from `.env.local.dist`, otherwise it will get created the first time you run `make` 
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

```sh
$ make image/build
```

Run
---

```sh
$ make app/run
```

Testing
-------

```
curl localhost:8000 -d @payload/merge_request-open-just-testing.json
```

Unit tests
----------

To run the tests locally, run the following command:

```sh
$ make test/unit
```

If you want to see how covered the project is, you can run the following command to get coverage report
```sh
$ make coverage/generate
```

Once the above has been run, it's time to see it in your browser. The following command will open a new tab in your browser with the code coverage.

```sh
$ make coverage/view
```
