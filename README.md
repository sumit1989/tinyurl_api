# tinyurl_api: Repository for shorten Url
## Download source

    cd $GOPATH/src
    git clone --recursive git@git.idfcbank.com:optimus/risk-monitor-api.git

## Prerequisites

* You need docker installed
* Docker setup
* Docker-compose
* Golang compiler version go1.12 with go mod enabled setup
* Make sure you have proxy setup

## Run with `go run main.go`

## Run Test case `go test ./... -coverprofile coverage.out`

## Run Integration Test Case

* cd integrationtest/
* go test -tags integration