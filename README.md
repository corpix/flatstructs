flatstructs
-----------

[![Build Status](https://travis-ci.org/corpix/flatstructs.svg?branch=master)](https://travis-ci.org/corpix/flatstructs)

Provides helpers to flatten nested structures in go.

## Example

``` go
package main

import (
	"fmt"

	"github.com/corpix/flatstructs"
)

type Scope struct {
	Request
	Connection
}

type Request struct {
	Headers
}

type Headers struct {
	UserAgent string
	Referer   string
}

type Connection struct {
	Host string
	Port int
}

type Event struct {
	ID     int
	Source string
	Scope
}

func main() {
	event := &Event{
		ID:     1,
		Source: "system",
		Scope: Scope{
			Request: Request{
				Headers: Headers{
					UserAgent: "curl",
					Referer:   "http://example.com",
				},
			},
			Connection: Connection{
				Host: "127.0.0.1",
				Port: 1337,
			},
		},
	}

	keys, err := flatstructs.Keys(event)
	if err != nil {
		panic(err)
	}

	values, err := flatstructs.Values(event)
	if err != nil {
		panic(err)
	}

	fmt.Println("flatstructs.Keys(event) -> ", keys)
	fmt.Println("flatstructs.Values(event) -> ", values)
}
```

Now save this code into `snippet.go` and:

``` shell
go run ./snippet.go
flatstructs.Keys(event) ->  [ID Source ScopeRequestHeadersUserAgent ScopeRequestHeadersReferer ScopeConnectionHost ScopeConnectionPort]
flatstructs.Values(event) ->  [1 system curl http://example.com 127.0.0.1 1337]
```

## Limitations

> Some of them are not limitations actually but it is worth to mention them here.

* If getting `Values()` and somewhere in the chain there is a slice of `struct`'s it will be returned untouched
* Map's is also returned untouched when getting `Values()` and `Keys()` is not going inside map's

## Customization

### Builder

There is things you could customize:

* Key delimiter
* Key name
* Key tag name

This customization's are available via custom `Builder`:

``` go
builder := flatstructs.NewBuilder("key", ".")

builder.Keys(...)
builder.Values(...)
builder.Map(...)
```
