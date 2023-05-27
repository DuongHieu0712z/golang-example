package uow

import (
	"example/data/repository"
	"example/db"
)

type UnitOfWork interface {
	Todos() repository.TodoRepository
	Users() repository.UserRepository
}

type unitOfWork struct {
	db *db.Database
}

func NewUnitOfWork(db *db.Database) UnitOfWork {
	return &unitOfWork{db: db}
}

func (uow *unitOfWork) Todos() repository.TodoRepository {
	return repository.NewTodoRepository(uow.db)
}

func (uow *unitOfWork) Users() repository.UserRepository {
	return repository.NewUserRepository(uow.db)
}
