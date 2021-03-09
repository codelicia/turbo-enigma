Turbo Enigma ⚡️🔋
=================

Build
-----

```
docker build -t turbo-enigma .
```

Environment
-----------

```
export WEBHOOK_URL="https://find-me-on.gitlab.com"
```

Run
---

```
docker run -it --rm -p 8000:80 -e WEBHOOK_URL=$WEBHOOK_URL turbo-enigma
```

Testing
-------

```
curl localhost:8000 -d @payload.yaml
```
