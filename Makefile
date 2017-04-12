all: gulp

node_modules:
	npm install typescript gulp
	npm install

.PHONY: build
build: node_modules
	node_modules/gulp/bin/gulp.js build

.PHONY: gulp
gulp: node_modules
	node_modules/gulp/bin/gulp.js
