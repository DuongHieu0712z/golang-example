package usecase

import (
	"context"
	"example/common/errs"
	"example/common/pagination"
	"example/db"
	"example/dto"
	"example/form"
	"example/model"
	"example/repository"
	"example/uow"

	"github.com/devfeel/mapper"
)

type TodoUsecase interface {
	GetPagedList(ctx context.Context, params pagination.PagingParams) *pagination.PagedList
	GetById(ctx context.Context, id string) *dto.TodoDto
	Create(ctx context.Context, form form.TodoForm) *dto.TodoDto
	Update(ctx context.Context, id string, form form.TodoForm)
	Delete(ctx context.Context, id string)
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

func (uc *todoUsecase) GetPagedList(
	ctx context.Context,
	params pagination.PagingParams,
) *pagination.PagedList {
	data := uc.todoRepo.GetPagedList(ctx, params)

	// Convert Todo object to Todo obj
	var obj []dto.TodoDto
	if err := mapper.MapperSlice(data.Data, &obj); err != nil {
		panic(errs.BadRequestError(err))
	}
	data.Data = obj

	return data
}

func (uc *todoUsecase) GetById(ctx context.Context, id string) *dto.TodoDto {
	data := uc.todoRepo.GetById(ctx, id)

	// Convert Todo object to Todo dto
	obj := &dto.TodoDto{}
	if err := mapper.Mapper(data, obj); err != nil {
		panic(errs.BadRequestError(err))
	}
	return obj
}

func (uc *todoUsecase) Create(ctx context.Context, form form.TodoForm) *dto.TodoDto {
	// Convert Todo form to Todo object
	data := &model.Todo{}
	if err := mapper.Mapper(&form, data); err != nil {
		panic(errs.BadRequestError(err))
	}

	uc.todoRepo.Create(ctx, data)

	// Convert Todo object to Todo dto
	obj := &dto.TodoDto{}
	if err := mapper.AutoMapper(data, obj); err != nil {
		panic(errs.BadRequestError(err))
	}
	return obj
}

func (uc *todoUsecase) Update(ctx context.Context, id string, form form.TodoForm) {
	// Get Todo object by ID
	data := uc.todoRepo.GetById(ctx, id)

	// Override Todo form into above Todo object
	if err := mapper.Mapper(&form, data); err != nil {
		panic(errs.BadRequestError(err))
	}

	uc.todoRepo.Update(ctx, data)
}

func (uc *todoUsecase) Delete(ctx context.Context, id string) {
	uc.todoRepo.Delete(ctx, id)
}
