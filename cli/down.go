package cli

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

var (
	downCmd = cli.Command{
		Name:        "down",
		Usage:       "stop containers",
		Description: "",
		Action:      down,
	}
)

func init() {
	appendCmd(downCmd)
}

func down(c *cli.Context) {
	tc, err := clouderFromFile(cloudFileVal)
	if err != nil {
		log.Fatalln(err)
	}
	if err := tc.Load(); err != nil {
		log.Fatalln(err)
	}

	log.Fatalln(tc.Stop())
}
