#!/bin/bash

docker network create traefik

MY_IF=$(netstat -rn | awk '/^0.0.0.0/ {thif=substr($0,74,10); print thif;} /^default.*UG/ {thif=substr($0,65,10); print thif;}' | head -n 1 | xargs)
MY_IP=$(ifconfig $MY_IF | grep -Eo 'inet (ad+r:)?([0-9]*\.){3}[0-9]*' | grep -Eo '([0-9]*\.){3}[0-9]*' | grep -v '127.0.0.1')
echo "Localhost IP: ${MY_IP}"

export IP_LOCALHOST=$MY_IP

docker-compose down
docker-compose -f backends/docker-compose.yml down
docker-compose -f frontends/docker-compose.yml down
