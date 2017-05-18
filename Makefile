.PHONY: all
all: gofmt reactapp

docker_run_node = docker run --rm -t -i -v $$(pwd):/opt/floorplan -w /opt/floorplan -u $$(id -u):$$(id -g)
docker_run_go = docker run --rm -t -v $$(pwd):/go/src/github.com/lanets/floorplanets -w /go/src/github.com/lanets/floorplanets -u $$(id -u):$$(id -g) floorplan-golang

#######################
## FRONT-END TARGETS ##
#######################

.node-build-image:
	docker build --build-arg userid=$$(id -u) -f docker/node -t floorplan-node docker
	touch .node-build-image

.PHONY: nodebuild
nodebuild: node_modules .node-build-image
	$(docker_run_node) floorplan-node npm run build

flow-typed: node_modules .node-build-image
	$(docker_run_node) floorplan-node npm run flowtyped

.PHONY: reactapp
reactapp: node_modules flow-typed .node-build-image
	$(docker_run_node) -p 3000:3000 floorplan-node npm start

node_modules: .node-build-image
	$(docker_run_node) floorplan-node npm install

.PHONY: nodetest
nodetest: node_modules .node-build-image
	$(docker_run_node) floorplan-node npm test

.PHONY: eslint
eslint: node_modules .node-build-image
	$(docker_run_node) floorplan-node ./node_modules/eslint/bin/eslint.js src

.PHONY: flow
flow: node_modules .node-build-image flow-typed
	$(docker_run_node) -e CI=true floorplan-node npm run flow

.PHONY: nodetest-CI
nodetest-CI: node_modules .node-build-image eslint flow
	$(docker_run_node) -e CI=true floorplan-node npm test

#####################
## BACKEND TARGETS ##
#####################

.golang-build-image:
	docker build --build-arg userid=$$(id -u) -f docker/golang -t floorplan-golang docker
	touch .golang-build-image

.PHONY: gobuild
gobuild: .golang-build-image
	$(docker_run_go) bash -c "go get -v ./... && go build ./... && go build ./.../cmd/..."

.PHONY: gofmt
gofmt: .golang-build-image
	$(docker_run_go) go fmt ./...

.PHONY: gotest
gotest: .golang-build-image
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
	rm -f .node-build-image
	rm -f .golang-build-image
