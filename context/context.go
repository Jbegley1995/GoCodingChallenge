//Package context contains all context related functionality.
package context

import (
	"database/sql"
	"fmt"
	"os"
)

//DatabaseConnectionSt Information to connect to the database.
type DatabaseConnectionSt struct {
	User     string
	Host     string
	Password string
	Name     string
}

//Initialize intitializes database information to make it easy to use later.
func (db *DatabaseConnectionSt) Initialize() error {
	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	//Check Environment variables here:

	//--
	db.User = dbUser
	db.Host = dbHost
	db.Password = dbPassword
	db.Name = dbName
	return nil
}

//Context is the global context used to maintain an environment.
type Context struct {
	db DatabaseConnectionSt
}

//Initialize intitializes context information to make it easy to use later.
func (ctxt *Context) Initialize() error {
	var (
		err error
	)

	if err = ctxt.db.Initialize(); err != nil {
		return err
	}
	return nil
}

//GetDB Gets the database through the context and it's environment.
func (ctxt *Context) GetDB() (*sql.DB, error) {
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		ctxt.db.Host,
		ctxt.db.User,
		ctxt.db.Password,
		ctxt.db.Name)
	return sql.Open("postgres", dbinfo)
}