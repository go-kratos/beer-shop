# beer-shop
An online shop application, the complete microservices demo for kratos.

## project structure
```
.
├── project  // all kratos microservices projects
│   └── pkg  // common used packages
│   └── user-service
│   └── cart-service
│   └── payment-service
│   └── order-service
│   └── catalog-service
│   └── shipping-service
│   └── shop-interface    // BFF for customer
│   └── shop-admin   // BFF for management
├── deploy  // deployment configuration
│   ├── docker-compose
│   └── kubernetes
├── docs  
└── frontend  // web ui projects
    ├── admin  // for management
    └── shop  // for customer
```