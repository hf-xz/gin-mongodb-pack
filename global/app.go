package global

import (
	"gin-mongodb-pack/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Application struct {
	Config      config.Configuration // 项目配置
	ConfigViper *viper.Viper         // 配置管理
	Log         *zap.Logger          // 日志系统
}

var App = new(Application)
