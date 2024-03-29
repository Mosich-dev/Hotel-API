package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/Mosich-dev/Hotel-API/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const USERCOLL = "users"

type Map map[string]any

type Droppable interface {
	Drop(ctx context.Context) error
}

type UserStore interface {
	Droppable
	GetUserByID(ctx context.Context, id string) (*types.User, error)
	GetUsers(ctx context.Context) ([]*types.User, error)
	InsertUser(ctx context.Context, user *types.User) (*types.User, error)
	DeleteUser(ctx context.Context, userID string) error
	UpdateUser(ctx context.Context, id string, params types.UpdateUserParams) error
}

type MongoUserStore struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client, DBName string) *MongoUserStore {
	return &MongoUserStore{
		client:     client,
		collection: client.Database(DBName).Collection(USERCOLL),
	}
}

func (s MongoUserStore) Drop(ctx context.Context) error {
	fmt.Println("-------Dropping User Collection-------")
	return s.collection.Drop(ctx)
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
	deleteResult, err := s.collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	if deleteResult.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

// UpdateUser TODO: refactor
func (s MongoUserStore) UpdateUser(ctx context.Context, id string, params types.UpdateUserParams) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	values := params.ToBsonM()
	if values.M["firstName"] != nil {
		if values.Len("firstName") < types.MinFirstName {
			return errors.New(fmt.Sprintf("first name length most be atleast %d characters", types.MinFirstName))
		}
	}

	if values.M["lastName"] != nil {
		if values.Len("lastName") < types.MinLastName {
			return errors.New(fmt.Sprintf("last name length most be atleast %d characters", types.MinLastName))
		}
	}
	update := bson.D{{"$set", values.M}}
	_, err = s.collection.UpdateByID(ctx, oid, update)
	if err != nil {
		return err
	}
	return nil
}
