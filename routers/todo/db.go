//Package todo A simple API package that manages todos.
package todo

import (
	"database/sql"
	"fmt"
)

//selectQry is the query to read todo's from the database.
var selectQry = "SELECT id, title, status FROM todo "

//read Reads the todos from the database.
func read(db *sql.DB, todoID int) ([]Todo, error) {
	var (
		todoList       = []Todo{}
		dbParams       = []interface{}{}
		whereStatement string
	)

	if todoID != -1 {
		whereStatement += "WHERE id = ? "
		dbParams = append(dbParams, todoID)
	}
	rows, err := db.Query(selectQry+whereStatement, dbParams...)
	if err != nil {
		//normally would need to do something with this error, logging most likely.
		fmt.Println(err)
		return nil, fmt.Errorf("Failed to build todo list")
	}
	defer rows.Close()

	for rows.Next() {
		todo := Todo{}
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Status); err != nil {
			//normally would need to do something with this error, logging most likely.
			fmt.Println(err)
			return nil, fmt.Errorf("Failed to build todo list")
		}

		todoList = append(todoList, todo)
	}
	return todoList, nil
}

//list lists out all of the todos in the database.
func list(db *sql.DB) ([]Todo, error) {
	return read(db, -1)
}

//get gets a todo by its id from the database.
func get(db *sql.DB, todoID int) (*Todo, error) {
	todos, err := read(db, -1)
	if err != nil {
		return nil, err
	}
	//should be caught by a sql exception but just in case.
	if len(todos) == 0 {
		return nil, fmt.Errorf("there is no todo for the ID sent")
	}
	return &todos[0], nil
}

//create creates a todo in the database
func create(db *sql.DB, todo CreateTodo) (*Todo, error) {
	var (
		todoID int
	)

	insertStmt := fmt.Sprintf(`INSERT INTO todo (title, status) VALUES ('%s', '%s') RETURNING id`, todo.Title, todo.Status)

	// Insert and get back newly created todo ID
	if err := db.QueryRow(insertStmt).Scan(&todoID); err != nil {
		return nil, fmt.Errorf("Failed to save to db: %s", err.Error())
	}

	fmt.Printf("Todo Created -- ID: %d\n", todoID)

	return get(db, todoID)
}

//update updates the todo based with the same id, and returns the todo's state back to the user.
func update(db *sql.DB, todoID int, todo CreateTodo) (state string, _ error) {
	updateStmt := fmt.Sprintf(`UPDATE todo SET title = ? AND status = ? WHERE id = ?`)

	// Insert and get back newly created todo ID
	if _, err := db.Exec(updateStmt); err != nil {
		return "", fmt.Errorf("Failed to save to db: %s", err.Error())
	}

	createdTodo, err := get(db, todoID)
	if err != nil {
		return "", err
	}
	return createdTodo.Status, nil
}
