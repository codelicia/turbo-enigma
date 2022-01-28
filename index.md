## Helm Chart Repository

### How to use it

On the first time we have to add the repository:
```
helm repo add turbo-enigma https://codelicia.github.io/turbo-enigma/
helm install my-enigma turbo-enigma/turbo-enigma
```

To release a new versions repository must be updated first:
```
helm repo update
helm upgrade my-enigma turbo-enigma/turbo-enigma
```

⎈Happy Helming!⎈

[index.yaml](index.yaml) gererated by [chart-releaser-action](https://github.com/helm/chart-releaser-action).
