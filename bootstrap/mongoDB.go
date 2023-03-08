package bootstrap

import (
	"context"
	"gin-mongodb-pack/global"
	"go.mongodb.org/mongo-driver/mongo"
	options2 "go.mongodb.org/mongo-driver/mongo/options"
)

func InitializeMongoDB() *mongo.Client {
	// 获取数据库配置
	conf := global.App.Config.MongoDB

	// 配置数据库
	o := options2.Client()
	o.ApplyURI(conf.URI)
	o.SetMaxPoolSize(conf.MaxPoolSize)

	// 尝试连接数据库
	global.App.Log.Info("connecting mongoDB on ", conf.URI)
	client, err := mongo.Connect(context.TODO(), o)
	if err != nil {
		global.App.Log.Fatal("mongoDB connect failed! ", err)
	}

	// 测试连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		global.App.Log.Fatal("mongoDB connect failed! ", err)
	}

	/*
		返回的客户端是池化过的，可以直接用于并发的访问，文档描述如下：
		Client is a handle representing a pool of connections to a MongoDB deployment.
		It is safe for concurrent use by multiple goroutines.
	*/
	return client
}

func CloseMongoDB() {
	if global.App.DB == nil {
		return
	}
	if err := global.App.DB.Disconnect(context.TODO()); err != nil {
		global.App.Log.Fatal("mongoDB disconnect failed! ", err)
	}
	global.App.Log.Info("mongoDB disconnect success.")
}
