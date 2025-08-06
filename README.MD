# Example of test Webapp and its deploy.
## Purpose
The `Webapp` exposes REST-API through which an user can gather information about:
1. number of available CPUs for app (`/cpu` path)
2. number of IPv4 addresses and name of interfaces to which these IPs have assigned (`/ip` path)
3. current date (`/date` path)
4. current time (`/time` path)

It's supposed that this app will have implementations in few languages. Currently there are following language implementation:
1. [java](./webapp-spring-boot/README.MD)
2. [golang](./webapp-golang/README.MD)

## How to use it
To deploy and for trying it, a number of steps have to be done:
1. Install [VirtualBox](https://www.virtualbox.org/wiki/Downloads) and [Vagrant](https://www.vagrantup.com/downloads.html) environments for Linux
2. Clone this repository from shell
3. Go to the cloned repository directory, launch command `vagrant up` and wait until the machine will be provisioned
4. Invoke the commands:

```bash
for item in cpu ip date time; do
    echo "\nRequesting of item $item:\n";
    curl -s -L --insecure http://localhost:8080/$item | python -m json.tool;
done
```

and watch the output. You shall see something like:

```bash

Requesting of item cpu:

{
    "type": "CPU_AMOUNT",
    "payload": "2"
}

Requesting of item ip:

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

Requesting of item date:

{
    "type": "DATE_NOW",
    "payload": "2024 10 08"
}

Requesting of item time:

{
    "type": "TIME_NOW",
    "payload": "02:57:30"
}
```

5. Once it's done, to remove the VM, the command is helpful:

```bash
vagrant destroy default
```
