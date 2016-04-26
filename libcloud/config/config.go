package config

import (
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
)

const (
	DefaultCloudConfigFile = "CloudConfig.yml"
)

type CloudConfig struct {
	Name         string   `yaml:"name"`
	ClusterStore string   `yaml:"ClusterStore"`
	Nodes        []string `yaml:"nodes"`
	ComposeFiles []string `yaml:"composeFiles"`
}

func NewCloudConfig(configPath string) (cc *CloudConfig, err error) {
	f, err := os.Open(configPath)
	if err != nil {
		return
	}
	defer f.Close()

	content, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}

	cc = &CloudConfig{}
	err = yaml.Unmarshal(content, cc)
	return
}
