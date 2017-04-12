# floorplan-2
New floorplan for lanets.ca


## Installing dev environment

Make sure you have [node](https://nodejs.org/en/download/) installed on your machine.

Then install the required packages
```
$ make node_modules
```

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
