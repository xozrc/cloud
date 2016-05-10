package cli

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

const ()

var (
	restartCmd = cli.Command{
		Name:        "restart",
		Usage:       "restart the cloud environment",
		Description: "",
		Action:      restart,
	}
)

func init() {
	appendCmd(restartCmd)
}

func restart(c *cli.Context) {
	tc, err := clouderFromFile(cloudFileVal)
	if err != nil {
		log.Fatalln(err)
	}

	if err := tc.Load(); err != nil {
		log.Fatalln(err)
	}

	log.Fatalln(tc.Start())
}
