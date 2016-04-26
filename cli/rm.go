package cli

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/xozrc/cloud/libcloud/cloud"
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
	tc := cloud.NewClouder()
	tc.Destroy()
	log.Info("remove cloud successfully")
}
