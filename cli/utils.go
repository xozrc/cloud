package cli

import (
	"path/filepath"

	"github.com/xozrc/cloud/libcloud/cloud"
	"github.com/xozrc/cloud/libcloud/config"
)

func clouderFromFile(f string) (c cloud.Clouder, err error) {

	f, err = filepath.Abs(f)
	if err != nil {
		return
	}

	cfg, err := config.NewCloudConfig(f)
	if err != nil {
		return
	}
	tc, err := cloud.NewClouder(cfg)
	if err != nil {
		return
	}
	c = tc
	return
}
