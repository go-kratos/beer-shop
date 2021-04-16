# beer-shop
An online shop application, the complete microservices demo for kratos.

## project structure
```
.
├── api  // API&Error proto&generated code
│   ├── shop
│   │   └── interface
│   └── user
│       └── service
├── app  // kratos microservices projects
│   ├── shop
│   │   └── interface
│   └── user
│       └── service
├── deploy
│   ├── build
│   ├── docker-compose
│   └── kubernetes
├── docs
├── pkg  // common used packages
└── web  // web ui projects
    ├── admin  // for management
    └── shop  // for customer
```