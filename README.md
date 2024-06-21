# JSH

A POSIX compliant shell that's capable of interpreting shell commands, running external 
programs and builtin commands like cd, pwd, echo etc.

## Requirements

`go` v1.22

## Makefile

To simplify the build and run process, you can use the following Makefile commands:

`make all` - test and build the application

`make build` - build the application and output the executable into /dist

`make run` - build and run the application

`make fmt` - format the codebase with go fmt

`make test` - run the unit tests

## Docker

To build and run the JSH shell using Docker, you can use the following commands:

`make docker-build` - build the docker image

`make docker-run` - run the image exposed on port 8080

## Usage

`cd` - change directory (supports absolute and relative paths)

`pwd` - print current directory

`type <command>` - is it a native / supported command e.g. `type pwd` returns type pwd: is a shell built in command

`exit` - exit code 0

`echo hello world` - echo a string of text