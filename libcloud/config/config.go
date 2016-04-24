package config

const (
	CloudConfigFile = "CloudConfig.yml"
)

type CloudConfig struct {
	Driver       string   `yaml:"driver"`
	NodesNum     int      `yaml:"nodesNum"`
	ClusterStore string   `yaml:"cluterStore"`
	ComposeFiles []string `yaml:"composeFiles"`
}

type ClusterConfig struct {
}

func NewClusterConfig(clusterPath string) (cc *ClusterConfig, err error) {
	return
}
