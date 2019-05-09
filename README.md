# gofast http server

Simple http server written in Go. Also it comes with a few tests to check the endpoint results.

> Gotta go fast!

## Folder structure

```
gofast
    ├── LICENSE.md
    ├── README.md
    ├── api
    │   └── endpoints
    │       ├── endpoints.go
    │       └── endpoints_test.go
    └── main
        └── main.go
```
## Requirements

This project has been coded under **go1.12.5** version and **$GOPATH** must be defined.

## Checkout 

Clone the repository:

```
$ cd $GOPATH/src
$ git checkout url
```

## Build

Once source code is downloaded go to gofast root folder and build the main.go file to make the executable and place 
it in the main sub folder:

```
$ cd $GOPATH/src/gofast
$ go build -o main/main main/main.go
```

## Test

To run tests just execute:

```
$ go test -json ./...
$ ok  	gofast/api/endpoints
$ ?   	gofast/main	[no test files]
```

And, obtaining the coverage:

```
$ go test -coverprofile=coverage.out ./...
$ ok  	gofast/api/endpoints	0.040s	coverage: 94.4% of statements
$ ?   	gofast/main	[no test files]
```

There is no tests for main source code file yet 😅

## Install

If you want to install the http server main file executable:

```
$ cd $GOPATH/src/gofast/main
$ go install
```

## Run server

Just run the main executable:

```
$ ./main
```

Now, with the http server up and running, we can make requests:

### Endpoints

* 	http://localhost:3003/health-check
* 	http://localhost:3003/ping
* 	http://localhost:3003/teapot
*   http://localhost:3003/api/gofast

## Author

Esteban Martín Busto

## License 

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
