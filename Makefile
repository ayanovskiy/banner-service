.PHONY: build
build:
	CGO_ENABLED=0 go build -v ./cmd/app.go

.PHONY: test
test:
	go test ./... -v -race -count=10

.PHONY: lint
lint:
	golangci-lint run

.PHONY: docker-build
docker-build:
	docker-compose -f ./.docker/docker-compose.yaml build

.PHONY: docker-up
docker-up:
	docker-compose -f ./.docker/docker-compose.yaml up -d --remove-orphans

.PHONY: docker-down
docker-down:
	docker-compose -f ./.docker/docker-compose.yaml down --remove-orphans
