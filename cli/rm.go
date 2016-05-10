package cli

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

const ()

var (
	rmCmd = cli.Command{
		Name:        "rm",
		Usage:       "rm the cloud environment",
		Description: "",
		Action:      rm,
	}
)

func init() {
	appendCmd(rmCmd)
}

func rm(c *cli.Context) {
	tc, err := clouderFromFile(cloudFileVal)
	if err != nil {
		log.Fatalln(err)
	}

	tc.Destroy()
	log.Info("remove cloud successfully")
}
