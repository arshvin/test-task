# Test WebApp (Golang implementation)
## Description
It should have the same route list and return the same JSON objects as the app that was described in the parent [README.MD](../README.MD) file:

```bash
GET /cpu ->
{
    "type": "CPU_AMOUNT",
    "payload": "2"
}

GET /ip ->
{
    "type": "IP_V4",
    "count": "EVEN",
    "payload": [
        {
            "name": "eth0",
            "ip": "10.0.2.15"
        },
        {
            "name": "lo",
            "ip": "127.0.0.1"
        }
    ]
}

GET /date ->
{
    "type": "DATE_NOW",
    "payload": "2024 10 08"
}

GET /time ->
{
    "type": "TIME_NOW",
    "payload": "02:57:30"
}
```

Also, additionally it has capability to print available route list if sent request ot `/`:
```bash
GET / ->
[
  {
    "method": "GET",
    "path": "/cpu",
    "name": "cpu-amount"
  },
  {
    "method": "GET",
    "path": "/ip",
    "name": "ip-addr-amount"
  },
  {
    "method": "GET",
    "path": "/date",
    "name": "date-now"
  },
  {
    "method": "GET",
    "path": "/time",
    "name": "time-amount"
  },
  {
    "method": "GET",
    "path": "/",
    "name": "available-routes"
  }
]
```

## How to build
It's required to have installed `mage` tool, which will help to build the main app. To reach out of it, the following commands will help:
```bash
go install github.com/magefile/mage@latest

mage
```
If all went well the new binary file will be created inside of `webapp-golang` folder.
