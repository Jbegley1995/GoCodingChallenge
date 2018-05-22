package routers

import (
	"log"

	"github.com/gocraft/web"
	"github.com/jbegley1995/GoCodingChallenge/context"
	"github.com/jbegley1995/GoCodingChallenge/response"
	"github.com/jbegley1995/GoCodingChallenge/routers/todo"
)

//Build builds out all of the API routers for the application without cluttering up main.
func Build(router *web.Router) *web.Router {
	router.Get("/", Status)

	return todo.Build(router)
}

// Status :=
func Status(ctxt *context.Context, w web.ResponseWriter, r *web.Request) {
	log.Println("Status Request Received")
	response.OK(w, "OK\n")
}
