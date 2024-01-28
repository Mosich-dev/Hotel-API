package API

import (
	"context"
	"github.com/Mosich-dev/Hotel-API/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"testing"
)

const (
	testDbUri  = "mongodb://localhost:27017"
	testDbName = "hotel-api-test"
)

type testdDb struct {
	db.UserStore
}

func (tdb testdDb) tearDown(t *testing.T) {
	if err := tdb.UserStore.Drop(context.Background()); err != nil {
		t.Fatal(err)
	}
}

func setup(t *testing.T) *testdDb {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(testDbUri))
	if err != nil {
		log.Fatal(err)
	}

	return &testdDb{
		UserStore: db.NewMongoUserStore(client, testDbName),
	}
}

func TestInsertUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.tearDown(t)
	t.Fail()
}
