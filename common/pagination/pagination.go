package pagination

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Pagination(
	ctx context.Context,
	coll *mongo.Collection,
	params PagingParams,
	filter interface{},
	opts ...*options.FindOptions,
) (*mongo.Cursor, int64, error) {
	opt := options.Find()
	opt.SetSkip((params.Page - 1) * params.Limit)
	opt.SetLimit(params.Limit)

	order := 1
	if params.Order == "desc" {
		order = -1
	}
	opt.SetSort(bson.D{{Key: params.Field, Value: order}})

	opts = append(opts, opt)

	cur, err := coll.Find(ctx, filter, opts...)
	if err != nil {
		return nil, 0, err
	}

	count, err := coll.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	return cur, count, nil
}
