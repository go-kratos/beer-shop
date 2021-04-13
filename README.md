# beer-shop
An online shop application, the complete microservices demo for kratos.

## project structure
```
.
├── apps  // all kratos microservice projects
│   └── pkg  // common used packages
│   └── user
│   └── cart
│   └── payment
│   └── order
│   └── catalog
│   └── shipping
│   └── shop-gateway
│   └── admin-gateway
├── deploy  // deployment configuration
│   ├── docker-compose
│   └── kubernetes
├── docs  
└── frontend  // web ui projects
    ├── admin  // for sku management
    └── shop  // for customer
```