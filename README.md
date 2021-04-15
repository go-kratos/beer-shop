# beer-shop
An online shop application, the complete microservices demo for kratos.

## project structure
```
.
├── api // all api
├── app  // all kratos microservices projects
│   └── user-service
│   └── cart-service
│   └── payment-service
│   └── order-service
│   └── catalog-service
│   └── shipping-service
│   └── shop-interface    // BFF for customer
│   └── shop-admin   // BFF for management
├── pkg  // common used packages
├── deploy  // deployment configuration
│   ├── docker-compose
│   └── kubernetes
├── docs  
└── web  // web ui projects
    ├── admin  // for management
    └── shop  // for customer
```