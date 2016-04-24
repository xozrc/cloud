package cloud

import (
	"fmt"

	"github.com/docker/machine/libmachine"
	"github.com/docker/machine/libmachine/host"
)

var ()

type Cluster interface {
	SwarmMaster() *host.Host
	Hosts() []*host.Host
	AppendHost(h *host.Host) error
}

type cluster struct {
	hosts       []*host.Host
	swarmMaster *host.Host
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

func NewCluster(api libmachine.API) (c Cluster, err error) {
	hostNames, err := api.List()
	if err != nil {
		return
	}

	c = &cluster{}

	for _, hostName := range hostNames {
		h, err := api.Load(hostName)
		if err != nil {
			return nil, err
		}

		if err = c.AppendHost(h); err != nil {
			return nil, err
		}
	}

	return
}
