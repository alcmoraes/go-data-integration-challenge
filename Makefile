.PHONY: help
.DEFAULT_GOAL := help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

start: ## Starts the containers and the API listening on port 8080
	- docker-compose up -d

stop: ## Stops the containers
	- docker-compose stop

before_check:
	- docker exec -it go_dic /bin/sh -c "CC_TEST_REPORTER_ID=$(CCT) ./cc-test-reporter before-build"

check: ## Run unit tests
	- docker exec -it go_dic /bin/sh -c "go test -coverprofile c.out -v ./..."

after_check:
	- docker exec -it go_dic /bin/sh -c "CC_TEST_REPORTER_ID=$(CCT) ./cc-test-reporter after-build $(TRAVIS_TEST_RESULT)"

docs: ## Serves a swagger server containing the specs for this api 
	- docker exec -it go_dic /bin/sh -c "PORT=3002 swagger serve --no-open swagger.json"

setup: remove ## Builds the containers
	- docker-compose build

remove: ## Removes containers and volumes created by docker
	- docker-compose down -v


