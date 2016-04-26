package cloud_test

import (
	"testing"
)
import (
	"github.com/stretchr/testify/assert"
	"github.com/xozrc/cloud/libcloud/cloud"
)

func TestNewCloud(t *testing.T) {

}

func TestNewCluster(t *testing.T) {
	clusterName := "test"
	nodes := []string{"./nodes/master.yaml", "./nodes/node1.yaml", "./nodes/node2.yaml"}
	clusterStore := "consul://192.168.99.101:8500/"
	cls, err := cloud.NewCluster(clusterName, nodes, clusterStore)
	if !assert.NoError(t, err, "") {
		return
	}

	err = cls.Load()
	defer cls.Destroy()
	assert.NoError(t, err, "")

}
