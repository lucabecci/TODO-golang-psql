package models

import (
	"log"

	"github.com/lucabecci/TODO-golang-psql/src/database"
)

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
	return Todo{todoId, description}, true
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

//GetAll is a function for get all todos of the db
func GetAll() []Todo {
	db := database.GetConnection()

	rows, err := db.Query("SELECT * FROM todos ORDER BY id")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var todos []Todo

	for rows.Next() {
		t := Todo{}

		var ID int
		var description string

		err := rows.Scan(&ID, &description)

		if err != nil {
			log.Fatal(err)
		}

		t.ID = ID
		t.Description = description

		todos = append(todos, t)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return todos
}

//Delete is for delete a todo of the db
func Delete(id string) (Todo, bool) {
	db := database.GetConnection()

	var todoId int

	db.QueryRow("DELETE FROM todos WHERE id = $1 RETURNING id", id).Scan(&todoId)

	if todoId == 0 {
		return Todo{}, false
	}

	return Todo{todoId, ""}, true
}

//Update is for update a todo of the db
func Update(id string, description string) (Todo, bool) {
	db := database.GetConnection()

	var todoId int
	db.QueryRow("UPDATE todos SET description = $1 WHERE id = $2 RETURNING id", description, id).Scan(&todoId)

	if todoId == 0 {
		return Todo{}, false
	}

	return Todo{todoId, description}, true
}
