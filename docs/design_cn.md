# 设计
这是一个使用go语言和Kratos框架搭建的微服务演示项目。
项目模拟销售啤酒的电商平台

## 特性
* 展示mono-repo风格的项目结构（不同于kratos-layout创建的风格）
* 实践多个服务间通信
* 实例展示kratos和多个基础设施服务结合，例如databases，caches和消息队列mq。
* 这个项目只是模拟微服务的方案，请大胆发挥想象力。

## 组件
这一章节描述项目的各个组件。

### `/api/`
所有的 API `.proto` 文件和生成的 `.go` 文件都在此目录。
此目录结构和目录 `/app/` 类似。

### `/app/`
所有服务的源码都放在这个目录。

#### Service: `/app/catalog/service`
##### 作用/功能
此服务管理所有销售的啤酒信息。
##### 特性
* 和Ent框架集成
* 服务注册

#### Service: `/app/cart/service`
##### 作用/功能
购物车服务，管理用户将要购买的啤酒
##### 特性
* 集成MongoDB
* 服务注册

#### Service: `/app/user/service`
##### 作用/功能
用户服务，管理用户信息。
##### 特性
* 集成Ent框架
* 服务注册

#### Service: `/app/shipping/service`
##### 作用/功能
此服务模拟消息队列MQ生产者，会发送购物消息到消息队列。
##### 特性
* 集成MQ
* 服务注册

#### Job: `/app/courier/job`
##### 作用/功能
此服务模拟消息队列MQ消费者，会从消息队列中消费消息。
##### 特性
* 集成MQ

#### Service: `/app/order/service`
##### 作用/功能
订单服务，管理用户的订单。
##### 特性
* 集成GORM框架
* 服务注册

#### Service: `/app/payment/service`
##### 作用/功能
支付服务，模拟支付认证。
##### 特性
* 服务注册

#### Admin: `/app/shop/admin`
##### 作用/功能
管理端的BFF服务，管理整个商城
##### 特性
* 服务发现
* 和其他服务交互

#### Interface: `/app/shop/interface`
##### 作用/功能
管理端的BFF服务，管理整个商城
##### 特性
* 服务发现
* 和其他服务交互
* 集成缓存

### `/pkg/`
公共的包，各个服务都可以引用。

### `/deploy/`
dockerfile和部署脚本

### `/web/`
前端项目

## 架构
整个项目的架构蓝图[TBD]
