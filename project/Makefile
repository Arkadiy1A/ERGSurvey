FRONT_END_BINARY=frontApp
BACK_END_BINARY=backApp

build_all: build_front build_back

## build_front: builds the front end binary
build_front:
	@echo "Building front end binary..."
	cd ../front && env GOOS=linux CGO_ENABLED=0 go build -o ${FRONT_END_BINARY} ./cmd/web
	@echo "Done!"

## build_back: builds the back end binary
build_back:
	@echo "Building front end binary..."
	cd ../back && env GOOS=linux CGO_ENABLED=0 go build -o ${BACK_END_BINARY} ./app
	@echo "Done!"

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_all
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

docker_deploy_all: docker_deploy_front docker_deploy_back docker_deploy_caddy

docker_deploy_front: build_front
	cd ../front && docker build -f erg.survey.front.dockerfile -t arkadiyagapov/erg-survey-front:1.0.0 .
	docker push arkadiyagapov/erg-survey-front:1.0.0

docker_deploy_back: build_back
	cd ../back && docker build -f erg.survey.back.dockerfile -t arkadiyagapov/erg-survey-back:1.0.0 .
	docker push arkadiyagapov/erg-survey-back:1.0.0

docker_deploy_caddy:
	cd ../project && docker build -f caddy.dockerfile -t arkadiyagapov/erg-survey-caddy:1.0.0 .
	docker push arkadiyagapov/erg-survey-caddy:1.0.0
