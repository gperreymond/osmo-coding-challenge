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

You need to finish the mongo installation, not so long don't worry :)

```sh
https://mongo-express.docker.localhost
```

- Create a new database: __osmo__
- Create a new collection: __eventstore__

__Infrastructure is ready!__

## Listing of infrastructure components

#### Commons

- traefik: https://traefik.docker.localhost

#### Backends

- rabbitmq: https://portainer.docker.localhost
- nats: https://nats.docker.localhost
- mongodb: https://mongo-express.docker.localhost
- logstash
- elasticsearch

#### Frontends

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
