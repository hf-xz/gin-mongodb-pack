package config

type mongoDB struct {
	URI         string `mapstructure:"uri"`
	MaxPoolSize uint64 `mapstructure:"max_pool_size"`
}
