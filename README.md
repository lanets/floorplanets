# floorp[LanETS](https://lanets.ca) || floor[(p)LanETS](https://www.nsec.io/2014/04/okiok-win-nsec14-competition/) || floorplan[ETS](https://etsmtl.ca)
[![Build Status](https://travis-ci.org/lanets/floorplanets.svg?branch=master)](https://travis-ci.org/lanets/floorplan-2)

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

The application should be running at : [localhost:3000](http://localhost:3000)
Dev mode runs the application with a watcher, reloading the app on source code change.


To build the application for production, run:
```
$ make build
```


## Running the tests

To run the tests, run:
```
$ make test
```
