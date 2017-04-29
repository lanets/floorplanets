all: gulp

docker_run_node = docker run --rm -t -v $$(pwd):/opt/floorplan -w /opt/floorplan -u $$(id -u):$$(id -g) node

node_modules:
	$(docker_run_node) npm install typescript gulp
	$(docker_run_node) npm install

.PHONY: build
build: node_modules
	$(docker_run_node) node node_modules/gulp/bin/gulp.js build

.PHONY: gulp
gulp: node_modules
	$(docker_run_node) node node_modules/gulp/bin/gulp.js

.PHONY: test
test: node_modules
	$(docker_run_node) npm test

.PHONY: clean
clean:
	rm -rf node_modules
	rm -f npm-debug.log
