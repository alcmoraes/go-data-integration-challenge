# Data Integration Challenge

[![Code Climate](https://codeclimate.com/github/alcmoraes/go-data-integration-challenge/badges/gpa.svg)](https://codeclimate.com/github/alcmoraes/go-data-integration-challenge)
[![GoReport](https://goreportcard.com/badge/github.com/alcmoraes/go-data-integration-challenge)](https://goreportcard.com/report/github.com/alcmoraes/go-data-integration-challenge)
[![Travis CI](https://api.travis-ci.org/alcmoraes/go-data-integration-challenge.svg?branch=master)](https://travis-ci.org/alcmoraes/go-data-integration-challenge)

**Obs.: That's a technical challenge applied by @alcmoraes to Neoway**

**Obs².: Refer to [data-integration-challenge](https://github.com/alcmoraes/data-integration-challenge) for the Node version of this challenge**

## Requirements

1. Docker
2. Docker Compose
3. Ports 3002 (for swagger), 8080 (for api) and 27017 (for mongo)

*You can change the exposed ports on `docker-compose.yml` in case you have no interest on exposing `mongo` or `swagger` ports. It's your choice*

## Features

It will build a mongo and go docker containers, where the go container `automatically starts listening` on port `8080`
CSV's uploaded via API imported/merged into database right away.

## The Goroutine issue

For some reason, unit tests that executes goroutines are not working.
Unfortunately I couldn't go further on debugging that issue.
To meet the deadline, I've chosen to let those goroutines as synchronous.
*I'm not proud of this decision*

## Usage

The default usage for this project is:

```
make setup
make start
```

That will start the API on port `8080`

In the other hand, typing `make` show the available commands in your terminal

| Command   |      Description     |
|-----------|:---------------------|
| help      | Gets this table in your terminal | 
| setup     | Builds the containers |
| start     | Starts the containers and the API listening on port `8080` |
| stop      | Stops the containers |
| check     | Run unit tests |
| docs      | Serve the Swagger Explorer UI on port `3002` |
| remove    | Removes containers and volumes created by docker |

## API Documentation (Swagger)

Ensure the project is running (`make start`).

Execute `make docs`, this will start a swagger explorer on port `3002`.

