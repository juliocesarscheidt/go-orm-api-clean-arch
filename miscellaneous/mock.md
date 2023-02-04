# Mocking API with Prism

## Docs

> [https://meta.stoplight.io/docs/prism/f51bcc80a02db-installation](https://meta.stoplight.io/docs/prism/f51bcc80a02db-installation)

> [https://meta.stoplight.io/docs/prism/83dbbd75532cf-http-mocking](https://meta.stoplight.io/docs/prism/83dbbd75532cf-http-mocking)

## Install Prism CLI
```bash
curl -L https://raw.githack.com/stoplightio/prism/master/install | sh

# run proxy
prism proxy $(pwd)/openapi.yaml http://localhost:8000 -h 0.0.0.0 --errors
```

## Running with Docker

```bash
docker container run --init \
  --rm -d --name prism \
  -v $(pwd)/openapi.yaml:/opt/openapi/openapi.yaml:rw \
  -p 4010:4010 stoplight/prism:4 \
  mock -h 0.0.0.0 "/opt/openapi/openapi.yaml"

docker container logs -f --tail 100 prism

docker container rm -f prism
```
