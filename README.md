# Product Consumer

## App Overview

This application keeps record of customers information obtained from NetSuite (and other sources) through Pub/Sub and provide end points to get the information of customers.

## Getting Started

1. Install go (version 1.9.2)
2. Install [Google Cloud SDK](https://cloud.google.com/sdk/)
  * Minimum permissions:
      * KE Developer
      * Pub/Sub Editor
3. Clone this repository in $GOPATH/src
4. Export environment variables
5. Setup to use GCP user credentail
6. `go run customer_consumer.go`
7. Visit `localhost:8082`

### Dependency Management

We use [dep](https://github.com/golang/dep) to manage dependencies. Check out [Daily Dep](https://golang.github.io/dep/docs/daily-dep.html) on how to use `dep ensure` for day-to-day use.

### Environment Variables

| Variables | Dev | Prod |
|-----------|----------|-----------|
| GOOGLE_CLOUD_PROJECT | chefd-dev-190417| |
|RAW_CUSTOMER_DATA|raw_customer_data||
|GET_RAW_CUSTOMER_DATA|get_raw_customer_data||
|PROCESSED_CUSTOMER_DATA|processed_customer_data||
|GET_PROCESSED_CUSTOMER_DATA|get_processed_customer_data||
|DATA_STORE_KIND_CUSTOMER|CUSTOMER_KIND|| 
|DATA_STORE_NAME_CUSTOMER|CUSTOMER| 

## Containerization

1. Update Dockerfile (need to include env variables)
2. `docker build -t customer-consumer:<tag> .`
3. `docker run -it -p 8082:<port> customer-consumer:<tag>`

## Setting Up Pub/Sub Emulator

1. Install [Pub/Sub Emulator](https://cloud.google.com/pubsub/docs/emulator)
2. Setup environment variables
2. Create Topic
3. Subscribe

## Deploy

TBD