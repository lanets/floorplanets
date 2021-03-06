# floorp[LanETS](https://lanets.ca) || floor[(p)LanETS](https://www.nsec.io/2014/04/okiok-win-nsec14-competition/) || floorplan[ETS](https://etsmtl.ca)
[![Build Status](https://travis-ci.org/lanets/floorplanets.svg?branch=master)](https://travis-ci.org/lanets/floorplanets)
[![GoDoc](https://godoc.org/github.com/lanets/floorplanets/backend?status.svg)](https://godoc.org/github.com/lanets/floorplanets/backend)

New floorplan for lanets.ca

## Installing dev environment

Make sure you have [docker](https://www.docker.com/), [docker-compose](https://docs.docker.com/compose/) and [Make](https://www.gnu.org/software/make/) installed on your machine.

## Building

To build the Floorplan API and client, run:
```
$ make
```
## Starting the Floorplan API
To start the Floorplan API, run:
```
$ make run-floorplanets
```
The API should start listening for requests at [localhost:8080](http://localhost:8080).

## Starting the Floorplan client
To start the Floorplan client in dev mode, run:
```
$ make reactapp
```
The application should be running at [localhost:3000](http://localhost:3000).
Dev mode runs the application with a watcher, reloading the app on source code change.


## Running the tests

To run the tests, run:
```
$ make test
```
