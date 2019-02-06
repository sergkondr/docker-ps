docker-ps
===

## Purpose
As for me, `docker ps` is uninformative and inconvenient. Here I'm trying to create something for convient displaying running containers.

## Compile
```
go build -o docker-ps .
```

## Usage
```
$ docker-ps
clever_noyce
    Container ID:       f7ea38fdcc90
    Image:              ubuntu:18.04
    Command:            bash
    Created:            14 seconds ago
    Status:             Up 13 seconds
    Network:            default
    IP-address:         172.17.0.4/16
    Container mounts:   /home/skondrashov:/home/user

friendly_dubinsky
    Container ID:       aeb5ad7b4b6e
    Image:              ubuntu:18.04
    Command:            bash -c 'echo 123 ; echo 123 ; echo 123 ; echo 123 ; echo 123 ; echo 123 ; bash'
    Created:            6 seconds ago
    Status:             Up 5 seconds
    Network:            default
    IP-address:         172.17.0.3/16
    Ports:              0.0.0.0:2223 -> 22/tcp, 0.0.0.0:8081 -> 80/tcp

awesome_thompson
    Container ID:       c896cbd44823
    Image:              ubuntu:18.04
    Command:            /entrypoint.sh bash
    Created:            42 minutes ago
    Status:             Up 42 minutes
    Network:            default
    IP-address:         172.17.0.2/16
    Ports:              0.0.0.0:2222 -> 22/tcp, 0.0.0.0:8080 -> 80/tcp

 3 containers are running
```

It also supports filtering of displaying containers by some string(it can be part of any parameter, e.g.: IP, volume, ):
```
$ docker-ps -m 172.17.0.2
awesome_thompson
    Container ID:       c896cbd44823
    Image:              ubuntu:18.04
    Command:            /entrypoint.sh bash
    Created:            About an hour ago
    Status:             Up About an hour
    Network:            default
    IP-address:         172.17.0.2/16
    Ports:              0.0.0.0:2222 -> 22/tcp, 0.0.0.0:8080 -> 80/tcp

  3 containers are running
$ docker-ps -m home
clever_noyce
    Container ID:       f7ea38fdcc90
    Image:              ubuntu:18.04
    Command:            bash
    Created:            About a minute ago
    Status:             Up About a minute
    Network:            default
    IP-address:         172.17.0.4/16
    Container mounts:   /home/skondrashov:/home/user

  3 containers are running
```