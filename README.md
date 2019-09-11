# OSMO Coding Challenge

## Architecture

Based on a solid experience with __moleculer__ for my current company.
- https://moleculer.services/

I decided to try the golang version (currently in progress), but fully compatible with the test.
- https://moleculer-go-site.herokuapp.com/

Due to the lack of time, I implemented a very very light version of CQRS/ES.
To be honest it's more than a start, than a full version, but you know what I mean.

I just want to demonstrate how I will make the architecture if I had to begin such a big project.

I don't use an __Observer Pattern__ because when I want to create a new achievment, I just need to replay past events.
And also because __Eventsourcing__ is a pretty good pattern in case of debug, data, etc.

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
