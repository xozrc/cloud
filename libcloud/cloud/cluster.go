package cloud

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/machine/libmachine/host"
	cloudutils "github.com/xozrc/cloud/libcloud/utils"
	"github.com/xozrc/deploy"

	deployConfig "github.com/xozrc/deploy/config"
)

type Cluster interface {
	Load() error
	ClusterName() string
	SwarmMaster() *host.Host
	Nodes() []*host.Host
	Destroy() error
}

type cluster struct {
	name        string
	nodes       []*host.Host
	swarmMaster *host.Host
	nodesFile   []string
}

func (c *cluster) Load() error {
	log.Infoln("trying to load cluster nodes")

	for _, nodeFile := range c.nodesFile {
		//nodeName
		nc, err := deployConfig.NodeConfigFromFile(nodeFile)
		if err != nil {
			return err
		}

		//rewrite cluster config
		nc.StorePath = cloudutils.GetBaseDir()
		nc.Name = c.ClusterName() + "-" + nc.Name

		//config swarm
		if nc.Swarm == nil {
			nc.Swarm = &deployConfig.SwarmConfig{}
			nc.Swarm.Swarm = true
		}

		h, err := deploy.Deploy(nc)

		//fix: if need to ignore error
		if err != nil {
			return err
		}

		err = c.appendHost(h)
		//fix: if need to ignore error
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *cluster) Destroy() error {
	log.Infoln("trying to destroy cluster nodes")

	for _, nodeFile := range c.nodesFile {
		//nodeName
		nc, err := deployConfig.NodeConfigFromFile(nodeFile)
		if err != nil {
			return err
		}
		nc.StorePath = cloudutils.GetBaseDir()
		nc.Name = c.ClusterName() + "-" + nc.Name
		deploy.Undeploy(nc)
	}
	return nil
}

func (c *cluster) ClusterName() string {
	return c.name
}

func (c *cluster) Nodes() []*host.Host {
	return c.nodes
}

func (c *cluster) SwarmMaster() *host.Host {
	return c.swarmMaster
}

func (c *cluster) appendHost(n *host.Host) (err error) {

	if n.HostOptions.SwarmOptions == nil || !n.HostOptions.SwarmOptions.IsSwarm {
		return fmt.Errorf("cluster should be swarm")
	}

	if n.HostOptions.SwarmOptions.Master {
		// err = fmt.Errorf("cluster should only one master")
		// return err
		c.swarmMaster = n
	}

	c.nodes = append(c.nodes, n)
	return
}

func NewCluster(clusterName string, nodes []string) (c Cluster, err error) {
	tc := &cluster{}
	tc.name = clusterName
	tc.nodesFile = nodes
	c = tc
	return
}
