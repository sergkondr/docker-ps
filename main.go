package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/go-units"
	"github.com/mgutz/ansi"
	"golang.org/x/net/context"
)

var (
	all   bool
	match string

	seporator = ", "
)

func init() {
	flag.BoolVar(&all, "all", false, "show containers with all statuses")
	flag.BoolVar(&all, "a", false, "show containers with all statuses (short)")

	flag.StringVar(&match, "match", "", "show only containers match with string")
	flag.StringVar(&match, "m", "", "show only containers match with string (short)")

	flag.Parse()
}

func main() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containerListOptions := types.ContainerListOptions{}

	containers, err := cli.ContainerList(context.Background(), containerListOptions)
	if err != nil {
		panic(err)
	}

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 20, 24, 4, '\t', 0)
	wFormat := "    %v\t%v\n"
	defer w.Flush()

	containersInfo := []containerInfo{}
	for _, container := range containers {
		ports := getContainerPorts(container)
		mounts := getContainerMounts(container)
		ip := getIPAddress(container)

		c := containerInfo{
			Name:        container.Names[0][1:],
			ID:          container.ID[:12],
			Image:       container.Image,
			Command:     container.Command,
			CreatedTime: units.HumanDuration(time.Now().UTC().Sub(time.Unix(container.Created, 0))) + " ago",
			Status:      container.Status,
			Network:     container.HostConfig.NetworkMode,
			IPAddresses: container.NetworkSettings.Networks,
			Ports:       container.Ports,
			Mounts:      container.Mounts,
		}

		if match == "" || (match != "" && c.containsString(match)) {
			containersInfo = append(containersInfo, c)
			fmt.Printf("%v\n", ansi.Color(c.Name, "cyan+b"))
			fmt.Fprintf(w, wFormat, ansi.Color("Container ID:", "+b"), c.ID)
			fmt.Fprintf(w, wFormat, ansi.Color("Image:", "+b"), c.Image)
			fmt.Fprintf(w, wFormat, ansi.Color("Command:", "+b"), c.Command)
			fmt.Fprintf(w, wFormat, ansi.Color("Created:", "+b"), c.CreatedTime)
			fmt.Fprintf(w, wFormat, ansi.Color("Status:", "+b"), c.Status)
			fmt.Fprintf(w, wFormat, ansi.Color("Network:", "+b"), c.Network)

			if len(ip) != 0 {
				fmt.Fprintf(w, wFormat, ansi.Color("IP-address:  ", "+b"), getIPAddress(container))
			}
			if len(ports) != 0 {
				fmt.Fprintf(w, wFormat, ansi.Color("Ports:", "+b"), ports)
			}
			if len(mounts) != 0 {
				fmt.Fprintf(w, wFormat, ansi.Color("Container mounts:", "+b"), mounts)
			}
			fmt.Fprintf(w, "\n")
		}
	}

	description := ""
	switch containerAmount := strconv.Itoa(len(containers)); containerAmount {
	case "0":
		description = "There is no containers running"
	case "1":
		description = ansi.Color("1", "+b") + " container is running"
	default:
		description = ansi.Color(containerAmount, "+b") + " containers are running"
	}
	fmt.Fprintf(w, "  %s\n", description)
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
	return strings.Join(ip, seporator)
}

func getContainerMounts(container types.Container) string {
	var mounts []string
	for _, m := range container.Mounts {
		mount := fmt.Sprintf("%v:%v", m.Source, m.Destination)
		mounts = append(mounts, mount)
	}
	return strings.Join(mounts, seporator)
}
