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

	"github.com/devfeel/mapper"
	"github.com/go-playground/validator/v10"
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
	validate *validator.Validate
}

func NewTodoUsecase(db *db.Database) TodoUsecase {
	usecase := &todoUsecase{
		uow:      uow.NewUnitOfWork(db),
		validate: validator.New(),
	}
	usecase.todoRepo = usecase.uow.Todos()
	return usecase
}

func (uc *todoUsecase) GetPagedList(
	ctx context.Context,
	param pagination.PagingParam,
) (*pagination.PagedList, error) {
	data, err := uc.todoRepo.GetPagedList(ctx, param)
	if err != nil {
		return nil, err
	}

	// Convert Todo object to Todo obj
	var obj []dto.TodoDto
	if err := mapper.MapperSlice(data.Data, &obj); err != nil {
		return nil, err
	}
	data.Data = obj

	return data, nil
}

func (uc *todoUsecase) GetById(ctx context.Context, id string) (*dto.TodoDto, error) {
	data, err := uc.todoRepo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	// Convert Todo object to Todo dto
	obj := &dto.TodoDto{}
	if err := mapper.Mapper(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (uc *todoUsecase) Create(ctx context.Context, form form.TodoForm) (*dto.TodoDto, error) {
	// Validate Todo form
	if err := uc.validate.Struct(form); err != nil {
		return nil, err
	}

	// Convert Todo form to Todo object
	data := &model.Todo{}
	if err := mapper.Mapper(&form, data); err != nil {
		return nil, err
	}

	if err := uc.todoRepo.Create(ctx, data); err != nil {
		return nil, err
	}

	// Convert Todo object to Todo dto
	obj := &dto.TodoDto{}
	if err := mapper.AutoMapper(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (uc *todoUsecase) Update(ctx context.Context, id string, form form.TodoForm) error {
	// Validate Todo form
	if err := uc.validate.Struct(form); err != nil {
		return err
	}

	// Get Todo object by ID
	data, err := uc.todoRepo.GetById(ctx, id)
	if err != nil {
		return err
	}

	// Override Todo form into above Todo object
	if err := mapper.Mapper(&form, data); err != nil {
		return err
	}

	return uc.todoRepo.Update(ctx, data)
}

func (uc *todoUsecase) Delete(ctx context.Context, id string) error {
	return uc.todoRepo.Delete(ctx, id)
}
