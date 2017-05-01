docker_run_node = docker run --rm -t -i -p 3000:3000 -v $$(pwd):/opt/floorplan -w /opt/floorplan -u $$(id -u):$$(id -g) node
docker_run_node_CI = docker run --rm -t -i -p 3000:3000 -v $$(pwd):/opt/floorplan -w /opt/floorplan -u $$(id -u):$$(id -g) -e CI=true node

all:
	$(docker_run_node) npm start

node_modules:
	$(docker_run_node) npm install

.PHONY: build
build: node_modules
	$(docker_run_node) npm run build

.PHONY: test
test: node_modules
	$(docker_run_node) npm test

.PHONY: test-CI
test-CI: node_modules
	$(docker_run_node_CI) npm test

.PHONY: clean
clean:
	rm -rf node_modules
	rm -f npm-debug.log
