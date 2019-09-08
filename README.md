# OSMO Coding Challenge

## Infrastructure

Read the README in __infra__ directory for infrastructure details.

```sh
cd infra
./start.sh
cd ..
```

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
