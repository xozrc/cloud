package cloud

import (
	"encoding/json"
	"fmt"
	"log"
)
import (
	"github.com/docker/machine/commands/mcndirs"
	"github.com/docker/machine/libmachine"
	"github.com/docker/machine/libmachine/drivers"
	"github.com/docker/machine/libmachine/host"
)

type Deployer interface {
	Deploy() error
	Undeploy() error
	Cluster() Cluster
}

func NewDeployer(api libmachine.API, cloudName string, nodesNum int, clusterStore string) (d Deployer, err error) {
	cl, err := NewCluster(api)
	if err != nil {
		return
	}
	d = &deployer{
		cl:           cl,
		api:          api,
		cloudName:    cloudName,
		nodesNum:     nodesNum,
		clusterStore: clusterStore,
	}
	return
}

type deployer struct {
	cl           Cluster
	api          libmachine.API
	cloudName    string
	nodesNum     int
	clusterStore string
}

func (d *deployer) Cluster() Cluster {
	return d.cl
}

func (d *deployer) Deploy() (err error) {

	//check machine nums
	if d.nodesNum <= 1 {
		return fmt.Errorf("cluster machines should more than 1")
	}

	//deploy config
	driverName := "virtualbox"
	hostName := "test"
	var driverOpts drivers.DriverOptions
	var hostOpts *host.Options
	curNum := len(d.cl.Hosts())

	//create new machines
	if curNum < d.nodesNum {
		log.Printf("there is %d nodes, less than config nodes num %d,trying to create new\n", curNum, d.nodesNum)
		//deploy master
		if d.cl.SwarmMaster() == nil {
			log.Println("there is no swarm master node ,trying to create")
			h, err := d.deployHost(true, driverName, driverOpts, hostName, hostOpts)
			if err != nil {
				return err
			}

			if err = d.cl.AppendHost(h); err != nil {
				return err
			}
		}

		//deploy remains
		curNum = len(d.cl.Hosts())
		if curNum < d.nodesNum {
			h, err := d.deployHost(false, driverName, driverOpts, hostName, hostOpts)
			if err != nil {
				return err
			}

			if err = d.cl.AppendHost(h); err != nil {
				return err
			}
		}
	}

	return

	// var (
	// 	certDir         string
	// 	caCertPath      string
	// 	caPriverKeyPath string
	// 	clientCertPath  string
	// 	clientKeyPath   string
	// 	serverCertPath  string
	// 	serverKeyPath   string
	// 	storePath       string
	// 	serverCertSANs  []string
	// )

	// var (
	// 	engineOpt              []string
	// 	engineEnv              []string
	// 	engineInsecureRegistry []string
	// 	engineLabel            []string
	// 	engineRegistryMirror   []string
	// 	engineStorageDirver    string
	// 	engineTLSVerify        bool
	// 	engineInstallURL       string
	// )

	// var (
	// 	isSwarm           bool
	// 	swarmImage        string
	// 	swarmMaster       bool
	// 	swarmDiscovery    string
	// 	swarmAddr         string
	// 	swarmHost         string
	// 	swarmStrategy     string
	// 	swarmOpts         []string
	// 	swarmExperimental bool
	// )

	// h.HostOptions = &host.Options{
	// 	AuthOptions: &auth.Options{
	// 		CertDir:          certDir,
	// 		CaCertPath:       caCertPath,
	// 		CaPrivateKeyPath: caPriverKeyPath,
	// 		ClientCertPath:   clientCertPath,
	// 		ClientKeyPath:    clientKeyPath,
	// 		ServerCertPath:   serverCertPath,
	// 		ServerKeyPath:    serverKeyPath,
	// 		StorePath:        storePath,
	// 		ServerCertSANs:   serverCertSANs,
	// 	},
	// 	EngineOptions: &engine.Options{
	// 		ArbitraryFlags:   engineOpt,
	// 		Env:              engineEnv,
	// 		InsecureRegistry: engineInsecureRegistry,
	// 		Labels:           engineLabel,
	// 		RegistryMirror:   engineRegistryMirror,
	// 		StorageDriver:    engineStorageDirver,
	// 		TLSVerify:        engineTLSVerify,
	// 		InstallURL:       engineInstallURL,
	// 	},
	// 	SwarmOptions: &swarm.Options{
	// 		IsSwarm:        isSwarm,
	// 		Image:          swarmImage,
	// 		Master:         swarmMaster,
	// 		Discovery:      swarmDiscovery,
	// 		Address:        swarmAddr,
	// 		Host:           swarmHost,
	// 		Strategy:       swarmStrategy,
	// 		ArbitraryFlags: swarmOpts,
	// 		IsExperimental: swarmExperimental,
	// 	},
	// }

}

func (d *deployer) Undeploy() (err error) {
	return
}

func (d *deployer) startHost() (err error) {
	return
}

//deploy host
func (d *deployer) deployHost(master bool, driverName string, driverOpts drivers.DriverOptions, hostName string, hostOpts *host.Options) (h *host.Host, err error) {

	rawDriver, err := json.Marshal(&drivers.BaseDriver{
		MachineName: hostName,
		StorePath:   mcndirs.GetBaseDir(),
	})

	if err != nil {
		return
	}

	h, err = d.api.NewHost(driverName, rawDriver)
	if err != nil {
		return
	}

	if hostOpts != nil {
		h.HostOptions = hostOpts
	}

	//set create config
	if driverOpts != nil {
		if err = h.Driver.SetConfigFromFlags(driverOpts); err != nil {
			return
		}
	}

	err = d.api.Create(h)
	if err != nil {
		return
	}

	err = d.api.Save(h)
	if err != nil {
		return
	}
	return
}
