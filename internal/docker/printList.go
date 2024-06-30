package docker

import (
	"bytes"
	"fmt"
	"io"
	"path"
	"strings"
	"text/template"
	"time"

	"github.com/docker/cli/cli/command/formatter"
	"github.com/docker/docker/api/types"
	"github.com/docker/go-units"
)

var (
	sep                       = ", "
	containerInfoTemplateFile = "internal/docker/template.tmpl"
)

type ContainerInfo struct {
	Name        string
	ID          string
	Image       string
	Command     string
	CreatedTime string
	Status      string
	Network     string
	IPAddresses string
	Ports       string
	Mounts      string
}

func PrintContainersList(containers []types.Container, w io.Writer) error {
	if len(containers) == 0 {
		return nil
	}

	for i, container := range containers {
		x := convertContainerToContainerInfo(container)
		result, err := x.render()
		if err != nil {
			return fmt.Errorf("can't render container info: %v", err)
		}

		fmt.Fprintf(w, result)
		if i != len(containers)-1 {
			fmt.Fprintf(w, "\n")
		}
	}

	return nil
}

func convertContainerToContainerInfo(c types.Container) ContainerInfo {
	return ContainerInfo{
		Name:        c.Names[0][1:], // originally it starts from "/"
		ID:          c.ID[:12],
		Image:       c.Image,
		Command:     c.Command,
		CreatedTime: units.HumanDuration(time.Now().UTC().Sub(time.Unix(c.Created, 0))) + " ago",
		Status:      c.Status,
		Network:     c.HostConfig.NetworkMode,
		IPAddresses: getIPAddress(c),
		Ports:       formatter.DisplayablePorts(c.Ports),
		Mounts:      getContainerMounts(c),
	}
}

func getIPAddress(c types.Container) string {
	var ip []string
	if c.NetworkSettings != nil {
		for _, val := range c.NetworkSettings.Networks {
			if val.IPAddress != "" {
				address := fmt.Sprintf("%v/%v", val.IPAddress, val.IPPrefixLen)
				ip = append(ip, address)
			}
		}
	}
	return strings.Join(ip, sep)
}

func getContainerMounts(c types.Container) string {
	var mounts []string
	for _, m := range c.Mounts {
		mount := fmt.Sprintf("%v:%v", m.Source, m.Destination)
		mounts = append(mounts, mount)
	}
	return strings.Join(mounts, sep)
}

func (c *ContainerInfo) render() (string, error) {
	tmpl := template.Must(template.New(path.Base(containerInfoTemplateFile)).ParseFiles(containerInfoTemplateFile))

	buff := &bytes.Buffer{}
	err := tmpl.Execute(buff, c)
	if err != nil {
		fmt.Printf("error executing template: %v", err)
	}

	return buff.String(), nil
}
