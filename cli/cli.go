package cli

import (
	"os"
	"path"

	"github.com/codegangsta/cli"
	"github.com/xozrc/discovery/version"
)

//env
const (
	debugEnv         = "XO_CLOUD_DEBUG"
	storagePathEnv   = "XO_CLOUD_STORAGE_PATH"
	tlsCaKeyEnv      = "XO_CLOUD_TLS_CA_KEY"
	tlsCaCertEnv     = "XO_CLOUD_TLS_CA_CERT"
	tlsClientKeyEnv  = "XO_CLOUD_TLS_CLIENT_KEY"
	tlsClientCertEnv = "XO_CLOUD_TLS_CLIENT_CERT"
)

//flag variables
var (
	debugVal         bool   = true
	storagePathVal   string = ""
	tlsCaKeyVal      string = ""
	tlsCaCertVal     string = ""
	tlsClientKeyVal  string = ""
	tlsClientCertVal string = ""
)

//flags
var (
	debugFlag = cli.BoolFlag{
		Name:        "debug,D",
		Usage:       "debug mode",
		Destination: &debugVal,
		EnvVar:      debugEnv,
	}

	storagePathFlag = cli.StringFlag{
		Name:        "storage-path,s",
		Usage:       "configures storage path",
		Value:       storagePathVal,
		Destination: &storagePathVal,
		EnvVar:      storagePathEnv,
	}

	tlsCaKeyFlag = cli.StringFlag{
		Name:        "tls-ca-key",
		Usage:       "Private key to generate certificates",
		Value:       tlsCaKeyVal,
		Destination: &tlsCaKeyVal,
		EnvVar:      tlsCaKeyEnv,
	}

	tlsCaCertFlag = cli.StringFlag{
		Name:        "tls-ca-cert",
		Usage:       "CA to verify remotes against",
		Value:       tlsCaCertVal,
		Destination: &tlsCaCertVal,
		EnvVar:      tlsCaCertEnv,
	}

	tlsClientKeyFlag = cli.StringFlag{
		Name:        "tls-client-key",
		Usage:       "Private key used in client TLS auth",
		Value:       tlsClientKeyVal,
		Destination: &tlsClientKeyVal,
		EnvVar:      tlsClientKeyEnv,
	}

	tlsClientCertFlag = cli.StringFlag{
		Name:        "tls-cert-key",
		Usage:       "Client cert to use for TLS",
		Value:       tlsClientCertVal,
		Destination: &tlsClientCertVal,
		EnvVar:      tlsClientCertEnv,
	}
)

func Run() {
	app := cli.NewApp()
	app.Name = path.Base(os.Args[0])
	app.Usage = "A cloud cluster"
	app.Version = version.VERSION + " (" + version.GITCOMMIT + ")"

	app.Author = ""
	app.Email = ""

	app.Flags = []cli.Flag{
		debugFlag,
		storagePathFlag,
		tlsCaKeyFlag,
		tlsCaCertFlag,
		tlsClientKeyFlag,
		tlsClientCertFlag,
	}

	app.Commands = commands
	app.Run(os.Args)
}
