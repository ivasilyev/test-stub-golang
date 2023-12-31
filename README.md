# test-stub-golang
Golang Test Stub.

Simple.

## Get the program

```bash
curl -fsSL "$(
    curl -s "https://api.github.com/repos/ivasilyev/test-stub-golang/releases/latest" \
    | grep -E '\"browser_download_url\":.*\linux.*' \
    | sed -E 's/.*"([^"]+)".*/\1/' 
)" -o "test-stub-golang"
```

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

## Access program monitoring metrics

```bash
curl "http://hostname:10100/actuator/prometheus"
```
