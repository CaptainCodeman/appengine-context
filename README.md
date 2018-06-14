# appengine-context

Http middleware to create an AppEngine context from the request
and make it easy to retrieve it further down the call stack by
saving it into the request context. 

See [this issue](https://github.com/golang/appengine/issues/99) for
why this is necessary - calls to AppEngine services must be made 
using a context created using the _original_ request and if you
have middleware storing things into the context, this replaces
the request with a clone which then interferes with your ability
to call AppEngine services.

The code is tiny but I end up needing it in nearly every AppEngine
project so decided to package it up.

## Installation

Install using `go get`

    go get -i github.com/captaincodeman/appengine-context

## Usage

Add the `Middleware` to your Router package of choice (or wrap
the standard http package Mux) and use the `Context` function
whenever you need to retrieve the AppEngine context to call any
AppEngine services. The middleware should be added _before_ any
other middleware that adds things to the request context.

Example:

```go
package demo

import (
	"fmt"

	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"

	"github.com/captaincodeman/appengine-context"
)

func main() {
	handler := http.HandlerFunc(handle)
	http.Handle("/", gaecontext.Middleware(handler))
	appengine.Main()
}

func handle(w http.ResponseWriter, r *http.Request) {
	ctx := gaecontext.Context(r)

    // simple example just to demonstrate a call to
    // an Appengine Service (stackdriver logging)
    // using the context. This would work even if
    // some other middleware added things into the
    // request context ...
	log.Debugf(ctx, "saying hello")

	fmt.Fprint(w, "hello world!")
}
```