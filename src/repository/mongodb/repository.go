package mongodb

import (
	"context"
	"time"

	"github.com/andrei-dascalu/shortener/src/shortener"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoRepository struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

func newMongoClient(mongoURL string, mongoTimeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(mongoTimeout)*time.Second,
	)

	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))

	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to mongo")
	}

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		return nil, errors.Wrap(err, "failed to set primary instance")
	}

	return client, nil
}

func NewMongoRepository(mongoURL, mongoDB string) (shortener.LinkRepository, error) {
	timeout := 2000
	repo := &mongoRepository{
		timeout:  time.Duration(timeout) * time.Second,
		database: mongoDB,
	}

	client, err := newMongoClient(mongoURL, timeout)

	if err != nil {
		return nil, errors.Wrap(err, "Failed to create client")
	}

	repo.client = client

	return repo, nil
}

func (r *mongoRepository) Find(code string) (*shortener.LinkRedirect, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)

	defer cancel()

	redirect := &shortener.LinkRedirect{}

	collection := r.client.Database(r.database).Collection("redirects")
	filter := bson.M{"code": code}

	err := collection.FindOne(ctx, filter).Decode(redirect)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(err, "Document not found")
		}

		return nil, errors.Wrap(err, "Generic mongo error")
	}

	return redirect, nil
}

func (r *mongoRepository) Store(link *shortener.LinkRedirect) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)

	log.Warn().Msg("Using Mongo Store")

	defer cancel()
	collection := r.client.Database(r.database).Collection("redirects")

	_, err := collection.InsertOne(
		ctx,
		bson.M{
			"code":       link.Code,
			"url":        link.URL,
			"created_at": link.CreatedAt,
		},
	)

	if err != nil {
		return errors.Wrap(err, "Could not insert")
	}

	return nil
}
