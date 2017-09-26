# This is how we want to name the binary output
BINARY=moor.bin

# These are the values we want to pass for Version and BuildTime
BUILD_TIME=`date +%FT%T%z`

all:
	    go build -o ${BINARY}
install:
	    go install ${LDFLAGS} ./...
clean:
	    if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
