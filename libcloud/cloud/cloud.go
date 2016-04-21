package cloud

import (
	"github.com/docker/libcompose/docker"
	"github.com/docker/libcompose/project"
	"golang.org/x/net/context"
)

type Clouder interface {
	Start() error
	Stop() error
}

type clouder struct {
	Deployer
}

func (c *clouder) Start() (err error) {

	//check
	err = c.Deploy()
	if err != nil {
		return
	}

	//set docker env

	//run
	project, err := docker.NewProject(&docker.Context{
		Context: project.Context{
			ComposeFiles: []string{"docker-compose.yml"},
			ProjectName:  "my-compose",
		},
	})

	if err != nil {
		return
	}

	err = project.Up()
	return
}

func (c *clouder) Stop() (err error) {
	//stop
	return
}

func NewClouder() (c Clouder) {
	ctx := context.Background()
	d := NewDeployer(ctx)
	return &clouder{Deployer: d}
}
