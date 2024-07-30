package main

import (
	"fmt"
	"log"
	"os"

	"github.com/docker/cli/cli-plugins/manager"
	"github.com/docker/cli/cli-plugins/plugin"
	"github.com/docker/cli/cli/command"
	"github.com/sergkondr/docker-ps/internal/docker"
	"github.com/spf13/cobra"
)

var (
	version = "dev"
)

func main() {
	plugin.Run(
		func(dockerCli command.Cli) *cobra.Command {
			var (
				showAll bool
			)

			cmd := &cobra.Command{
				Use:   "cps",
				Short: "The cps (custom ps) is a Docker plugin that displays a list of running containers in a more readable and informative format than the standard docker ps command.",
				Long: `The cps (custom ps) is a Docker plugin that displays a list of both running and stopped containers 
in a more readable and informative format than the standard docker ps command. 
It helps users quickly access important details about container status and other relevant information.`,
				Run: func(cmd *cobra.Command, args []string) {
					opts := docker.GetOpts{
						Client:  dockerCli,
						ShowAll: showAll,
					}

					containers, err := docker.GetList(opts)
					if err != nil {
						log.Fatal("can't list containers:", err)
					}

					if len(containers) == 0 {
						fmt.Println("no running containers")
						return
					}

					err = docker.PrintContainersList(containers, os.Stdout)
					if err != nil {
						log.Fatal("can't print containers:", err)
					}
				},
			}

			flags := cmd.Flags()
			flags.BoolVarP(&showAll, "all", "a", false, "Display all containers")

			cmd.AddCommand()
			return cmd
		},

		manager.Metadata{
			SchemaVersion:    "0.1.0",
			Vendor:           "Sergei Kondrashov",
			Version:          version,
			ShortDescription: "Docker convenient ps",
			URL:              "https://github.com/sergkondr/docker-ps",
		})
}
