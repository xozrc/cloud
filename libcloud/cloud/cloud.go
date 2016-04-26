package cloud

import "github.com/xozrc/cloud/libcloud/config"

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

	//set docker env
	//swarmMaster := c.Cluster().SwarmMaster()

	//run
	// project, err := docker.NewProject(&docker.Context{
	// 	Context: project.Context{
	// 		ComposeFiles: []string{"docker-compose.yml"},
	// 		ProjectName:  "my-compose",
	// 	},
	// })

	// if err != nil {
	// 	return
	// }

	// err = project.Up()
	return
}

func (c *clouder) Stop() (err error) {

	//stop
	return
}

func NewClouder(cfg *config.CloudConfig) (c Clouder, err error) {
	tc := &clouder{}
	tc.Cluster, err = NewCluster(cfg.Name, cfg.ComposeFiles, cfg.ClusterStore)
	if err != nil {
		return
	}
	tc.composeFiles = cfg.ComposeFiles
	c = tc
	return
}
