start:
	- docker-compose up -d

before_check:
	- docker exec -it go_dic /bin/sh -c "CC_TEST_REPORTER_ID=$(CCT) ./cc-test-reporter before-build"

check:
	- docker exec -it go_dic /bin/sh -c "go test -coverprofile=c.out -v ./..."
	- docker exec -it go_dic /bin/sh -c "gocov convert c.out | gocov annotate -"

after_check:
	- docker exec -it go_dic /bin/sh -c "CC_TEST_REPORTER_ID=$(CCT) ./cc-test-reporter after-build $(TRAVIS_TEST_RESULT)"

setup:
	- docker-compose down -v
	- docker-compose build
