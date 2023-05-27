package repository

import (
	"context"
	"example/common/base"
	"example/common/errs"
	"example/data/entity"
	"example/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	GetById(ctx context.Context, id string) *entity.User
	GetByUsername(ctx context.Context, username string) *entity.User
	Create(ctx context.Context, data *entity.User)
}

type userRepository struct {
	base.Repository
}

func NewUserRepository(db *db.Database) UserRepository {
	return &userRepository{
		Repository: *base.NewRepository(db, "user"),
	}
}

func (repo userRepository) GetById(ctx context.Context, id string) *entity.User {
	_id, _ := primitive.ObjectIDFromHex(id)
	cur := repo.Collection.FindOne(ctx, bson.M{"_id": _id})

	data := &entity.User{}
	if err := cur.Decode(data); err != nil {
		panic(errs.BadRequestError(err))
	}

	return data
}

func (repo userRepository) GetByUsername(ctx context.Context, username string) *entity.User {
	cur := repo.Collection.FindOne(ctx, bson.M{"username": username})

	data := &entity.User{}
	if err := cur.Decode(data); err != nil {
		panic(errs.BadRequestError(err))
	}

	return data
}

func (repo userRepository) Create(ctx context.Context, data *entity.User) {
	data.SetTime(true)

	result, err := repo.Collection.InsertOne(ctx, data)
	if err != nil {
		panic(errs.BadRequestError(err))
	}

	data.SetId(result.InsertedID)
}
