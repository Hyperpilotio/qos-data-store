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
curl -H "Content-Type: application/json" -XPOST localhost:7781/v1/apps/goddd/metrics/qos -d '{"value": 1.45}'
# Query value
curl http://localhost:7781/v1/metrics
```
# Switch on
curl -XPOST "http://localhost:7781/v1/switch/on

# Switch off
curl -XPOST "http://localhost:7781/v1/switch/off
