# Osmo Development Bootstrap

Nerd stuff - Launch Osmo production on your local station w/ Docker.  
First time ELK stack is created, it could be long, because these three components have to be build from docker files.

## Prepare your computer

#### Troubles with elasticsearch

- sudo sysctl -w vm.max_map_count=262144
- sudo sysctl -w fs.file-max=65536

#### Docker engines

Having __docker__ and __docker-compose__ ready to use.

## First time running

You need to finish the postgres installation, not so long don't worry :)

- Create a new database: __osmo__

You need to finish the rethinkdb installation, not so long don't worry :)

```sh
https://rethinkdb.docker.localhost/
```

- Create a new database: __osmo__
- Create a new table: __eventstore__

![rethinkdb](rethinkdb.png?raw=true)

__Infrastructure is ready!__

## Listing of infrastructure components

#### Commons

- traefik: https://traefik.docker.localhost

#### Backends

- rabbitmq: https://portainer.docker.localhost
- nats: https://nats.docker.localhost
- rethinkdb: https://rethinkdb.docker.localhost
- logstash
- elasticsearch

#### Frontends

- grafana : https://grafana.docker.localhost
- kibana (elk webui): https://kibana.docker.localhost

## Usages

#### Running the infrastructure
Start all components.

```sh
./start.sh
```

#### Stopping the infrastructure
Stop and remove all components.

```sh
./stop.sh
```
