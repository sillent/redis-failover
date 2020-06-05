#.PHONY: clean all build
DEST_DIR=build
COMPILE=go

all: build

clean:
	rm -rf ./$(DEST_DIR)

build: clean
	CGO_ENABLED=0 
	GOOS=linux
	mkdir -p ./$(DEST_DIR)/cmd/failover 
	$(COMPILE) build -ldflags '-extldflags "-static"' -a -v -o $(DEST_DIR)/cmd/failover ./cmd/failover/ 

docker: clean
	docker build . -f docker/Dockerfile -t failover:latest
