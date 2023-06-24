package services

import (
	"API/model"
	"API/repositories"
)

func Insert(todo model.Todo) (id int64, err error) {
	return repositories.Insert(todo)
}

func Update(id int64, todo model.Todo) (int64, error) {
	return repositories.Update(id, todo)
}

func Delete(id int64) (int64, error) {
	return repositories.Delete(id)
}

func Get(id int64) (model.Todo, error) {
	return repositories.Get(id)
}

func GetAll() ([]model.Todo, error) {
	return repositories.GetAll()
}
