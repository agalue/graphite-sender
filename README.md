Graphite Playground with OpenNMS
====

## Requirements

All you need is having Docker up and running on your machine.

When using Docker for Mac or Docker for Windows, make sure to have at least 4GB of RAM and 2 CPUs assigned to the Docker VM.

## Build and Run

```bash
docker-compose up -d --build
```

## Configuration

* Find the IP address of the container associated with the generator:

```bash
docker-compose exec generator ip addr | grep inet | grep -v 127.0.0.1
```

For example:

```bash
inet 172.19.0.5/16 brd 172.19.255.255 scope global eth0
```

* Create a requisition and add a node that represents the generator:

```
docker-compose exec opennms bash
```

Then,

```bash
provision.pl requisition add Docker
provision.pl node add Docker generator generator
provision.pl interface add Docker generator 172.19.0.5
provision.pl interface set Docker generator 172.19.0.5 snmp-primary N
provision.pl requisition import Docker
```

> *WARNING*: Replace 127.19.0.5 with the correct container IP

Finally, exit the container.

## Grafana

* Enable Helm

* Create a Datasource for OpenNMS Performance, using `http://opennms:8980/opennms` as the URL.

* Import the provided Dashboard (assuming the generator was created as shown above).

