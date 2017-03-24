# qos-data-store

## Getting started

This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`.

Running it then should be as simple as:

```console
$ make build
$ ./bin/qos-data-store
```

### Testing

``make test``

### Quick Start

```{shell}

# start server.
make dev
# or
docker run -d -p 7781:7781 hyperpilot/qos-data-store:latest

# Insert value
curl -XPOST "http://localhost:7781/v1/metrics/httprequest?val=100"
# Query value
curl -XGET"http://localhost:7781/v1/metrics/httprequest"
```
