BINARY=moor.bin
MOOR_DOCKER_NAME ?= moor

all: clean compile docker-build
docker-build:
		docker build -t moor-image -f Dockerfile .
docker-run:
		docker run -d -p 7999:7999 --name ${MOOR_DOCKER_NAME} -i moor-image
docker-clean:
		docker container rm ${MOOR_DOCKER_NAME}
docker-stop:
		docker stop ${MOOR_DOCKER_NAME}
compile:
		CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o moor.bin .
clean: docker-stop docker-clean remove-binary
remove-binary:
		if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
