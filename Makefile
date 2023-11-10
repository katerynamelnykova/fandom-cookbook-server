help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

db_clean: ## Stop and remove the mongo container and the image
	-docker stop fandom-cookbook_mongodb
	-docker rm fandom-cookbook_mongodb
	-docker rmi fandom-cookbook/mongodb

db_build: ## Build mongo image
	docker build -f Dockerfile -t fandom-cookbook/mongodb .
	echo "`docker images | grep fandom-cookbook`"

db_run: db_clean db_build ## rebuild and run the mongo image
	docker run -d -p 27017:27017 --name fandom-cookbook_mongodb fandom-cookbook/mongodb
	echo "Point your DB client to localhost:27017 to connect to this DB"

build: ## Build the service in a temp directory
	echo "building the fandom-cookbook service"
	go build -o /tmp/fandom-cookbook

run: build ## build and run service
	echo "running the service"
	/tmp/fandom-cookbook


.PHONY: docker_run docker_build
