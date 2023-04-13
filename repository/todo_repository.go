package repository

import (
	"context"
	"example/common/pagination"
	"example/db"
	"example/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoRepository interface {
	GetPagedList(ctx context.Context, param pagination.PagingParam) (*pagination.PagedList, error)
	GetById(ctx context.Context, id string) (*model.Todo, error)
	Create(ctx context.Context, data *model.Todo) error
	Update(ctx context.Context, data *model.Todo) error
	Delete(ctx context.Context, id string) error
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

func (repo *todoRepository) GetPagedList(ctx context.Context, param pagination.PagingParam) (*pagination.PagedList, error) {
	// Get cursor, count of documents and error
	cur, count, err := pagination.Pagination(repo.collection, ctx, param, bson.M{})
	if err != nil {
		return nil, err
	}

	// Read data from cursor, and decode to Todo list
	var data []model.Todo
	if err := cur.All(ctx, &data); err != nil {
		return nil, err
	}

	return pagination.NewPagedList(data, param.Page, param.Limit, count), nil
}

func (repo *todoRepository) GetById(ctx context.Context, id string) (*model.Todo, error) {
	_id, _ := primitive.ObjectIDFromHex(id)
	res := repo.collection.FindOne(ctx, bson.M{"_id": _id})

	// Decode above result to Todo object
	var data model.Todo
	if err := res.Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (repo *todoRepository) Create(ctx context.Context, data *model.Todo) error {
	// Generate ObjectID
	data.Id = primitive.NewObjectID()
	// Assign timestamps
	data.CreatedAt, data.UpdatedAt = time.Now(), time.Now()

	_, err := repo.collection.InsertOne(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (repo *todoRepository) Update(ctx context.Context, data *model.Todo) error {
	// Assign timestamps
	data.UpdatedAt = time.Now()

	_, err := repo.collection.UpdateByID(ctx, data.Id, bson.M{"$set": data})
	if err != nil {
		return err
	}

	return nil
}

func (repo *todoRepository) Delete(ctx context.Context, id string) error {
	_id, _ := primitive.ObjectIDFromHex(id)

	_, err := repo.collection.DeleteOne(ctx, bson.M{"_id": _id})
	if err != nil {
		return err
	}

	return nil
}
