# Test WebApp (Java implementation)
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

## How to build
The following command will help to build the app:
```bash
./gradlew build
```
If all went well the new JAR-file will be created inside of `webapp-spring-boot/build/libs` folder.
