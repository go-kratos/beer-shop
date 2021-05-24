# [WIP] beer-shop
An online shop application, mono-repo microservices demo for kratos.

本项目为一个使用kratos框架创建的，简单的微服务电商项目。

由于example较碎片化，未能体现出完整的项目的样子，因此我们创建了本项目，以完成如下目标：

* 演示kratos在mono-repo中的项目结构实践（与layout创建出来的略有不同）
* 提供多个微服务之间互相依赖调用和样例
* 提供与各种基础设施集成和部署的样例
* 主要为kratos框架使用演示，很多组件的设计做了简化或模拟处理，与实际的电商项目有一定出入，仅供参考

具体架构请参考文档：[Docs](https://go-kratos.github.io/beer-shop/#/)

**ATTENTION: This project is a Work-in-Progress.**

**注意，目前尚在开发，暂时无法运行，仅供代码参考。**

## Kratos Mono-Repo structure
```
.
├── api  // API&Error Proto files & Generated codes
│   ├── foo
│   │   ├── job
│   │   └── service
│   └── bar
│       └── interface
├── app  // kratos microservices projects
│   ├── foo
│   │   ├── job
│   │   └── service
│   └── bar
│       └── interface
├── pkg  // common used packages
└── docs

```
