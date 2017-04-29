# floorplan-2
[![Build Status](https://travis-ci.org/lanets/floorplan-2.svg?branch=master)](https://travis-ci.org/lanets/floorplan-2)

New floorplan for lanets.ca

## Installing dev environment

Make sure you have [docker](https://www.docker.com/) and [Make](https://www.gnu.org/software/make/) installed on your machine.

## Building

To compile and build the TypeScript code, run this gulp task:
```
$ make build
```

While developping, you can ask gulp to recompile the project
every time you change a file. Just start gulp in the background using :
```
$ make gulp
```


## Running the tests

We use [Jasmine](https://jasmine.github.io/) for our unit tests library.
To run the tests, run:
```
$ make test
```
