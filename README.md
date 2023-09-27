# test-stub-golang
Golang Test Stub.

Simple.


## Run the program

```bash
chmod +x "test-stub-golang"
./test-stub-golang
```

## Access test endpoint Web UI

```text
http://hostname:10100/endpoint/ui
```

## Access test endpoint via REST API

```bash
curl "http://hostname:10100/endpoint"
```

## Change response delay time (in milliseconds)

```bash
curl \
    --request POST \
    --data-binary '{"delayMs":555}' \
    --header 'Content-Type: application/json' \
    "http://hostname:10100/endpoint/api/set-delay"
```
