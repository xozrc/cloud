package cloud

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/docker/machine/libmachine"
	"github.com/docker/machine/libmachine/drivers"
	"github.com/docker/machine/libmachine/host"
	"github.com/docker/machine/libmachine/mcnerror"
	"golang.org/x/net/context"
)

type Deployer interface {
	Deploy() error
	Undeploy() error
}

func NewDeployer(ctx context.Context) Deployer {
	return &deployer{ctx}
}

type deployer struct {
	ctx context.Context
}

func (dd *deployer) Deploy() (err error) {
	log.Println("ads")
	hostName := "asd"
	driver := virtualbox.NewDriver(hostName, "/tmp/automatic")
	driver.CPU = 2
	driver.Memory = 2048

	data, err := json.Marshal(driver)
	if err != nil {
		return
	}

	_, err = createHost("virtualbox", data, hostName, "", "")
	return
}

func (dd *deployer) Undeploy() error {
	return nil
}

//create Host
func createHost(driver string, rawDriver []byte, hostName, storePath, certsDir string) (h *host.Host, err error) {

	api := libmachine.NewClient(storePath, certsDir)

	validName := host.ValidateHostName(hostName)
	if !validName {
		return fmt.Errorf("Error creating machine: %s", mcnerror.ErrInvalidHostname)
	}

	// TODO: Fix hacky JSON solution
	rawDriver, err := json.Marshal(&drivers.BaseDriver{
		MachineName: name,
		StorePath:   storePath,
	})

	if err != nil {
		return fmt.Errorf("Error attempting to marshal bare driver data: %s", err)
	}

	h, err := api.NewHost(driverName, rawDriver)
	if err != nil {
		return fmt.Errorf("Error getting new host: %s", err)
	}

	h.HostOptions = &host.Options{
		AuthOptions: &auth.Options{
			CertDir:          mcndirs.GetMachineCertDir(),
			CaCertPath:       tlsPath(c, "tls-ca-cert", "ca.pem"),
			CaPrivateKeyPath: tlsPath(c, "tls-ca-key", "ca-key.pem"),
			ClientCertPath:   tlsPath(c, "tls-client-cert", "cert.pem"),
			ClientKeyPath:    tlsPath(c, "tls-client-key", "key.pem"),
			ServerCertPath:   filepath.Join(mcndirs.GetMachineDir(), name, "server.pem"),
			ServerKeyPath:    filepath.Join(mcndirs.GetMachineDir(), name, "server-key.pem"),
			StorePath:        filepath.Join(mcndirs.GetMachineDir(), name),
			ServerCertSANs:   c.StringSlice("tls-san"),
		},
		EngineOptions: &engine.Options{
			ArbitraryFlags:   c.StringSlice("engine-opt"),
			Env:              c.StringSlice("engine-env"),
			InsecureRegistry: c.StringSlice("engine-insecure-registry"),
			Labels:           c.StringSlice("engine-label"),
			RegistryMirror:   c.StringSlice("engine-registry-mirror"),
			StorageDriver:    c.String("engine-storage-driver"),
			TLSVerify:        true,
			InstallURL:       c.String("engine-install-url"),
		},
		SwarmOptions: &swarm.Options{
			IsSwarm:        c.Bool("swarm"),
			Image:          c.String("swarm-image"),
			Master:         c.Bool("swarm-master"),
			Discovery:      c.String("swarm-discovery"),
			Address:        c.String("swarm-addr"),
			Host:           c.String("swarm-host"),
			Strategy:       c.String("swarm-strategy"),
			ArbitraryFlags: c.StringSlice("swarm-opt"),
			IsExperimental: c.Bool("swarm-experimental"),
		},
	}

	exists, err := api.Exists(h.Name)
	if err != nil {

		return
	}

	if exists {
		err = mcnerror.ErrHostAlreadyExists{
			Name: h.Name,
		}
		return
	}

	// driverOpts is the actual data we send over the wire to set the
	// driver parameters (an interface fulfilling drivers.DriverOptions,
	// concrete type rpcdriver.RpcFlags).
	mcnFlags := h.Driver.GetCreateFlags()
	driverOpts := getDriverOpts(c, mcnFlags)

	if err := h.Driver.SetConfigFromFlags(driverOpts); err != nil {
		return fmt.Errorf("Error setting machine configuration from flags provided: %s", err)
	}

	if err := api.Create(h); err != nil {
		// Wait for all the logs to reach the client
		time.Sleep(2 * time.Second)

		vBoxLog := ""
		if h.DriverName == "virtualbox" {
			vBoxLog = filepath.Join(api.GetMachinesDir(), h.Name, h.Name, "Logs", "VBox.log")
		}

		err = fmt.Errorf("crash")
		return
	}

	if err = api.Save(h); err != nil {
		return nil, err
	}

	log.Infof("To see how to connect your Docker Client to the Docker Engine running on this virtual machine, run: %s env %s", os.Args[0], name)

	return h, nil
}
