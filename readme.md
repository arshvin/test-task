# Example of test Webapp and its deploy.
## Purpose
The `Webapp` exposes REST-API through which an user can gather information about:
1. number of available CPUs for app
2. number of IPv4 addresses and name of interfaces to which these IPs have assigned
3. current date
4. current time

## How to use it
In order to deploy and to try it, a number of steps have to be done:
1. Install [VirtualBox](https://www.virtualbox.org/wiki/Downloads) and [Vagrant](https://www.vagrantup.com/downloads.html) environments for Linux
2. Clone this repository from shell
3. Go to the cloned repository directory, launch command `vagrant up` and wait until the machine will be provisioned
4. Invoke the commands:

```bash
for item in cpu ip date time; do
    curl -s -L --insecure http://localhost:8080/$item | python -m json.tool;
done
```
and watch the output
