# Clean Architecture in Go

This is my version of how to structure a GO application. 

to demonstrate my idea I'm going to use the weather API as an example


## Frameworks

The example uses the following frameworks:

* [Gin-Gonic](https://github.com/gin-gonic/gin) to make the API layer
* [Resty](https://github.com/go-resty/resty) to make API calls

## Run

Clone the repository and execute the following command on the root folder

``` bash
go run main.go
```

## Test

``` bash
go test ./...
```

## Backlog

> This is a work in progress. PRs are welcome :)
* Add (a LOT) of unit tests
* Use `Context` type to pass parameters across different layers
* Add an example of database handler module

