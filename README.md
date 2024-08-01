# docker-ps

[![License: MIT](https://img.shields.io/badge/License-MIT%202.0-blue.svg)](https://github.com/sergkondr/docker-ps/blob/main/LICENSE)
[![GitHub release](https://img.shields.io/github/release/sergkondr/docker-ps.svg)](https://github.com/sergkondr/docker-ps/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/sergkondr/docker-ps)](https://goreportcard.com/report/github.com/sergkondr/docker-ps)
[![Go](https://github.com/sergkondr/docker-ps/actions/workflows/go.yml/badge.svg)](https://github.com/sergkondr/docker-ps/actions/workflows/go.yml)

### Purpose
As for me, `docker ps` is uninformative and inconvenient.
This app is a Docker plugin that displays a list of running containers in a more 
readable and informative format than the standard `docker ps` command.

### Download
Compiled binary you can get [in releases](https://github.com/sergkondr/docker-ps/releases)

### Install
```
mv ./docker-cps ~/.docker/cli-plugins/docker-cps
```

### Usage
```
➜ docker cps -a
jovial_mendel
    Container ID:    70a838c8dcb7
    Image:           alpine:3.20
    Command:         ash
    Created:         1 second ago
    IP-address:      172.17.0.2/16
    Ports:           0.0.0.0:8080-8081->8080-8081/tcp, 0.0.0.0:9090->9090/tcp
    Status:          Up 1 second

tender_johnson
    Container ID:    e3913b691066
    Image:           alpine:3.20
    Command:         echo 123
    Created:         6 minutes ago
    Status:          Exited (0) 6 minutes ago

```
