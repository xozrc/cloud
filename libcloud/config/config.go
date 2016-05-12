package config

import (
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
)

type CloudConfig struct {
	Name  string   `yaml:"name"`
	Nodes []string `yaml:"nodes"`
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
