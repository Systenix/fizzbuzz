# FizzBuzz Microservice example

This is a production-level FizzBuzz microservice generated using [go-cloud](https://github.com/Systenix/go-cloud.git). By leveraging a blueprint, we quickly scaffolded the service and updated it to meet production standards, complete with monitoring and metrics.

## Table of Contents

- [Introduction](#introduction)
- [Setup and Installation](#setup-and-installation)
  - [Prerequisites](#prerequisites)
  - [Clone the Repository](#clone-the-repository)
  - [Running the Service](#running-the-service)
- [Project Structure](#project-structure)
- [Using go-cloud for Rapid Development](#using-go-cloud-for-rapid-development)
  - [The Blueprint](#the-blueprint)
  - [Generation Process](#generation-process)
  - [Benefits of go-cloud](#benefits-of-go-cloud)
- [Endpoints](#endpoints)
  - [/fizzbuzz](#fizzbuzz)
  - [/statistics](#statistics)
- [Monitoring with Prometheus and Grafana](#monitoring-with-prometheus-and-grafana)
  - [Accessing Prometheus and Grafana](#accessing-prometheus-and-grafana)
  - [Visualizing Metrics](#visualizing-metrics)

## Introduction

This microservice implements the classic FizzBuzz problem as a RESTful API and tracks the most frequent requests. It was generated using go-cloud, showcasing how quickly a production-ready service can be scaffolded and customized using a blueprint.

## Setup and Installation

### Prerequisites

	•	Docker and Docker Compose installed.
	•	Git installed.
	•	Go installed.

### Clone the Repository
```bash
git clone https://github.com/Systenix/fizzbuzz.git
cd fizzbuzz
```

### Running the Service

To build and run the service along with Redis, Prometheus, and Grafana:
```bash
docker-compose up -d
```

This command will start:
- fizzbuzz: The FizzBuzz microservice on port 8080.
- redis: For storing request statistics.
- prometheus: For scraping metrics from the microservice.
- grafana: For visualizing metrics.

To stop the services:
```bash
docker-compose down
```

## Project Structure

```
fizzbuzz/
├── Dockerfile
├── Makefile
├── bin/
├── cmd/
│   └── main.go
├── docker-compose.yml
├── go.mod
├── internal/
│   ├── infrastructures/
│   │   ├── redis/
│   │   │   └── redis.go
│   │   └── repositories/
│   │       └── statisticsrepository.go
│   ├── interfaces/
│   │   ├── handlers/
│   │   │   └── fizzbuzzhandler.go
│   │   └── middleware/
│   │       └── middleware.go
│   ├── models/
│   │   ├── fizzbuzzrequest.go
│   │   ├── fizzbuzzresponse.go
│   │   └── statsresponse.go
│   └── services/
│       └── fizzbuzzservice.go
├── prometheus/
│   └── prometheus.yml
└── grafana/
    └── provisioning/
        └── datasources/
            └── datasource.yml
```

- cmd/: Entry point of the application.
- internal/: Application logic organized by domain.
- Dockerfile: Docker build configuration.
- docker-compose.yml: Configuration to run services together.

## Using go-cloud for Rapid Development

### The Blueprint

We defined the service as per the requirements using the following blueprint:

```yaml
go_version: "1.22.2"
port: 8080

services:
  - name: "FizzBuzzService"
    type: "rest"
    models:
      - "FizzBuzzRequest"
    repositories:
      - "StatisticsRepository"
    handlers:
      - name: "FizzBuzzHandler"
        service: "FizzBuzzService"
        routes:
          - path: "/fizzbuzz"
            verb: "GET"
            method: "FizzBuzz"
            request_model: "FizzBuzzRequest"
            response_model: "FizzBuzzResponse"
            middleware:
              - "MetricsMiddleware"
          - path: "/statistics"
            verb: "GET"
            method: "GetStatistics"
            response_model: "StatsResponse"
    methods:
      - name: "FizzBuzz"
        params:
          - name: "ctx"
            type: "context.Context"
          - name: "request"
            type: "*models.FizzBuzzRequest"
        returns:
          - name: "result"
            type: "[]string"
          - name: "err"
            type: "error"
      - name: "GetStatistics"
        params:
          - name: "ctx"
            type: "context.Context"
        returns:
          - name: "result"
            type: "*models.StatsResponse"
          - name: "err"
            type: "error"

middleware:
  - name: "LoggingMiddleware"
    type: "logging"
    scope: "global"
  - name: "RecoveryMiddleware"
    type: "recovery"
    scope: "global"
  - name: "MetricsMiddleware"
    type: "metrics"
    scope: "route"
    options:
      endpoint: "/metrics"

repositories:
  - name: "StatisticsRepository"
    type: "redis"
    model: "StatsResponse"
    settings:
      address: "redis:6379"
      password: ""
      db: 0
    methods:
      - name: "Increment"
        params:
          - name: "ctx"
            type: "context.Context"
          - name: "request"
            type: "*models.FizzBuzzRequest"
        returns:
          - name: "err"
            type: "error"
      - name: "GetMostFrequent"
        params:
          - name: "ctx"
            type: "context.Context"
        returns:
          - name: "result"
            type: "*models.StatsResponse"
          - name: "err"
            type: "error"

models:
  - name: FizzBuzzRequest
    fields:
      - name: int1
        type: int
        json_name: int1
      - name: int2
        type: int
        json_name: int2
      - name: limit
        type: int
        json_name: limit
      - name: str1
        type: string
        json_name: str1
      - name: str2
        type: string
        json_name: str2
  - name: FizzBuzzResponse
    fields:
      - name: "result"
        type: "[]string"
  - name: StatsResponse
    fields:
      - name: "most_frequent"
        type: "*FizzBuzzRequest"
      - name: "hits"
        type: "int"

docker:
  enabled: true

third_party:
  prometheus:
    enabled: true
    port: "9090"
  grafana:
    enabled: true
    port: "3000"
```

### Generation Process

By running [go-cloud](https://github.com/Systenix/go-cloud.git) with this blueprint, the tool generated:
- The project directory structure.
- Boilerplate code for models, services, handlers, repositories, and middleware and main files.
- Configuration files for Docker, Docker Compose, Prometheus, and Grafana.

### Benefits of go-cloud

- Rapid Development: Generated a functional scaffold in minutes.
- Consistency: Ensured a consistent code structure and coding standards.
- Customization: Allowed us to focus on business logic and customize where needed.
- Scalability: Made it easy to extend the service with additional features.
- Integration: Included monitoring tools out-of-the-box.

## Endpoints

### /fizzbuzz

- Method: GET
- Description: Generates a FizzBuzz sequence based on provided parameters.
- Query Parameters:
  - int1 (int): Divisor for the first word substitution.
  - int2 (int): Divisor for the second word substitution.
  - limit (int): Upper limit of the sequence.
  - str1 (string): Word to replace multiples of int1.
  - str2 (string): Word to replace multiples of int2.
- Example Request:

```
GET /fizzbuzz?int1=3&int2=5&limit=15&str1=Fizz&str2=Buzz
```

- Example Response:

```json
{
  "result":["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]
}
```

### /statistics

- Method: GET
- Description: Retrieves the most frequent /fizzbuzz request parameters and the number of times it was called.
- Example Response:

```json
{
  "most_frequent": {
    "int1": 3,
    "int2": 5,
    "limit": 15,
    "str1": "Fizz",
    "str2": "Buzz"
  },
  "hits": 42
}
```

## Monitoring with Prometheus and Grafana

The service includes monitoring capabilities using Prometheus and Grafana to track metrics such as request rates, durations, and error rates.

### Accessing Prometheus and Grafana

- Prometheus UI: http://localhost:9090
- Grafana UI: http://localhost:3000
  - Username: admin
  - Password: admin (please change this in a production environment, this is only for development purposes)

### Visualizing Metrics

- Dashboards: Create custom dashboards in Grafana to visualize metrics.
- Metrics Exposed:
  - http_requests_total: Total number of HTTP requests.
  - http_request_duration_seconds: Duration of HTTP requests.
  - Custom application metrics can be added as needed.
