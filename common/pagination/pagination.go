package pagination

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Pagination(
	coll *mongo.Collection,
	ctx context.Context,
	param PagingParam,
	filter interface{},
	opts ...*options.FindOptions,
) (*mongo.Cursor, int64, error) {
	opt := options.Find()
	opt.SetSkip((param.Page - 1) * param.Limit)
	opt.SetLimit(param.Limit)

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
