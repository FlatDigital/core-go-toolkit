<div id="top"></div>

<!-- PROJECT SHIELDS -->
<!--
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/FlatDigital/core-go-toolkit">
    <img src="https://i.ibb.co/H49WMvQ/6106732-overlay-overlay-1-removebg-preview.png" alt="6106732-overlay-overlay-1-removebg-preview" alt="Logo" width="340" height="150">
  </a>

<h3 align="center">Project: Core-Go-Toolkit</h3>

  <p align="center">
    Team: Tech-Core
    <br />
    <a href="https://github.com/FlatDigital/core-go-toolkit/issues">Report Bug</a>
    ·
    <a href="https://github.com/FlatDigital/core-go-toolkit/issues">Request Feature</a>
  </p>
</div>

<!-- ABOUT THE PROJECT -->
## About The Project

Core Go Toolkit is a set of libraries aimed at writing Go code within the Flat ecosystem easier

<p align="right">(<a href="#top">back to top</a>)</p>

### Built With

* [Go](https://go.dev/) Language

### Features

* Custom Flat Context.
* Logger library for golang.
* Functionality needed for initializing a web server with sane settings based on a given Flat scope.
* Database connector (PostgreSQL) && Database Operations
* Error handling library
* Datadog Custom Metrics (WIP)
* Simple HTTP and REST client library for Go ([Resty](https://github.com/go-resty/resty))
* Secrets management service

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- USAGE EXAMPLES -->
# Usage

### Custom Flat Context

In order to use `flat.Context` package, you should define your HTTP handlers as follow:

```go
func ControllerHandler(c *gin.Context, ctx *flat.Context) {
    // ...
}
```

When hooking tour handler to GIN you must encapsulate it in `flat.Handler` function.

```go
v1 := g.Group("/v1")
v1.POST("/", flat.Handler(ControllerHandler))
```

The flat context gives you typed access to request data, and a configured logger that will automatically add the `request_id` to any logged lines.

### Logger library for golang

Import and use it.

```go
import (
 "github.com/gin-gonic/gin"
 "net/http"
 "github.com/FlatDigital/core-go-toolkit/core/libs/go/logger"
)

func Ping(c *gin.Context) {
 logger := logger.LoggerWithName(c, "Ping")

 logger.Info("PingLog", logger.Attrs{
  "atts_1": "value 1",
  "...":    "value ...",
  "atts_n": "value n",
 })

 c.String(http.StatusOK, "pong")
}
```

Methods

```go
LoggerWithName(c *gin.Context, name string) *Logger 
```

Returns a logger pointer with gin context information. The name is used for more information.

```go
(l *Logger) LogWithLevel(level string, event string, attrs ...Attrs) *Logger 
```

Basic log method. Level `INFO`, `DEBUG`, `WARN` or `ERROR`, and an event are required.
Optional attrs of type `Attrs - map[string]interface{}` can be passed.

Direct methods for each level are provided.

```go
(l *Logger) Debug(event string, attrs ...Attrs) *Logger

(l *Logger) Error(event string, attrs ...Attrs) *Logger

(l *Logger) Warning(event string, attrs ...Attrs) *Logger

(l *Logger) Info(event string, attrs ...Attrs) *Logger
```

### Functionality needed for initializing a web server

This package is though for using when bootstraping the `main` package of a new application.

Toolkit uses scopes for defining which functionality of an application a given scaling group should respond to.

For example, given a scope `production-read` the application should bootstrap itself so that it only accept requests pertaining `read` functionality and respond `404` for every other endpoint; using server defaults for a `production` environment.

The scope format used is: `environment`-`role`-`tag` (the `tag` segment is optional).

```go
package main

import (
 "github.com/FlatDigital/core-go-toolkit/core/libs/go/server"
)

routes := server.RoutingGroup{
    server.RoleRead: func (g *gin.RouterGroup) {
    g.GET("/reader", func (c *gin.Context) {})
    },
    
    server.RoleWrite: func (g *gin.RouterGroup) {
    g.POST("/writer", func (c *gin.Context) {})
    },
}

srv, err := server.NewEngine("develop-read", routes)
if err != nil {
    log.Fatal(err)
}

srv.Run(":8080")
```

Alternatively you can use a controller with a method that returns a function with signature `func(*gin.RouterGroup)` and then use that to register the appropriate routes.

In effect you'll be delegating the knowledge of the routes to the module, instead of bundling it with your main program.

```go
reader := reader.NewReader(container)

routes := server.RoutingGroup{
    server.RoleRead: reader.RegisterRoutes(basePath),
    // ...
}
```

### Support for Database connector

Import lib and use it.

```go
import (
    "github.com/FlatDigital/core-go-toolkit/database"
)

dbConfig := database.ServiceConfig{
DBUsername: config.DBReadUsername(),
DBPassword: config.DBReadPassword(),
DBHost:     config.DBHost(),
DBName:     config.DBName(),

MaxIdleConns:     config.DBPoolMaxIdleConns(),
MaxOpenConns:     config.DBPoolMaxOpenConns(),
ConnMaxLifetime:  time.Duration(config.DBConnMaxLifetime()),
ConnReadTimeout:  config.DBConnectionReadTimeout(),
ConnWriteTimeout: config.DBConnectionWriteTimeout(),
ConnTimeout:      config.DBConnectionTimeout(),

// MaxConnectionRetries default = 3
// MaxConnectionRetries: ,
}

// Connect to database
db, err := database.NewService(dbConfig)
if err != nil {
panic(err.Error())
}
```

ConnReadTimeout and ConnWriteTimeout not supported by postgres connetion string. If you are interested in being able to set these values, look at this workarround -> <https://github.com/Kount/pq-timeouts>

`config` contains all settings for a given environment.

### Support for Database Operations

List of basic operations

```go
  Execute(dbc *DBContext, query string, params ...interface{}) (*DBResult, error)
  ExecuteEnsuringOneAffectedRow(dbc *DBContext, query string, params ...interface{}) error
  Select(dbc *DBContext, query string, forUpdate bool, params ...interface{}) (*DBResult, error)
  SelectUniqueValue(dbc *DBContext, query string, forUpdate bool, params ...interface{}) (*DBRow, error)
  SelectUniqueValueNonEmpty(dbc *DBContext, query string, forUpdate bool, params ...interface{}) (*DBRow, error)
  
  Begin(dbc *DBContext) (*DBContext, error)
  Commit(dbc *DBContext) error
  Rollback(dbc *DBContext) error
  Connection() (*DBContext, error)
  Close(dbc *DBContext) error
```

For more information check the Database lib inside the toolkit

Simple example of a SELECT statement:

```go
const (
 getPropertyByID = `
  SELECT
   *
  FROM
   rs_webproperty
  WHERE
   id = $1;
 `
)

func (repository *propertyListingDatabase) GetPropertyListing(id int64) (*entities.PropertyListing, errorWrapper.Wrapper) {
 params := []interface{}{
  id,
 }

 dbRow, err := repository.database.SelectUniqueValue(nil, getPropertyByID, false, params...)
 if err != nil {
  return &entities.PropertyListing{}, errorWrapper.Wrap(err)
 }

 // No record found
 if dbRow == nil {
  return nil, nil
 }

 propertyListing, errConvert := convertToPropertyListing(dbRow)
 if errConvert != nil {
  return nil, errConvert
 }

 // done
 return propertyListing, nil
}
```

### Error handling library

This lib has everything you need to handle errors in our application.

Import lib and use it.

```go
import (
  errorWrapper "github.com/FlatDigital/core-go-toolkit/error"
)
```

Most important methods:

```go
// New returns a new err container from the given error
func New(format string, params ...interface{}) Wrapper {
  ...
}

// Wrap wraps an error in an error container
func Wrap(err error) Wrapper {
  ...
}

func ReturnError(c *gin.Context, errWrapped Wrapper) {
  ...
}

func ReturnWrappedErrorFromStatus(statusCode int, err error) Wrapper {
  ...
}
```

Most common uses:

Example 1:

```go
// check if id is empty
if len(propertyListingIDStr) == 0 {
  return 0, errorWrapper.New("property listing id can not be empty")
}
```

Example 2:

```go
dbRow, err := repository.database.SelectUniqueValue(nil, getPropertyByID, false, params...)
if err != nil {
  return &entities.PropertyListing{}, errorWrapper.Wrap(err)
}
```

Example 3:

```go
propertyListing, err := handler.usecase.GetPropertyListing(id)
if err != nil {
  errorWrapper.ReturnError(c, err)
  return
}
```

Example 4:

```go
status, responseRaw, err := repository.resty.MakeGetRequest(nil, url, http.Header{})
if err != nil {
  return weatherMsg, errorWrapper.ReturnWrappedErrorFromStatus(status, fmt.Errorf("error executing GET %s [status:%v][response:%s]", url, status, responseRaw))
}
```

### Datadog Custom Metrics (WIP)

Import lib and use it.

```go
import (
  "github.com/FlatDigital/core-go-toolkit/src/api/libs/godog"
)
```

Example:

```go
tags := new(godog.Tags).
 Add("core_property_listings_api_env", string(application.Context().Environment())).
 Add("property_listing_id", fmt.Sprint(id))

godog.RecordSimpleMetric("application.core_property_listings_api.get_property_listing_by_id", 1, tags.ToArray()...)
```

### Simple HTTP and REST client library for Go (Resty)

Import lib and use it.

```go
import (
    "github.com/FlatDigital/core-go-toolkit/src/api/libs/rest"
)
```

```go
// Default RestyClient
restyClient := rest.NewRestyService()

// If you need you can create a restyclient with config
// Configure according to your needs
requestConfig := rest.RequestConfig{
    DisableTimeout: false,
    Timeout:        3 * time.Second,
    ConnectTimeout: 1500 * time.Millisecond,
}
serviceConfig := rest.ServiceConfig{
    BaseURL:             config.FlatURL(),
    MaxIdleConnsPerHost: config.MaxIdleConnsPerHost(),
    RequestConfig:       &requestConfig,
}

restyClient := rest.NewRestyServiceWithConfig(serviceConfig)
```

Most common uses:

```go
// GET
status, responseRaw, err := repository.resty.MakeGetRequest(nil, url, http.Header{})
if err != nil {
    //do some stuff
}

// POST
status, responseRaw, err := repository.resty.MakePostRequest(nil, url, body, http.Header{})
if err != nil {
    //do some stuff
}
...
```

Check all the operations available inside the rest library

### Secrets service

This service retrieves secrets that the application needs. Currently the secrets are obtained from environment variables -> `os.Getenv(key)`

Import lib and use it.

```go
import (
  "github.com/FlatDigital/core-go-toolkit/src/api/libs/secrets"
)

// Example
secrets := secrets.NewService()
scope, err := secrets.Get("SCOPE")
if scope == "" || err != nil {
  panic(fmt.Errorf("application initialization error - No scope defined"))
}
```

## Execute test with coverage

``` sh
  go test -cover ./... 
```

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- CONTACT -->
## Contact

Slack: `#tech-core`

Nico Albani: nico@flat.mx

Nico de Lara: nicolas.delara@flat.mx

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://badges.flat.mx/github/contributors/FlatDigital/core-go-toolkit.svg?style=for-the-badge
[contributors-url]: https://github.com/FlatDigital/core-go-toolkit/graphs/contributors
[forks-shield]: https://badges.flat.mx/github/forks/FlatDigital/core-go-toolkit.svg?style=for-the-badge
[forks-url]: https://github.com/FlatDigital/core-go-toolkit/network/members
[stars-shield]: https://badges.flat.mx/github/stars/FlatDigital/core-go-toolkit.svg?style=for-the-badge
[stars-url]: https://github.com/FlatDigital/core-go-toolkit/stargazers
[issues-shield]: https://badges.flat.mx/github/issues/FlatDigital/core-go-toolkit.svg?style=for-the-badge
[issues-url]: https://github.com/FlatDigital/core-go-toolkit/issues
