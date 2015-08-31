package server

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/docker/libcompose/cli/app"
)

func ServerCommand(factory app.ProjectFactory) cli.Command {
	return cli.Command{
		Name:  "server",
		Usage: "Run a Remote Compose server",
		Action: func(c *cli.Context) {
			log.Infof("Starting Server...")
			StartServer()
		},
	}
}
