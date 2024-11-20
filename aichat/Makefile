.PHONY: build

build:
	cd backend && go build -o aichatapp ./cmd/aichat
	cd backend && cp aichatapp ../build

run:
	cd backend && go mod tidy
	cd backend && go build -o main ./cmd/aichat
	docker-compose up

deploy: build
	cd backend && scp aichatapp root@68.183.225.233:/root