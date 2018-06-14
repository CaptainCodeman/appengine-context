package gaecontext // import "github.com/captaincodeman/appengine-context"

import (
	"context"

	"net/http"

	"google.golang.org/appengine"
)

// contextKey is a value for use with context.WithValue. It's use of
// defining context keys was copied from Go 1.7's new use of context in net/http.
type contextKey struct {
	name string
}

var contextKeyContext = &contextKey{"gae-context"}

// Middleware creates an AppEngine context that can be used to
// call AppEngine services and saves it into the context of the
// request
func Middleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		appengineCtx := appengine.NewContext(r)
		ctx := context.WithValue(r.Context(), contextKeyContext, appengineCtx)
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

// Context returns the previously created AppEngine context so
// calls to AppEngine services can be made even if the request
// has been replaced by other middleware
func Context(r *http.Request) context.Context {
	return r.Context().Value(contextKeyContext).(context.Context)
}
