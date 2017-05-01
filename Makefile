docker_run_node = docker run --rm -t -i -p 3000:3000 -v $$(pwd):/opt/floorplan -w /opt/floorplan -u $$(id -u):$$(id -g) node

all:
	$(docker_run_node) npm start

node_modules:
	$(docker_run_node) npm install

.PHONY: build
build: node_modules
	$(docker_run_node) npm build

.PHONY: test
test: node_modules
	$(docker_run_node) npm test

.PHONY: clean
clean:
	rm -rf node_modules
	rm -f npm-debug.log
