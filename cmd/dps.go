package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/sergkondr/docker-ps/internal/docker"
)

var (
	flagAll = false
)

func main() {
	flag.BoolVar(&flagAll, "all", false, "Display all containers")
	flag.Parse()

	socketAddr, err := getDockerSocketAddress()
	if err != nil {
		log.Fatal("can't get docker socket location: ", err)
	}

	opts := docker.GetOpts{
		SocketAddr: socketAddr,
		ShowAll:    flagAll,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	containers, err := docker.GetList(ctx, opts)
	if err != nil {
		log.Fatal("can't list containers:", err)
	}
	err = docker.PrintContainersList(containers)
	if err != nil {
		log.Fatal("can't print containers:", err)
	}
}

func getDockerSocketAddress() (string, error) {
	if dockerHost := os.Getenv("DOCKER_HOST"); dockerHost != "" {
		return dockerHost, nil
	}

	socketProto := "unix://"
	socketPath := ""
	switch runtime.GOOS {
	case "darwin":
		if homeDir := os.Getenv("HOME"); homeDir != "" {
			socketPath += homeDir + "/.docker/run/docker.sock"
		}
	case "linux":
		socketPath += "/var/run/docker.sock"
	// TODO: windows support
	default:
		return "", fmt.Errorf("unsupported platform")
	}

	_, err := os.Stat(socketPath)
	if errors.Is(err, os.ErrNotExist) {
		return "", fmt.Errorf("can't find docker socket at %s, check that the docker daemon is running", socketPath)
	}
	if errors.Is(err, os.ErrPermission) {
		return "", fmt.Errorf("permissions denied: %s", socketPath)
	}
	if err != nil {
		return "", fmt.Errorf("can't check docker socket file at %s: %v", socketPath, err)
	}

	return socketProto + socketPath, nil
}
