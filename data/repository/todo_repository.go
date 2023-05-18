package repository

import (
	"context"
	"example/common/errs"
	"example/common/pagination"
	"example/data/entity"
	"example/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoRepository interface {
	GetPagedList(ctx context.Context, params pagination.PagingParams) *pagination.PagedList
	GetById(ctx context.Context, id string) *entity.Todo
	Create(ctx context.Context, data *entity.Todo)
	Update(ctx context.Context, data *entity.Todo)
	Delete(ctx context.Context, id string)
}

type todoRepository struct {
	name       string
	db         *db.Database
	collection *mongo.Collection
}

func NewTodoRepository(db *db.Database) TodoRepository {
	repo := &todoRepository{
		name: "todo",
		db:   db,
	}
	repo.collection = db.GetCollection(repo.name)
	return repo
}

func (repo *todoRepository) GetPagedList(
	ctx context.Context,
	params pagination.PagingParams,
) *pagination.PagedList {
	cur, count, err := pagination.Pagination(ctx, repo.collection, params, bson.M{})
	if err != nil {
		panic(errs.BadRequestError(err))
	}

	var data []entity.Todo
	if err := cur.All(ctx, &data); err != nil {
		panic(errs.BadRequestError(err))
	}

	return pagination.NewPagedList(data, params.Page, params.Limit, count)
}

func (repo *todoRepository) GetById(ctx context.Context, id string) *entity.Todo {
	_id, _ := primitive.ObjectIDFromHex(id)
	res := repo.collection.FindOne(ctx, bson.M{"_id": _id})

	data := &entity.Todo{}
	if err := res.Decode(data); err != nil {
		panic(errs.BadRequestError(err))
	}

	return data
}

func (repo *todoRepository) Create(ctx context.Context, data *entity.Todo) {
	data.CreatedAt, data.UpdatedAt = time.Now(), time.Now()

	result, err := repo.collection.InsertOne(ctx, data)
	if err != nil {
		panic(errs.BadRequestError(err))
	}

	data.Id = result.InsertedID.(primitive.ObjectID)
}

func (repo *todoRepository) Update(ctx context.Context, data *entity.Todo) {
	data.UpdatedAt = time.Now()

	_, err := repo.collection.UpdateByID(ctx, data.Id, bson.M{"$set": data})
	if err != nil {
		panic(errs.BadRequestError(err))
	}
}

func (repo *todoRepository) Delete(ctx context.Context, id string) {
	_id, _ := primitive.ObjectIDFromHex(id)

	_, err := repo.collection.DeleteOne(ctx, bson.M{"_id": _id})
	if err != nil {
		panic(errs.BadRequestError(err))
	}
}
