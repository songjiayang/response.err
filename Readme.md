# response.err

A error helper with gogo api response.

### Installation

Via go-get

```bash
go get github.com/songjiayang/response.err
```

### Usage

```golang
package main

import (
	"github.com/dolab/gogo"
	"github.com/songjiayang/response.err"
)

func main() {
	app := gogo.New("development", "")

	app.GET("/err", func(ctx *gogo.Context) {
		ctx.Json(errors.NewErrorResponse(ctx.RequestID(), ctx.RequestURI(), errors.InvalidParameter))
	})

	app.Run()
}

```

When start app and visit `/err` page, you will see responsed message like:

```json
{
code: "InvalidParameter",
message: "A parameter specified in a request is not valid, is unsupported, or cannot be used.",
resource: "/err",
request_id: "5a178eb13d151e8040000002",
error: ""
}
```


### Custom Error

This package just define response error schema but don't contain many errors.

Don't worry, you can custom your error easily.

```
# custom_error.go

import (
	"github.com/dolab/gogo"
	"github.com/songjiayang/response.err"
)

var (
	CustomErr2     = errros.Error{400, "xxx", "xxx"}
  CustomErr1     = errros.Error{403, "xxx", "xxxx"}
)

```
