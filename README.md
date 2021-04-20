# [WIP] beer-shop
An online shop application, the complete microservices demo for kratos.

本项目为一个使用kratos框架创建的，简单却功能尽量完整的微服务电商项目。旨在演示kratos在mono-repo（单体仓库，即多个服务代码维护在同一git仓库中）中的项目结构实践（与layout创建出来的略有不同），以及提供多个微服务之间互相依赖调用的样例。

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
