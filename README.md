# OSMO Coding Challenge

[![CircleCI](https://circleci.com/gh/gperreymond/osmo-coding-challenge.svg?style=shield)](https://circleci.com/gh/gperreymond/osmo-coding-challenge) [![Coverage Status](https://coveralls.io/repos/github/gperreymond/osmo-coding-challenge/badge.svg?branch=master&kill_cache=2)](https://coveralls.io/github/gperreymond/osmo-coding-challenge?branch=master)

## TODO

This list need to be added of course, I'm running out of time, but it's not a really hard work to do.

- Env vars, from docker (K8S) and/or config file
- Implements all the Actions to retrieve achievement (in go) ; I put all the __RethinkDB__ RQL, just need to be translated.

I will finish, the stuff!

## Architecture

Based on a solid experience with __moleculer__ used in my current company.
- https://moleculer.services/

I decided to try the golang version (currently in progress), but fully compatible with osmo test.
- https://moleculer-go-site.herokuapp.com/

Due to the lack of time, I implemented a very very light version of CQRS/ES ; With more time, in a real context of work, I will do more of course.

__Eventsourcing__ is better than __Observer Pattern__ for an achievement system.
Because when I want to create a new achievement, I just need to replay past events.
__Eventsourcing__ is also a pretty good pattern in case of debug, data, etc.

## Infrastructure

Read the README in __infra__ directory for infrastructure details.

__Nats__ is used as service discovery, and __rabbitmq__ if I need a state machine. This is a very good combo in microservices architectures.

```sh
cd infra
./start.sh
cd ..
```

## Logs, Dashboards, Stress tests

I don't have time to implements application logs for the test, but Moleculer give us metrics to monitor all the services/actions.
Grafana dashboards for moleculer are just awesome.

You can also monitor the go in runtime, but you need to configure grafana.  

```sh
https://grafana.docker.localhost/
```

Set up interaction between Grafana and InfluxDB Grafana (Grafana main page -> Top left corner -> Datasources -> Add new datasource):

![](https://i.stack.imgur.com/7o7VR.png)

Import dashboard #3242 from https://grafana.com (Grafana main page -> Top left corner -> Dashboard -> Import):

![](https://i.stack.imgur.com/ZyHlx.png)

Finally, launch your application: it will transmit runtime metrics to the contenerized Influxdb.

#### Real stress test on 100 games played auto
![stress test](osmo-stress-test-100.png?raw=true)

## Tests unit or integration ?

When I work with Moleculer and microservices I prefer to use a BDD approach and integration tests.

__BEFORE RUNNING TESTS YOU NEED TO INITILIAZE THE DATABASES__  
see HOW TO USE > INITILIAZING (next doc part)

- https://github.com/onsi/ginkgo

```sh
go get -u github.com/onsi/ginkgo/ginkgo  # installs the ginkgo CLI
go get -u github.com/onsi/gomega/...     # fetches the matcher library

ginkgo -r -v -cover
```

## How to use

#### Build the binary

```sh
go get -u github.com/golang/dep/cmd/dep
dep ensure
go build main.go
```

#### Initializing

Only one time stuff:

- Create __Postgres__ database
- Create __RethinkDB__ database
- Add the players in postgres.

```sh
./main --initialize
```

#### Run the all the services together

```sh
./main --start-services
```

__Nota__   
We could add some arguments to start only selected microservices. This is a better approch for scalabilty.

## Games

On another CLI:

You can play 100 games in a row, around 3 minutes execution per 100 games.

```sh
./playGames.sh
```
Or play only one:

```sh
./main --play-game
```

## Achievements

- https://rethinkdb.docker.localhost

![stress test](infra/rethinkdb.png?raw=true)

__RethinkDB__ is a powerfull database for map reducers, aggregate data, etc.  
For the test I used this one, but I look closely __Couchbase__ to replace it.

To execute the query and play with events, you can go to __Data Explorer__ on the admin page.

#### “Sharpshooter Award”
A user receives this for landing 75% of their attacks, assuming they have at least attacked once.

The __RethinkDB__ query:

```js
```

####  “Big Winner” Award
A user receives this for having 200 wins

The __RethinkDB__ query:
```js
r.db("osmo").table("eventstore")
.filter({
  "AggregateID":  "<INSERT AggregateID of a player>",
  "AggregateType":  "Player",
  "EventType": "GameWonFinished"
})
.count()
.gt(200);
```

#### “Veteran” Award
A user receives this for playing more than 1000 games in their lifetime.

The __RethinkDB__ query:
```js
r.db("osmo").table("eventstore")
.filter({
  "AggregateID":  "<INSERT AggregateID of a player>",
  "AggregateType":  "Player",
  "EventType": "GameStarted"
})
.count()
.ge(1000);
```

#### “Bruiser” Award
A user receives this for doing more than 500 points of damage in one game

This is one is little bit tricky, you have to map/reduce on "TotalAmountOfDamageDone" per "Game", and get the max :)
The __RethinkDB__ query:

```js
r.db("osmo").table("eventstore")
.filter({
  "AggregateID":  "<INSERT AggregateID of a player>",
  "AggregateType":  "Player",
  "EventType":  "TotalAmountOfDamageDoneUpdated"
}).map(function(val) {
  return val("Data")
})
.max("TotalAmountOfDamageDone")
.getField("TotalAmountOfDamageDone")
.ge(500);
```

As example, I show here the "translated" ReQL in GO:

```go
res, _ := r.Table("eventstore").Filter(Filter{
  AggregateID:   aggregateID,
  AggregateType: aggregateType,
  EventType:     eventType,
}).Map(func(row r.Term) interface{} {
  return row.Field("Data")
}).Max("TotalAmountOfDamageDone").Field("TotalAmountOfDamageDone").Ge(500).Run(session)
```

How to control if a player have this achievement:

```sh
go run main.go --bruiser-award <INSERT AggregateID of a player>
```

But you can also do it by calling the microservice:

```go
// Control BruiserAward for a player
res := <-bkr.Call("Achievement.ControlBruiserAward", payload.New(map[string]string{
  "AggregateID":   aggregateID,
}))
```
