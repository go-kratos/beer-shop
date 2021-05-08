# Design
This project is a demonstration for building microservices system with golang and Kratos framework. 
It simulates an e-commerce website that sells beers.

## Features
* Shows the project layout in mono-repo. (Which is different from the structures created with kratos-layout) 
* Presents the practice of multiple services' communication.
* Examples for kratos integrated with many infrastructure services such as databases, caches and message queues.
* This project is just a toy for simulating microservices scenario, feel free to play with it

## Components
this section describes the components in this repository.

### `/api/`
All the API `.proto` files and generated `.go` files are in this directory.
The directory structure is same as `/app/`.

### `/app/`
All the actual services source codes are located in there.

#### Service: `/app/catalog/service`
##### Functions
This service maintained all the beers which are selling in this shop.
##### Features
* Integration for Ent
* Service Registration

#### Service: `/app/cart/service`
##### Functions
The cart service, which can store the beers that users want to buy.
##### Features
* Integration for MongoDB
* Service Registration

#### Service: `/app/user/service`
##### Functions
The user service, which holds the users' information.
##### Features
* Integration for Ent
* Service Registration

#### Service: `/app/shipping/service`
##### Functions
This service stimulates a MQ producer. It will put the shipping package messages to the message queue.
##### Features
* Integration for MQ
* Service Registration

#### Job: `/app/courier/job`
##### Functions
This service stimulates a MQ consumer. It will receive(consume) the messages from the message queue.
##### Features
* Integration for MQ

#### Service: `/app/order/service`
##### Functions
The order service, which holds the users' order.
##### Features
* Integration for GORM
* Service Registration

#### Service: `/app/payment/service`
##### Functions
Just a stimulation of payment authentication.
##### Features
* Service Registration

#### Admin: `/app/shop/admin`
##### Functions
The backend for frontend(BFF) service for Administrator Web UI, to manage the shop.
##### Features
* Service Discovery 
* Communication with other services

#### Interface: `/app/shop/interface`
##### Functions
The backend for frontend(BFF) service for Administrator Web UI, to manage the shop.
##### Features
* Service Discovery 
* Communication with other services
* Integration for Cache

### `/pkg/`
The common packages which used by many services. 

### `/deploy/`
The dockerfile and deployment scripts.

### `/web/`
The frontend project.

## Architecture
This is a picture of the whole architecture.
[TBD]
