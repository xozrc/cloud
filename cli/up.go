package cli

import (
	"log"

	"github.com/codegangsta/cli"
	"github.com/xozrc/cloud/libcloud/cloud"
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
	tc := cloud.NewClouder()
	log.Fatalln(tc.Start())
}
