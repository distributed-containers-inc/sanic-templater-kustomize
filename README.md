# Sanic Templater for Kustomize
This is a [sanic](https://github.com/distributed-containers-inc/sanic) templater for [kustomize](https://github.com/kubernetes-sigs/kustomize)

## Structure
This templater expects your sanic.yaml's environment keys to match the directories in the "overlays" folder. I.e., `sanic.yaml` contains:
```
environments:
  dev:
    (the sanic configuration for dev)
  prod:
    (the sanic configuration for prod)
deploy:
  templaterImage: distributedcontainers/templater-kustomize
```

It expects your deploy folder to look like the following:
```
deploy
└── in
    ├── base
    │   ├── kustomization.yaml
    │   └── (other base yamls...)
    └── overlays
        ├── dev
        │   ├── kustomization.yaml
        │   └── (other patches...)
        └── prod
            ├── kustomization.yaml
            └── (other patches...)
```

Note that the directories in the overlays/ directory have a one-to-one coorespondance with the keys in the sanic.yaml environments: block.

## Example
To try the example in isolation, run the following:
```
docker build . -t distributedcontainers/templater-kustomize
docker run --rm \
    -v $(pwd)/deploy/in:/in:ro \
    -v $(pwd)/deploy/out:out:rw \
    -e SANIC_ENV=dev \
    distributedcontainers/templater-kustomize 
```
