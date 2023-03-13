# 开发指南

本文将带你了解如何使用快速本框架进行业务开发。

## 项目结构

| 目录        | 功能        |
|-----------|-----------|
| config    | 存放项目配置结构体 |
| global    | 存放全局变量    |
| bootstrap | 项目启动逻辑    |
| utils     | 项目工具      |
| models    | 存放业务模型    |

## 项目配置

项目配置文件路径在项目[根目录](./)，存放配置结构体的路径为 [config](./config) 。

### 选择配置

项目默认为开发环境，在生产部署时，请设置环境变量 `GO_ENV=prod`。

如需添加、修改环境配置文件，请在 [config.go](./bootstrap/config.go) 中进行修改。

### 添加配置项

在 config 目录中新建所需的配置结构体，在 [Config.go](./config/Config.go) 中引入即可。注意，务必为所添加的配置选项填写对应的 `mapstructure` 标签。

## 数据库访问

本项目默认使用mongoDB的单数据库，如有多数据库使用的需要，请自行在[app.go](./global/app.go)中自行添加，并在[mongoDB.go](./bootstrap/mongoDB.go)中修改初始化函数。

本框架不对数据库操作做日志记录，请在数据库端配置操作日志记录。

[//]: # (TODO)