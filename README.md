Turbo Enigma ⚡️🔋
=================

Build
-----

```
docker build -t turbo-enigma .
```

Run
---

```
docker run -it --rm -p 8000:80 turbo-enigma
```

Testing
-------

```
curl localhost:8000 -d payload.yaml
```
