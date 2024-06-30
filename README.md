# docker-ps [![Build Status](https://travis-ci.com/sergkondr/docker-ps.svg?branch=master)](https://travis-ci.com/sergkondr/docker-ps) [![Go Report Card](https://goreportcard.com/badge/github.com/sergkondr/docker-ps)](https://goreportcard.com/report/github.com/sergkondr/docker-ps)

### Purpose
As for me, `docker ps` is uninformative and inconvenient. Here I'm trying to create something for convenient displaying running containers.

### Download
Compiled binary you can get [in releases](https://github.com/sergkondr/docker-ps/releases)

### Compile
```
make build
```

### Usage
```
$ dps
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

```
