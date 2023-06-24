package repositories

import (
	"API/database"
	"API/model"
	"database/sql"
	"log"
)

var conn *sql.DB

/* func Insert(todo model.Todo) (id int64, err error) {
	fmt.Println(todo)
	database.CreateConnection(conn)
	fmt.Println("rodou aqui")
	defer conn.Close()
	fmt.Println("rodou aqui2")

	query := `INSERT INTO todos (title, description, is_done) VALUES ($1, $2, $3) RETURNING id`
	fmt.Println("rodou aqui3")
	err = conn.QueryRow(query, todo.Title, todo.Description, todo.IsDone).Scan(&id)

	fmt.Println("rodou aqui4")

	return
} */

func Insert(todo model.Todo) (id int64, err error) {
	conn, err := database.CreateConnection()
	if err != nil {
		log.Fatal("Cannot establish connection")
	}

	defer conn.Close()

	query := `INSERT INTO todos (title, description, is_done) VALUES ($1, $2, $3) RETURNING id`
	err = conn.QueryRow(query, todo.Title, todo.Description, todo.IsDone).Scan(&id)

	return
}

func Get(id int64) (todo model.Todo, err error) {
	conn, err := database.CreateConnection()
	if err != nil {
		log.Fatal("Cannot establish connection")
	}

	defer conn.Close()

	row := conn.QueryRow(`SELECT id, title, description, is_done FROM todos WHERE id = $1`, id)
	err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.IsDone)

	return
}

func GetAll() (todos []model.Todo, err error) {
	conn, err := database.CreateConnection()
	if err != nil {
		log.Fatal("Cannot establish connection")
	}

	defer conn.Close()

	rows, err := conn.Query(`SELECT id, title, description, is_done FROM todos`)
	if err != nil {
		log.Fatal("Can't get all todos")
		return
	}

	for rows.Next() {
		var todo model.Todo

		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.IsDone)
		if err != nil {
			continue
		}

		todos = append(todos, todo)
	}

	return
}

func Update(id int64, todo model.Todo) (int64, error) {
	conn, err := database.CreateConnection()
	if err != nil {
		log.Fatal("Cannot establish connection")
	}

	defer conn.Close()

	response, err := conn.Exec(`UPDATE todos set title=$1, description=$2, is_done=$3 WHERE id=$4`, todo.Title, todo.Description, todo.IsDone, id)

	if err != nil {
		return 0, err
	}

	return response.RowsAffected()
}

func Delete(id int64) (int64, error) {
	conn, err := database.CreateConnection()
	if err != nil {
		log.Fatal("Cannot establish connection")
	}

	defer conn.Close()

	response, err := conn.Exec(`DELETE FROM todos WHERE id=$1`, id)

	if err != nil {
		return 0, err
	}

	return response.RowsAffected()
}
