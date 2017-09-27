BINARY=moor.bin
MOOR_DOCKER_NAME ?=moor

all: clean build docker-build

.PHONY: docker-build docker-run docker-clean build install-dependencies compile clean remove-binary install

install: all docker-run

docker-build:
		@printf "[$@] Building the container from Dockerfile. This may take a while...\n"
		docker build -t moor-image -f Dockerfile .

docker-run:
		@printf "[$@] Starting the container...\n"
		docker run -d -p 7999:7999 --name ${MOOR_DOCKER_NAME} -i moor-image

docker-clean:
		@printf "[$@] Checking if container is still saved...\n"
		@if (docker ps -a |grep ${MOOR_DOCKER_NAME}); then \
			echo "[$@] It is. Removing it... "; \
			docker container rm ${MOOR_DOCKER_NAME} ; \
		fi
		@printf "[$@] Container removed.\n"

docker-stop:
		@printf "[$@] Checking if container is running...\n"
		@if (docker ps |grep ${MOOR_DOCKER_NAME}); then \
			echo "[$@] It is. Stopping it... "; \
			docker stop ${MOOR_DOCKER_NAME}; \
		fi

build: remove-binary install-dependencies compile

install-dependencies:
		@printf "[$@] Resolving dependencies (this may take a while)...\n"
		@go get ./...

compile:
		@printf "[$@] Starting compilation (this also may take a while)...\n"
		@CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o moor.bin .
		@printf "[$@] Done! Binary: ${BINARY}\n"

clean: docker-stop docker-clean remove-binary

remove-binary:
		@printf "[$@] Checking if the binary is still here...\n" \
		@if [ -f ${BINARY} ] ; then \
			echo "[$@] It is. Removing it... "; \
			rm ${BINARY}; \
		fi \