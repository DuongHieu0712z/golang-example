package repository

import (
	"context"
	"example/common/errs"
	"example/common/pagination"
	"example/db"
	"example/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoRepository interface {
	GetPagedList(ctx context.Context, params pagination.PagingParams) *pagination.PagedList
	GetById(ctx context.Context, id string) *model.Todo
	Create(ctx context.Context, data *model.Todo)
	Update(ctx context.Context, data *model.Todo)
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
	// Get cursor, count of documents and
	cur, count := pagination.Pagination(repo.collection, ctx, params, bson.M{})

	// Read data from cursor, and decode to Todo list
	var data []model.Todo
	if err := cur.All(ctx, &data); err != nil {
		panic(errs.BadRequestError(err))
	}

	return pagination.NewPagedList(data, params.Page, params.Limit, count)
}

func (repo *todoRepository) GetById(ctx context.Context, id string) *model.Todo {
	_id, _ := primitive.ObjectIDFromHex(id)
	res := repo.collection.FindOne(ctx, bson.M{"_id": _id})

	// Decode above result to Todo object
	data := &model.Todo{}
	if err := res.Decode(data); err != nil {
		panic(errs.BadRequestError(err))
	}

	return data
}

func (repo *todoRepository) Create(ctx context.Context, data *model.Todo) {
	// Assign timestamps
	data.CreatedAt, data.UpdatedAt = time.Now(), time.Now()

	result, err := repo.collection.InsertOne(ctx, data)
	if err != nil {
		panic(errs.BadRequestError(err))
	}

	data.Id = result.InsertedID.(primitive.ObjectID)
}

func (repo *todoRepository) Update(ctx context.Context, data *model.Todo) {
	// Assign timestamps
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
