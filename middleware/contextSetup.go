//Package middleware Handles all middlewares for GoCodingChallenge.
package middleware

import (
	"encoding/json"
	"fmt"

	"github.com/jbegley1995/GoCodingChallenge/context"

	"github.com/gocraft/web"
)

// ContextSetup is a middleware that sets up context variables.
func ContextSetup(ctxt *context.Context, w web.ResponseWriter, r *web.Request, next web.NextMiddlewareFunc) {
	//just wanted to show that I understand the concept of middleware. Here is where I can setup user sessions
	if err := ctxt.Initialize(); err != nil {
		//looks like something wasn't setup properly, handle the error in some way, for this test scenario we're just going
		//to print out an error.
		badError := fmt.Sprintf("Environment not configured properly. \n Error: %s", err)
		//normally I would have a type for errors, but i'm just going to use a small struct.
		jsonResp, _ := json.Marshal(badError)
		fmt.Fprintf(w, string(jsonResp))
		return
	}
	next(w, r)
}
