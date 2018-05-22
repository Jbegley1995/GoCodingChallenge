//Package middleware Handles all middlewares for GoCodingChallenge.
package middleware

import (
	"fmt"

	"github.com/jbegley1995/GoCodingChallenge/context"
	"github.com/jbegley1995/GoCodingChallenge/response"

	"github.com/gocraft/web"
)

// ContextSetup is a middleware that sets up context variables.
func ContextSetup(ctxt *context.Context, w web.ResponseWriter, r *web.Request, next web.NextMiddlewareFunc) {
	//just wanted to show that I understand the concept of middleware. Here is where I can setup user sessions
	if err := ctxt.Initialize(); err != nil {
		fmt.Println(err)
		//looks like something wasn't setup properly, handle the error in some way, for this test scenario we're just going
		//to print out an error.
		response.InternalServerError(w, fmt.Errorf("environment not configured properly"))
		return
	}
	next(w, r)
}
