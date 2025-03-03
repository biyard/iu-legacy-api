package mongodb

import (
	"context"
	"errors"

	"biyard.co/common/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CollectionDescriber interface {
	Collection() string
	Indexes() []mongo.IndexModel
}

type MongoDB struct {
	*mongo.Database

	cli *mongo.Client
	ctx context.Context
}

func New(ctx context.Context, uri, db string) *MongoDB {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI))
	if err != nil {
		panic(err)
	}

	d := client.Database(db)

	return &MongoDB{
		cli:      client,
		ctx:      ctx,
		Database: d,
	}
}

func (r *MongoDB) Close() {
	r.cli.Disconnect(r.ctx)
}

func NewCollection[T CollectionDescriber](ctx context.Context, db *MongoDB) (*Collection[T], error) {
	var doc T

	c := db.Database.Collection(doc.Collection())
	if c == nil {
		return nil, errors.New("collection not found")
	}
	if i := doc.Indexes(); len(i) > 0 {
		_, err := c.Indexes().CreateMany(ctx, i)
		if err != nil {
			return nil, err
		}
	}

	return &Collection[T]{
		Collection: c,
	}, nil
}

type Collection[T CollectionDescriber] struct {
	*mongo.Collection
}

func (r *Collection[T]) FindOne(ctx context.Context, filter any) (T, error) {
	var doc T
	logger.Debugf(ctx, "FindOne %+v", r.Collection)
	ret := r.Collection.FindOne(ctx, filter)
	if ret.Err() != nil {
		return doc, ret.Err()
	}

	ret.Decode(&doc)

	return doc, nil
}

func (r *Collection[T]) Find(ctx context.Context, filter any) ([]T, error) {
	docs := make([]T, 0)
	cur, err := r.Collection.Find(ctx, filter)
	if err != nil {
		return docs, err
	}
	err = cur.All(ctx, &docs)

	return docs, err
}

func (r *Collection[T]) Insert(ctx context.Context, doc T) error {
	_, err := r.Collection.InsertOne(ctx, doc)

	return err
}

func (r *Collection[T]) InsertMany(ctx context.Context, docs []T) error {
	newValue := make([]interface{}, len(docs))
	for i := range docs {
		newValue[i] = docs[i]
	}
	_, err := r.Collection.InsertMany(ctx, newValue)
	return err
}
