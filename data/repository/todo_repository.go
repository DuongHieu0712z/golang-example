package repository

import (
	"context"
	"errors"
	"example/common/base"
	"example/common/errs"
	"example/common/pagination"
	"example/data/entity"
	"example/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoRepository interface {
	GetPagination(ctx context.Context, params pagination.PagingParams) *pagination.PagedList
	GetById(ctx context.Context, id string) *entity.Todo
	Create(ctx context.Context, data *entity.Todo)
	Update(ctx context.Context, data *entity.Todo)
	Delete(ctx context.Context, id string)
}

type todoRepository struct {
	base.Repository
}

func NewTodoRepository(db *db.Database) TodoRepository {
	return &todoRepository{
		Repository: *base.NewRepository(db, "todo"),
	}
}

func (repo todoRepository) GetPagination(
	ctx context.Context,
	params pagination.PagingParams,
) *pagination.PagedList {
	cur, count, err := pagination.Pagination(ctx, repo.Collection, params, bson.M{})
	if err != nil {
		panic(errs.BadRequestError(err))
	}

	var data []entity.Todo
	if err := cur.All(ctx, &data); err != nil {
		panic(errs.BadRequestError(err))
	}

	return pagination.NewPagedList(data, params.Page, params.Limit, count)
}

func (repo todoRepository) GetById(ctx context.Context, id string) *entity.Todo {
	_id, _ := primitive.ObjectIDFromHex(id)
	cur := repo.Collection.FindOne(ctx, bson.M{"_id": _id})

	data := &entity.Todo{Entity: base.Entity{}}
	if err := cur.Decode(data); err != nil {
		panic(errs.BadRequestError(err))
	}

	return data
}

func (repo todoRepository) Create(ctx context.Context, data *entity.Todo) {
	data.SetTime(true)

	result, err := repo.Collection.InsertOne(ctx, data)
	if err != nil {
		panic(errs.BadRequestError(err))
	}

	data.SetId(result.InsertedID)
}

func (repo todoRepository) Update(ctx context.Context, data *entity.Todo) {
	data.SetTime(false)

	result, err := repo.Collection.UpdateByID(ctx, data.Id, bson.M{"$set": data})
	if err != nil {
		panic(errs.BadRequestError(err))
	}

	if result.MatchedCount < 1 {
		err := errors.New("mongo: no document is updated")
		panic(errs.BadRequestError(err))
	}
}

func (repo todoRepository) Delete(ctx context.Context, id string) {
	_id, _ := primitive.ObjectIDFromHex(id)

	result, err := repo.Collection.DeleteOne(ctx, bson.M{"_id": _id})
	if err != nil {
		panic(errs.BadRequestError(err))
	}

	if result.DeletedCount < 1 {
		err := errors.New("mongo: no document is deleted")
		panic(errs.BadRequestError(err))
	}
}
