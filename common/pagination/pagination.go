package pagination

import (
	"context"
	"example/common/errs"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Pagination(
	coll *mongo.Collection,
	ctx context.Context,
	params PagingParams,
	filter interface{},
	opts ...*options.FindOptions,
) (*mongo.Cursor, int64) {
	opt := options.Find()
	opt.SetSkip((params.Page - 1) * params.Limit)
	opt.SetLimit(params.Limit)
	opt.SetSort(bson.D{{Key: params.Field, Value: params.Order}})

	opts = append(opts, opt)

	cur, err := coll.Find(ctx, filter, opts...)
	if err != nil {
		panic(errs.BadRequestError(err))
	}

	count, err := coll.CountDocuments(ctx, filter)
	if err != nil {
		panic(errs.BadRequestError(err))
	}

	return cur, count
}
