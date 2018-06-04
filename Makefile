start:
	- docker-compose up -d

before_check:
	- docker exec -it go_dic /bin/sh -c "CC_TEST_REPORTER_ID=$(CCT) ./cc-test-reporter before-build"

check:
	- docker exec -it go_dic /bin/sh -c "go test -coverprofile c.out ./..."

after_check:
	- docker exec -it go_dic /bin/sh -c "CC_TEST_REPORTER_ID=$(CCT) ./cc-test-reporter after-build $(TRAVIS_TEST_RESULT)"

docs:
	- docker exec -it go_dic /bin/sh -c "PORT=3002 swagger serve swagger.json"

setup:
	- docker-compose down -v
	- docker-compose build
