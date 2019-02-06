package main

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/network"
	"strings"
)

type containerInfo struct {
	Name        string
	ID          string
	Image       string
	Command     string
	CreatedTime string
	Status      string
	Network     string
	IPAddresses map[string]*network.EndpointSettings
	Ports       []types.Port
	Mounts      []types.MountPoint
}

func (c *containerInfo) String() string {
	return ""
}

func (c *containerInfo) containsString(s string) bool {
	if strings.Contains(c.Name, s) ||
		strings.Contains(c.ID, s) ||
		strings.Contains(c.Image, s) ||
		strings.Contains(c.Command, s) ||
		strings.Contains(c.CreatedTime, s) ||
		strings.Contains(c.Status, s) ||
		strings.Contains(c.Network, s) {
		return true
	}
	for _, p := range c.Ports {
		port := fmt.Sprintf("%v %v %v %v", p.IP, p.PrivatePort, p.PublicPort, p.Type)
		if strings.Contains(port, s) {
			return true
		}
	}
	for _, m := range c.Mounts {
		mount := fmt.Sprintf("%v %v %v", m.Name, m.Destination, m.Source)
		if strings.Contains(mount, s) {
			return true
		}
	}
	for _, i := range c.IPAddresses {
		if strings.Contains(i.IPAddress, s) {
			return true
		}
	}

	return false
}
