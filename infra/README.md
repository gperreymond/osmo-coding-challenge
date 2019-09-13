# Osmo Development Bootstrap

Nerd stuff - Launch Osmo production on your local station w/ Docker.

## Prepare your computer

#### Docker engines

Having __docker__ and __docker-compose__ ready to use.

## Listing of infrastructure components

#### Commons

- traefik: https://traefik.docker.localhost

#### Backends

- rabbitmq: https://portainer.docker.localhost
- nats: https://nats.docker.localhost
- rethinkdb: https://rethinkdb.docker.localhost
- postgres

#### Frontends

- grafana : https://grafana.docker.localhost

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
