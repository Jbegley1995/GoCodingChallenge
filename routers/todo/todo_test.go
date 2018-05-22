package todo

import (
	"testing"

	"github.com/jbegley1995/GoCodingChallenge/context"
)

var (
	ctxt = context.Testing()
)

func TestList(t *testing.T) {
	if ctxt == nil {
		t.Fatal("Context not set up properly.")
	}

	db, err := ctxt.GetDB()
	if err != nil {
		t.Fatalf("Database Connection error: %s", err.Error())
	}

	_, err = list(db)
	if err != nil {
		t.Errorf("Could not list todos: %s", err.Error())
	}
}

//TestAPI A sequential test of api success.
func TestAPI(t *testing.T) {
	var (
		err          error
		todoToCreate CreateTodo
		todoToUpdate CreateTodo
		updatedTodo  *Todo
		createdTodo  *Todo
	)
	if ctxt == nil {
		t.Fatal("Context not set up properly.")
	}

	db, err := ctxt.GetDB()
	if err != nil {
		t.Fatalf("Database Connection error: %s", err.Error())
	}
	todoToCreate = CreateTodo{
		Title:  "Fizz",
		Status: "New",
	}
	createdTodo, err = create(db, todoToCreate)
	if err != nil {
		t.Errorf("Could not list todos: %s", err.Error())
	}

	todoToUpdate = CreateTodo{
		Title:  "Fizz",
		Status: "In Progress",
	}

	updatedTodo, err = update(db, createdTodo.ID, todoToUpdate)
	if err != nil {
		t.Errorf("Could not update todo: %s", err.Error())
	}

	if updatedTodo.Status != todoToUpdate.Status {
		t.Errorf("Read back the wrong status, wanted %s got %s", todoToUpdate.Status, updatedTodo.Status)
	}
	_, err = get(db, createdTodo.ID)
	if err != nil {
		t.Errorf("Could not get todo: %s", err.Error())
	}

	//Theres a bunch more tests we could do, but i'm going to cut it off here for time purposes.
}

//Run test on router relay down here, this can be complicated due to the request and response writer simulation
//so i'm just going to do the db tests, in the real world I would have something setup for this.
