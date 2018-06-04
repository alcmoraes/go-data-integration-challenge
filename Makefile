start:
	- docker-compose up -d

check:
	- docker exec -it go_dic /bin/sh -c "go test -coverprofile coverage.out -v ./..."

lcov:
	- docker exec -it go_dic /bin/sh -c "cat coverage.out"

setup:
	- docker-compose build
