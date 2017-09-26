BINARY=moor.bin

all:
	    go build -o ${BINARY}
install:
	    go install ${LDFLAGS} ./...
clean:
	    if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
