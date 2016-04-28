package cloud

import (
	"fmt"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/libcompose/docker"
	"github.com/docker/libcompose/project"
	"github.com/docker/machine/libmachine"
	"github.com/xozrc/cloud/libcloud/config"
	cloudutils "github.com/xozrc/cloud/libcloud/utils"
)

type Clouder interface {
	Load() error
	Start() error
	Stop() error
	Destroy() error
}

type clouder struct {
	Cluster
	composeFiles []string
}

func (c *clouder) Start() (err error) {
	log.Infoln("trying to start cloud")
	//set docker env
	tsm := c.SwarmMaster()

	api := libmachine.NewClient(cloudutils.GetBaseDir(), tsm.HostOptions.AuthOptions.CertDir)

	h, err := api.Load(tsm.Name)
	if err != nil {
		return
	}
	defer api.Close()

	if err = os.Setenv("DOCKER_TLS_VERIFY", "1"); err != nil {
		return
	}

	log.Infoln("export DOCKER_TLS_VERIFY='1'")

	tCertPath := tsm.AuthOptions().CertDir
	if err = os.Setenv("DOCKER_CERT_PATH", tCertPath); err != nil {
		return
	}

	log.Infoln("export DOCKER_CERT_PATH=" + tCertPath)

	tIP, err := h.Driver.GetIP()
	if err != nil {
		return
	}

	//split but remove //
	tparts := strings.Split(tsm.HostOptions.SwarmOptions.Host, ":")
	if len(tparts) != 3 {
		return fmt.Errorf("swarm host error format")
	}

	th := tparts[0] + "://" + tIP + ":" + tparts[2]

	if err = os.Setenv("DOCKER_HOST", th); err != nil {
		return
	}

	log.Infoln("export DOCKER_HOST=" + th)

	tname := tsm.Name
	if err = os.Setenv("DOCKER_MACHINE_NAME", tname); err != nil {
		return
	}

	log.Infoln("export DOCKER_MACHINE_NAME=" + tname)

	//run
	project, err := docker.NewProject(&docker.Context{
		Context: project.Context{
			ComposeFiles: c.composeFiles,
			ProjectName:  c.ClusterName(),
		},
	})

	if err != nil {
		return
	}

	err = project.Up()
	return
}

func (c *clouder) Stop() (err error) {

	cloudName := c.ClusterName()
	project, err := docker.NewProject(&docker.Context{
		Context: project.Context{
			ComposeFiles: c.composeFiles,
			ProjectName:  cloudName,
		},
	})
	if err != nil {
		return
	}

	project.Down()
	//stop
	return
}

func NewClouder(cfg *config.CloudConfig) (c Clouder, err error) {
	tc := &clouder{}
	tc.Cluster, err = NewCluster(cfg.Name, cfg.Nodes)
	if err != nil {
		return
	}
	tc.composeFiles = cfg.ComposeFiles
	c = tc
	return
}
