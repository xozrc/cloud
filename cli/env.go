package cli

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

var (
	envCmd = cli.Command{
		Name:        "env",
		Usage:       "cloud env",
		Description: "",
		Action:      env,
	}
)

func init() {
	appendCmd(envCmd)
}

func env(c *cli.Context) {
	tc, err := clouderFromFile(cloudFileVal)
	if err != nil {
		log.Fatalln(err)
	}
	if err := tc.Load(); err != nil {
		log.Fatalln(err)
	}

	if err := tc.Start(); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(tc.Env())
}
