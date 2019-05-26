# Example of test Webapp and its deploy.
## Destionation
The `Webapp` exposes REST-API through which an user can gather information about:
1. number of available CPUs for app
2. number of IPv4 addresses and name of interfaces to which these IPs have assigned
3. current date
4. current time
## How to use it
In order to deploy and to try it, a number of steps have to be done:
1. Install [VirtualBox](https://www.virtualbox.org/wiki/Downloads) and [Vagrant](https://www.vagrantup.com/downloads.html) environments for Linux
2. Clone this repository
3. From file-system directory of cloned repo, launch command `vagrant up` and wait until the machine will be provisioned
4. Launch the commands:
```
for path in cpu ip date time; do 
curl -s --insecure https://localhost:8443/$path | python -m json.tool; 
done
```
and watch the output
