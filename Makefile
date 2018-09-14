.SILENT:
.ONESHELL:
.NOTPARALLEL:
.EXPORT_ALL_VARIABLES:
.PHONY: run deps build clean exec

run: build exec clean

exec:
	./bin/app

build:
	go build -o bin/app -ldflags '-s -w -extldflags "-static"'

clean:
	rm -rf bin

deps:
	govendor init
	govendor add +e
	govendor update +v

dev:
	go get -u -v github.com/kardianos/govendor

docker-build:
	docker build -t chneau .

docker-up:
	docker stack up chneau --compose-file docker-compose.yml

docker-update:
	docker service update --force chneau_chneau

docker-down:
	docker stack down chneau
