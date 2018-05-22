package main

import (
	"log"
	"net/http"

	"github.com/jbegley1995/GoCodingChallenge/context"
	"github.com/jbegley1995/GoCodingChallenge/routers"

	"github.com/gocraft/web"
	"github.com/jbegley1995/GoCodingChallenge/middleware"
	_ "github.com/lib/pq"
)

func main() {
	router := web.New(context.Context{}).
		Middleware(middleware.ContextSetup)

	builtRouter := routers.Build(router)

	log.Println("Starting server...")

	// Make sure you have DB_USER, DB_PASSWORD and DB_NAME environment variables set.
	// We use them elsewhere
	log.Fatal(http.ListenAndServe(":8080", builtRouter))
}
