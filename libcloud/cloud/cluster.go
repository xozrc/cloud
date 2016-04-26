package cloud

import (
	"fmt"

	"github.com/docker/machine/libmachine/auth"
	"github.com/docker/machine/libmachine/host"
	cloudutils "github.com/xozrc/cloud/libcloud/utils"
	deployConfig "github.com/xozrc/deploy/config"
	"github.com/xozrc/deploy/libdeploy"
)

var ()

type Cluster interface {
	Load() error
	ClusterName() string
	SwarmMaster() *host.Host
	Hosts() []*host.Host
	AppendHost(h *host.Host) error
	Destroy() error
}

type cluster struct {
	name         string
	hosts        []*host.Host
	swarmMaster  *host.Host
	nodes        []string
	clusterStore string
}

func (c *cluster) Load() error {

	for _, node := range c.nodes {
		//nodeName
		mc, err := deployConfig.MachineConfigFromFile(node)
		if err != nil {
			return err
		}

		mc.Name = c.ClusterName() + "-" + mc.Name
		if mc.Auth == nil {
			mc.Auth = &auth.Options{}
		}
		mc.Auth.StorePath = cloudutils.GetBaseDir()
		h, err := libdeploy.Deploy(mc)
		if err != nil {
			return err
		}
		err = c.AppendHost(h)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *cluster) Destroy() error {

	for _, node := range c.nodes {
		//nodeName
		mc, err := deployConfig.MachineConfigFromFile(node)
		if err != nil {
			return err
		}
		if mc.Auth == nil {
			mc.Auth = &auth.Options{}
		}
		mc.Auth.StorePath = cloudutils.GetBaseDir()
		mc.Name = c.ClusterName() + "-" + mc.Name
		libdeploy.Undeploy(mc)
	}
	return nil
}

func (c *cluster) ClusterName() string {
	return c.name
}

func (c *cluster) Hosts() []*host.Host {
	return c.hosts
}

func (c *cluster) SwarmMaster() *host.Host {
	return c.swarmMaster
}

func (c *cluster) AppendHost(h *host.Host) (err error) {
	if h.HostOptions == nil || h.HostOptions.SwarmOptions == nil {
		return
	}

	if !h.HostOptions.SwarmOptions.IsSwarm {
		//err = fmt.Errorf("host [%s] should be swarm", h.Name)
		return
	}

	if h.HostOptions.SwarmOptions.Master {
		if c.SwarmMaster() != nil {
			err = fmt.Errorf("cluster should only one master")
			return err
		}
		c.swarmMaster = h
	}

	c.hosts = append(c.hosts, h)
	return
}

func NewCluster(clusterName string, nodes []string, clusterStore string) (c Cluster, err error) {
	tc := &cluster{}
	tc.name = clusterName
	tc.nodes = nodes
	tc.clusterStore = clusterStore
	c = tc
	return
}
