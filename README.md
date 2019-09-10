# OSMO Coding Challenge

## Infrastructure

Read the README in __infra__ directory for infrastructure details.

```sh
cd infra
./start.sh
cd ..
```

## Logs, Dashboards, Stress tests

You have in the boostrap an ELK stack, I don't have time to implements application logs for the test, but Moleculer give us metrics to monitor all the services/actions.
You can use grafana with a moleculer dashboard too.

You can monitor the go in runtime :

- Dashbord here : https://grafana.docker.localhost/
- Help to implement: https://stackoverflow.com/questions/24863164/how-to-analyze-golang-memory

#### Real stress test on 100 games played auto
![stress test](https://raw.githubusercontent.com/gperreymond/osmo-coding-challenge/master/osmo-stress-test-100.png?token=AAVKRIIT7G5HWAWXGRQNJGK5O634C)

## How to use

#### Build the binary

```sh
dep ensure
go build main.go
```

#### Run the all the services together

```sh
./main --start-services
```

__Nota__   
We could add some arguments to start only selected microservices. This is a better approch for scalabilty.

#### Initialize first players

```c
names := []string{"Thrall", "Rexxar", "Gul'Dan", "Malfurion", "Garrosh", "Uther", "Anduin", "Valeera", "Morgl", "Medivh"}
```

```sh
./main --initialize
```
