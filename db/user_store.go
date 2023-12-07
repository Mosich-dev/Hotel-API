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
	GetUsers(context.Context) ([]*types.User, error)
	InsertUser(ctx context.Context, user *types.User) (*types.User, error)
	DeleteUser(ctx context.Context, userID string) error
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

func (s MongoUserStore) GetUsers(ctx context.Context) ([]*types.User, error) {
	cur, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var user []*types.User
	if err := cur.All(ctx, &user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s MongoUserStore) InsertUser(ctx context.Context, user *types.User) (*types.User, error) {
	result, err := s.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (s MongoUserStore) DeleteUser(ctx context.Context, userID string) error {
	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	_, err = s.collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	return nil
}
