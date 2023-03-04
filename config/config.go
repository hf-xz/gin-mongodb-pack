package config

/*
这里是保存项目配置结构体的地方，新增配置请务必填写对应的 mapstructure 标签。
新增配置组请编写好配置的结构体，并在此引入。
*/

type Configuration struct {
	AppName string `mapstructure:"app_name"`
	Version string `mapstructure:"version"`
	Env     string `mapstructure:"env"`
	Port    string `mapstructure:"port"`

	MongoDBConfig mongoDBConfig `mapstructure:"mongodb"`
	Log           log           `mapstructure:"log"`
}
