package cli

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

const ()

var (
	upCmd = cli.Command{
		Name:        "up",
		Usage:       "starts and provisions the cloud environment",
		Description: "",
		Action:      up,
	}
)

func init() {
	appendCmd(upCmd)
}

func up(c *cli.Context) {
	tc, err := clouderFromFile(cloudFileVal)
	if err != nil {
		log.Fatalln(err)
	}
	if err := tc.Load(); err != nil {
		log.Fatalln(err)
	}

	log.Fatalln(tc.Start())
}
