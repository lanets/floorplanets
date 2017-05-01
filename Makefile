all: gofmt reactapp

docker_run_node = docker run --rm -t -i -p 3000:3000 -v $$(pwd):/opt/floorplan -w /opt/floorplan -u $$(id -u):$$(id -g)
docker_run_go = docker run --rm -t -v $$(pwd):/go/src/github.com/lanets/floorplan-2 -w /go/src/github.com/lanets/floorplan-2 -u $$(id -u):$$(id -g) floorplan-golang

#######################
## FRONT-END TARGETS ##
#######################

.PHONY: nodebuild
nodebuild: node_modules
	$(docker_run_node) node npm run build

.PHONY: reactapp
reactapp: node_modules
	$(docker_run_node) node npm start

node_modules:
	$(docker_run_node) node npm install

.PHONY: nodetest
nodetest: node_modules
	$(docker_run_node) node npm test

.PHONY: nodetest-CI
nodetest-CI: node_modules
	$(docker_run_node) -e CI=true node bash -c "npm test && npm run flow"

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
test: gotest nodetest-CI

.PHONY: clean
clean:
	rm -rf node_modules
	rm -f npm-debug.log
	rm -f floorplan-api
