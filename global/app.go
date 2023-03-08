package global

import (
	"gin-mongodb-pack/config"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Application struct {
	Config      config.Configuration // 项目配置
	ConfigViper *viper.Viper         // 配置管理
	Log         *zap.SugaredLogger   // 日志系统
	DB          *mongo.Client        // 数据库
}

var App = new(Application)
