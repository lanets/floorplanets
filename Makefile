all: gulp gofmt

docker_run_node = docker run --rm -t -v $$(pwd):/opt/floorplan -w /opt/floorplan -u $$(id -u):$$(id -g) node
docker_run_go = docker run --rm -t -v $$(pwd):/go/src/github.com/lanets/floorplan-2 -w /go/src/github.com/lanets/floorplan-2 -u $$(id -u):$$(id -g) floorplan-golang


#######################
## FRONT-END TARGETS ##
#######################

.PHONY: nodebuild
nodebuild: node_modules
	$(docker_run_node) node node_modules/gulp/bin/gulp.js build

.PHONY: gulp
gulp: node_modules
	$(docker_run_node) node node_modules/gulp/bin/gulp.js

node_modules:
	$(docker_run_node) npm install typescript gulp
	$(docker_run_node) npm install

.PHONY: nodetest
nodetest: node_modules
	$(docker_run_node) npm test

#####################
## BACKEND TARGETS ##
#####################

.PHONY: golang-build-image
golang-build-image:
	docker build --build-arg userid=$$(id -u) -f docker/golang -t floorplan-golang docker

.PHONY: gobuild
gobuild: golang-build-image
	$(docker_run_go) go build ./...

.PHONY: gofmt
gofmt: golang-build-image
	$(docker_run_go) go fmt ./...

.PHONY: gotest
gotest: golang-build-image
	$(docker_run_go) bash -c "go get -v -t ./... && go test ./..."


#####################
## GENERAL TARGETS ##
#####################

.PHONY: build
build: nodebuild gobuild

.PHONY: test
test: gotest nodetest

.PHONY: clean
clean:
	rm -rf node_modules
	rm -f npm-debug.log
	rm -f floorplan-api
