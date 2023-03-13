# 项目介绍

这是对 Golang 的 Gin 框架的一个简单封装，用于开发后端api，旨在创建一个开箱即用的基础框架。

技术栈为 Golang + Gin + MongoDB。

## 项目功能

本项目将会对以下功能进行封装：

- [x] 全局配置（如日志等级、后端地址等）
  - 基于 Viper 库实现
  - 多配置文件（开发环境，生产环境）
- [x] 日志系统
- [ ] MongoDB 数据库访问
  - 数据库模型管理
  - 数据库操作封装
- [ ] 登录接口和 jwt 鉴权中间件
- ...

## 项目结构

| 目录        | 功能        |
|-----------|-----------|
| config    | 存放项目配置结构体 |
| global    | 存放全局变量    |
| bootstrap | 项目启动逻辑    |
| utils     | 项目工具      |
| models    | 存放业务模型    |

## 相关资源

### 推荐IDE插件

| 用途                | GoLand               | VSCode          |
|-------------------|----------------------|-----------------|
| 高亮注释，安装以更好查阅本项目代码 | Comments Highlighter | Better Comments |

### 主要插件列表

| 名称         | 地址                                  | 说明     |
|------------|-------------------------------------|--------|
| gin        | `github.com/gin-gonic/gin`          | web框架  |
| viper      | `github.com/spf13/viper`            | 配置管理   |
| zap        | `go.uber.org/zap`                   | 日志管理   |
| lumberjack | `gopkg.in/natefinch/lumberjack.v2`  | 日志切割归档 |
| mongo      | `go.mongodb.org/mongo-driver/mongo` | 数据库驱动  |

### 参考链接

- [手把手，带你从零封装Gin框架 @jassue](https://juejin.cn/post/7016742808560074783)
- [Go语言配置管理神器——Viper中文教程](https://zhuanlan.zhihu.com/p/272508571)
- [mongodb驱动文档](https://www.mongodb.com/docs/drivers/go/current/)