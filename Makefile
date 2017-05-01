docker_run_cmd = docker run --rm -t -i -v $$(pwd):/opt/floorplan -w /opt/floorplan -u $$(id -u):$$(id -g)

all:
	$(docker_run_cmd) -p 3000:3000 node npm start

node_modules:
	$(docker_run_cmd) node npm install

.PHONY: build
build: node_modules
	$(docker_run_cmd) node npm run build

.PHONY: test
test: node_modules
	$(docker_run_cmd) node npm test

.PHONY: test-CI
test-CI: node_modules
	$(docker_run_cmd) -e CI=true node npm test && npm run flow

.PHONY: clean
clean:
	rm -rf node_modules
	rm -f npm-debug.log
