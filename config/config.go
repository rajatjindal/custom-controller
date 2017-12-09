package config

// ResourceType to be monitored
type ResourceType struct {
	Pod     bool
	Service bool
}

// Config to hold configurations
type Config struct {
	Resource ResourceType
}

// DefaultConfig Returns the default config
func DefaultConfig() *Config {
	return &Config{
		Resource: ResourceType{
			Pod: true,
		},
	}
}
