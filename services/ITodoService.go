package services

import "API/model"

type ITodoService interface {
	Insert(model.Todo) (int64, error)
	Update(int64, model.Todo) (int64, error)
	Delete(int64) (int64, error)
	Get(id int64) (model.Todo, error)
	GetAll() ([]model.Todo, error)
}
