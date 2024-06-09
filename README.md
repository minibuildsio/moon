# Moon

A CLI app that prints the current lunar phase.

## Build

```bash
go build
```

## Run

```bash
./moon
# prints "Current lunar phase: Waning Gibbous 🌖"
```

## Build and Run using Docker 

Build the docker image using docker builder to both build the Go application and package it as a docker image:

```
docker build -t moon .
```

Run the image:

```
docker run moon
```
