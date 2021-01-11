package models

import "github.com/lucabecci/TODO-golang-psql/src/database"

type Todo struct {
	ID          int    `json:id`
	Description string `json:string`
}

//Insert is a function for save the todo in the db
func Insert(description string) (Todo, bool) {
	db := database.GetConnection()
	var todoId int
	db.QueryRow("INSERT INTO todos(description) VALUES($1) RETURNING id", description).Scan(&todoId)
	if todoId == 0 {
		return Todo{}, false
	}
	return Todo{todoId, ""}, true
}

//Get is a function for get a todo of the db
func Get(id string) (Todo, bool) {
	db := database.GetConnection()

	row := db.QueryRow("SELECT * FROM todos WHERE id = $1", id)

	var ID int
	var description string

	err := row.Scan(&ID, &description)

	if err != nil {
		return Todo{}, false
	}

	return Todo{ID, description}, true
}
