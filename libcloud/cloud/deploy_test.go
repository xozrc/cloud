package cloud_test

import (
	"testing"

	"github.com/docker/machine/commands/mcndirs"
	"github.com/docker/machine/libmachine"
	"github.com/stretchr/testify/assert"
	"github.com/xozrc/cloud/libcloud/cloud"
)

func TestDeploy(t *testing.T) {
	api := libmachine.NewClient(mcndirs.GetBaseDir(), mcndirs.GetMachineCertDir())
	defer api.Close()
	deployer, err := cloud.NewDeployer(api, "test-cloud", 2, "")

	if !assert.NoError(t, err, "") {
		return
	}

	err = deployer.Deploy()
	if !assert.NoError(t, err, "") {
		return
	}
	return
}
