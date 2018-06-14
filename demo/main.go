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
