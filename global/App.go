package global

import (
	"gin-mongodb-pack/config"
	"github.com/spf13/viper"
)

type Application struct {
	Config      config.Configuration // 项目配置
	ConfigViper *viper.Viper         // 配置管理
}

var App = new(Application)
