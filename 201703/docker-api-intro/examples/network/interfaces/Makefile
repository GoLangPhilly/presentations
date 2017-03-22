SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

BINARY=interfaces

BUILD_TIME=`date +%FT%T%z`

LDFLAGS=-a -tags netgo -ldflags " -w" 

.DEFAULT_GOAL: $(BINARY)

$(BINARY): $(SOURCES)
	CGO_ENABLED=0 GOOS=linux go build ${LDFLAGS} -o ${BINARY} .

mac: $(SOURCES)
	go build -o ${BINARY}-mac .

docker: clean ${BINARY}
	docker build -t robertojrojas/${BINARY} .

docker_clean: clean
	docker rmi robertojrojas/${BINARY}

.PHONY: install
install:
	go install ${LDFLAGS} ./...

.PHONY: clean
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
	if [ -f ${BINARY}-mac ] ; then rm ${BINARY}-mac ; fi
