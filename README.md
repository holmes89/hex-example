# Gira (A Hex Example)

This is a sample project that was originally supposed to be and example of (hexagonal architecture)[http://www.joeldholmes.com/post/go-hex-arch/]. It has slowly evolved to incorporate other ideas and product integrations. 

Gira is a _really_ simple ticket management API. You can create and fetch ticket (that's it). The application is flexible in that it can run as either a AWS Lambda function or a standalong server. It also allows you to swap backend database (between Redis or Postgres). 

This project will continue to evolve and change depending on things I'm interested in learning.

## Explored Areas
* (Hexagonal Architecture)[http://www.joeldholmes.com/post/go-hex-arch/]
* (Docker)[http://www.joeldholmes.com/post/go-docker/]
* (FaaS)[http://www.joeldholmes.com/post/serverless-to-server/]
* (Mock Testing)[http://www.joeldholmes.com/post/go-mock-testing/]
* Terraform
* Jenkins Pipeline
* Robot Acceptance Testing

## Build

`go build main.go`

## Test

`go test ./...`

## Run

`./main` - runs in lambda mode with a redis instance

### Flags

`--database` - either `redis` or `psql`
`--server` - runs in sever mode

### Environment

`DATABASE_URL` - Override local database url
`REDIS_PASSWORD` - Override default Redis Password

## Run Acceptance Tests

`cd tests/acceptance`
`source .env`
`robot -v HOST:${endpoint} tickets.robot`
