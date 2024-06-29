package docker

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
	"time"

	"github.com/docker/cli/cli/command/formatter"
	"github.com/docker/docker/api/types"
	"github.com/docker/go-units"
)

var ()

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

func (c *ContainerInfo) Render() (string, error) {
	containerInfoTemplate := `{{ .Name }}
    Container ID:    {{ .ID }}
    Image:           {{ .Image }}
    Command:         {{ .Command }}
    Created:         {{ .CreatedTime }}
{{ if .Mounts }}    Mounts:          {{ .Mounts }}
{{ end }}    Network:         {{ .Network }}
{{ if .IPAddresses }}    IP-address:      {{ .IPAddresses }}
{{ end }}{{ if .Ports }}    Ports:           {{ .Ports }}
{{ end }}    Status:          {{ .Status }}
`

	tmpl, err := template.New("container").Parse(containerInfoTemplate)
	if err != nil {
		return "", fmt.Errorf("error parsing template: %v", err)
	}

	buff := &bytes.Buffer{}
	err = tmpl.Execute(buff, c)
	if err != nil {
		fmt.Printf("error executing template: %v", err)
	}

	return buff.String(), nil
}

func PrintContainersList(containers []types.Container) error {
	if len(containers) == 0 {
		return nil
	}

	for _, container := range containers {
		x := convertContainerToContainerInfo(container)
		result, err := x.Render()
		if err != nil {
			return fmt.Errorf("can't render container info: %v", err)
		}

		fmt.Println(result)
	}

	return nil
}

func convertContainerToContainerInfo(c types.Container) ContainerInfo {
	return ContainerInfo{
		Name:        c.Names[0][1:],
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
	return strings.Join(ip, ", ")
}

func getContainerMounts(c types.Container) string {
	var mounts []string
	for _, m := range c.Mounts {
		mount := fmt.Sprintf("%v:%v", m.Source, m.Destination)
		mounts = append(mounts, mount)
	}
	return strings.Join(mounts, ", ")
}
