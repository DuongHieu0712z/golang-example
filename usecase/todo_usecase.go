package usecase

import (
	"context"
	"example/common/pagination"
	"example/db"
	"example/dto"
	"example/form"
	"example/model"
	"example/repository"
	"example/uow"
)

type TodoUsecase interface {
	GetPagedList(ctx context.Context, param pagination.PagingParam) (*pagination.PagedList, error)
	GetById(ctx context.Context, id string) (*dto.TodoDto, error)
	Create(ctx context.Context, form form.TodoForm) (*dto.TodoDto, error)
	Update(ctx context.Context, id string, form form.TodoForm) error
	Delete(ctx context.Context, id string) error
}

type todoUsecase struct {
	uow      uow.UnitOfWork
	todoRepo repository.TodoRepository
}

func NewTodoUsecase(db *db.Database) TodoUsecase {
	usecase := &todoUsecase{
		uow: uow.NewUnitOfWork(db),
	}
	usecase.todoRepo = usecase.uow.Todos()
	return usecase
}

func (uc *todoUsecase) GetPagedList(ctx context.Context, param pagination.PagingParam) (*pagination.PagedList, error) {
	data, err := uc.todoRepo.GetPagedList(ctx, param)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (uc *todoUsecase) GetById(ctx context.Context, id string) (*dto.TodoDto, error) {
	_, err := uc.todoRepo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	var obj dto.TodoDto
	// automapper.MapLoose(data, &obj)
	return &obj, nil
}

func (uc *todoUsecase) Create(ctx context.Context, form form.TodoForm) (*dto.TodoDto, error) {
	var data model.Todo
	// automapper.MapLoose(form, &data)

	if err := uc.todoRepo.Create(ctx, &data); err != nil {
		return nil, err
	}

	var obj dto.TodoDto
	// automapper.MapLoose(data, &obj)
	return &obj, nil
}

func (uc *todoUsecase) Update(ctx context.Context, id string, form form.TodoForm) error {
	data, err := uc.todoRepo.GetById(ctx, id)
	if err != nil {
		return err
	}

	// automapper.MapLoose(form, &data)

	if err := uc.todoRepo.Update(ctx, data); err != nil {
		return err
	}

	return nil
}

func (uc *todoUsecase) Delete(ctx context.Context, id string) error {
	if err := uc.todoRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
