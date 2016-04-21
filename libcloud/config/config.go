package config

const (
	CloudConfigFile = "CloudConfig.yml"
)

type CloudConfig struct {
	NodesNum     int      `yaml:"nodesNum"`
	ClusterStore string   `yaml:"cluterStore"`
	ComposeFiles []string `yaml:"composeFiles"`
}
