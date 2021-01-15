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

Create a requisition and add a node that represents the generator:

```
docker-compose exec opennms bash
```

Then,

```bash
GENERATOR_IP=$(getent hosts generator | awk '{print $1}')
provision.pl requisition add Docker
provision.pl node add Docker generator generator
provision.pl interface add Docker generator $GENERATOR_IP
provision.pl interface set Docker generator $GENERATOR_IP snmp-primary N
provision.pl requisition import Docker
exit
```

## Visualization

There is ready to use dashboard available in Grafana, accessible through `http://localhost:3000`, called "Graphite Playground".

