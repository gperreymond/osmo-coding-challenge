# OSMO Coding Challenge

## Architecture

Based on a solid experience with __moleculer__ for my current company.
- https://moleculer.services/

I decided to try the golang version (currently in progress), but fully compatible with the test.
- https://moleculer-go-site.herokuapp.com/

Due to the lack of time, I implemented a very very light version of CQRS/ES.  
With more time, in a real context of work, I will do more of course.

I just want to demonstrate how I will make the architecture if I had to begin such a big project.

I don't use an __Observer Pattern__ because when I want to create a new achievement, I just need to replay past events.
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

You can monitor the go in runtime, but you need to configure grafana.  

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

## How to use

#### Build the binary

```sh
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

## Achievements

__RethinkDB__ is a powerfull database for map reducers, aggregate data, etc.  
For the test I used this one, at the moment I look closely __Couchbase__.

#### “Veteran” Award
A user receives this for playing more than 1000 games in their lifetime.

The __RethinkDB__ query:
```js
```

#### “Bruiser” Award
A user receives this for doing more than 500 points of damage in one game

The __RethinkDB__ query:

```js
r.db("osmo").table("eventstore")
.filter({
  "AggregateID":  "b3053a78-1b9f-4000-b6ed-70d9ca9b64a4",
  "AggregateType":  "Player",
  "EventType":  "TotalAmountOfDamageDoneUpdated"
})
.map(function(val) {
  return val("Data")
})
.without(["AggregateID", "Game"])
.sum("TotalAmountOfDamageDone")
.gt(500);
```
