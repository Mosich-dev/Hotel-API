package db

import (
	"context"
	"github.com/Mosich-dev/Hotel-API/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const USERCOLL = "users"

type UserStore interface {
	GetUserByID(context.Context, string) (*types.User, error)
}

type MongoUserStore struct {
	client     *mongo.Client
	dbname     string
	collection *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		client:     client,
		dbname:     DBNAME,
		collection: client.Database(DBNAME).Collection(USERCOLL),
	}
}

func (s MongoUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	var (
		user types.User
		oid  any
		err  error
	)

	if IsObjedID(id) {
		oid, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			panic(err)
		}
	} else {
		oid = id
	}

	if err := s.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
