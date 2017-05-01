# floorplan-2
[![Build Status](https://travis-ci.org/lanets/floorplan-2.svg?branch=master)](https://travis-ci.org/lanets/floorplan-2)

New floorplan for lanets.ca

## Installing dev environment

Make sure you have [docker](https://www.docker.com/) and [Make](https://www.gnu.org/software/make/) installed on your machine.

To install node dependencies, run :
```
$ make node_modules
```

## Building

To start the application in dev mode, run:
```
$ make
```
Dev mode runs the application with a watcher, reloading the app on source code change.


## Running the tests

To run the tests, run:
```
$ make test
```
