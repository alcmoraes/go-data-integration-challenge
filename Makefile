#Challenge Makefile

start:
	- docker-compose up -d

check:
	- docker exec -it go_dic bash -c "mv config.yml config.yml.bck && mv config.yml.test config.yml"
	- docker exec -it go_dic bash -c "go test -v ./..."
	- docker exec -it go_dic bash -c "mv config.yml config.yml.test && mv config.yml.bck config.yml"

setup:
	- docker-compose build
