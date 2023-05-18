package usecase

import (
	"context"
	"example/common/errs"
	"example/common/pagination"
	"example/data/entity"
	"example/data/repository"
	"example/data/uow"
	"example/db"
	"example/service/request"
	"example/service/response"

	"github.com/devfeel/mapper"
)

type TodoUsecase interface {
	GetPagedList(ctx context.Context, params pagination.PagingParams) *pagination.PagedList
	GetById(ctx context.Context, id string) *response.TodoResponse
	Create(ctx context.Context, request request.TodoRequest) *response.TodoResponse
	Update(ctx context.Context, id string, request request.TodoRequest)
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

	var obj []response.TodoResponse
	if err := mapper.MapperSlice(data.Data, &obj); err != nil {
		panic(errs.BadRequestError(err))
	}
	data.Data = obj

	return data
}

func (uc *todoUsecase) GetById(ctx context.Context, id string) *response.TodoResponse {
	data := uc.todoRepo.GetById(ctx, id)

	obj := &response.TodoResponse{}
	if err := mapper.Mapper(data, obj); err != nil {
		panic(errs.BadRequestError(err))
	}
	return obj
}

func (uc *todoUsecase) Create(ctx context.Context, request request.TodoRequest) *response.TodoResponse {
	data := &entity.Todo{}
	if err := mapper.Mapper(&request, data); err != nil {
		panic(errs.BadRequestError(err))
	}

	uc.todoRepo.Create(ctx, data)

	obj := &response.TodoResponse{}
	if err := mapper.AutoMapper(data, obj); err != nil {
		panic(errs.BadRequestError(err))
	}
	return obj
}

func (uc *todoUsecase) Update(ctx context.Context, id string, request request.TodoRequest) {
	data := uc.todoRepo.GetById(ctx, id)

	if err := mapper.Mapper(&request, data); err != nil {
		panic(errs.BadRequestError(err))
	}

	uc.todoRepo.Update(ctx, data)
}

func (uc *todoUsecase) Delete(ctx context.Context, id string) {
	uc.todoRepo.Delete(ctx, id)
}
