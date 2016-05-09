package cloud_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xozrc/cloud/libcloud/cloud"
	"github.com/xozrc/cloud/libcloud/config"
)

func TestNewCloud(t *testing.T) {
	tcfg, err := config.NewCloudConfig("./CloudConfig.yaml")

	if !assert.NoError(t, err, "") {
		return
	}
	tc, err := cloud.NewClouder(tcfg)
	if !assert.NoError(t, err, "") {
		return
	}
	err = tc.Load()
	if !assert.NoError(t, err, "") {
		return
	}

	err = tc.Start()

	if !assert.NoError(t, err, "") {
		return
	}

}

func TestNewCluster(t *testing.T) {
	clusterName := "test"
	nodes := []string{"./nodes/master.yaml", "./nodes/node1.yaml", "./nodes/node2.yaml"}

	cls, err := cloud.NewCluster(clusterName, nodes)
	if !assert.NoError(t, err, "") {
		return
	}

	err = cls.Load()

	assert.NoError(t, err, "")
}
