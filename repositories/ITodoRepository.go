package repositories

import "API/model"

type ITodoRepository interface {
	Insert(model.Todo) (int64, error)
	Get(int64) (model.Todo, error)
	GetAll() ([]model.Todo, error)
	Update(int64, model.Todo) (int64, error)
	Delete(int64) (int64, error)
}
