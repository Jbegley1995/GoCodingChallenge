package todo

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/jbegley1995/GoCodingChallenge/context"
	"github.com/jbegley1995/GoCodingChallenge/response"

	"github.com/gocraft/web"
)

const (
	//APIKey the api key to access when fetching parameters from the URL.
	APIKey = ":todoID"
)

//Build builds out the routes relating to the todo API.
func Build(router *web.Router) *web.Router {
	return router.Subrouter(context.Context{}, "todos").
		Get("", List).
		Get(APIKey, Get).
		Post("", Create).
		Put(APIKey, Update)
}

// Create will allow a user to create a new todo
// The supported body is {"title": "", "status": ""}
func Create(ctxt *context.Context, w web.ResponseWriter, r *web.Request) {
	db, err := ctxt.GetDB()
	if err != nil {
		//return error
		response.InternalServerError(w, err)
		return
	}
	var todo CreateTodo

	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		response.UnprocessableEntity(w, "Malformed todo sent, please check documentation for the correct structure.")
		return
	}

	if todo.Status == "" || todo.Title == "" {
		response.UnprocessableEntity(w, "Malformed todo sent, please check documentation for the correct structure.")
		return
	}

	invalidStatus := true
	for _, status := range allowedStatuses {
		if todo.Status == status {
			invalidStatus = false
			break
		}
	}

	if !invalidStatus {
		response.UnprocessableEntity(w, "The provided status is not supported")
		return
	}

	newTodo, err := create(db, todo)
	if err != nil {
		//return error
		response.InternalServerError(w, err)
		return
	}
	response.OK(w, newTodo)
}

// Update will allow a user to create a new todo
// The supported body is {"title": "", "status": ""}
func Update(ctxt *context.Context, w web.ResponseWriter, r *web.Request) {
	todoID, err := strconv.Atoi(r.PathParams[APIKey])
	if err != nil {
		//return error
		response.UnprocessableEntity(w, fmt.Sprintf("Need an id to update a Todo"))
		return
	}
	db, err := ctxt.GetDB()
	if err != nil {
		//return error
		response.InternalServerError(w, err)
		return
	}
	var todo CreateTodo

	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		response.UnprocessableEntity(w, "Malformed todo sent, please check documentation for the correct structure.")
		return
	}

	if todo.Status == "" || todo.Title == "" {
		response.UnprocessableEntity(w, "Malformed todo sent, please check documentation for the correct structure.")
		return
	}

	invalidStatus := true
	for _, status := range allowedStatuses {
		if todo.Status == status {
			invalidStatus = false
			break
		}
	}

	if !invalidStatus {
		response.UnprocessableEntity(w, "The provided status is not supported")
		return
	}

	newStatus, err := update(db, todoID, todo)
	if err != nil {
		//return error
		response.InternalServerError(w, err)
		return
	}
	response.OK(w, newStatus)
}

// List will provide a list of all current to-dos
func List(ctxt *context.Context, w web.ResponseWriter, r *web.Request) {
	db, err := ctxt.GetDB()
	if err != nil {
		//return error
		response.InternalServerError(w, err)
		return
	}
	todoList, err := list(db)
	if err != nil {
		//return error
		response.InternalServerError(w, err)
		return
	}

	response.OK(w, Todos{TodoList: todoList})
}

// Added a get in for fun, and since it is easy with the structure used.

// Get will provide a todo based on the id passed.
func Get(ctxt *context.Context, w web.ResponseWriter, r *web.Request) {
	todoID, err := strconv.Atoi(r.PathParams[APIKey])
	if err != nil {
		//return error
		response.UnprocessableEntity(w, fmt.Sprintf("Need an id to fetch a Todo"))
		return
	}
	db, err := ctxt.GetDB()
	if err != nil {
		//return error
		response.InternalServerError(w, err)
		return
	}
	todo, err := get(db, todoID)
	if err != nil {
		//return error
		response.InternalServerError(w, err)
		return
	}

	response.OK(w, Todos{TodoList: []Todo{*todo}})
}
