package API

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/Mosich-dev/Hotel-API/db"
	"github.com/Mosich-dev/Hotel-API/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	testDbUri  = "mongodb://localhost:27017"
	testDbName = "hotel-api-test"
)

type testDb struct {
	db.UserStore
}

func (tdb *testDb) tearDown(t *testing.T) {
	if err := tdb.UserStore.Drop(context.Background()); err != nil {
		t.Fatal(err)
	}
}

func setup(t *testing.T) *testDb {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(testDbUri))
	if err != nil {
		t.Error(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		t.Errorf("Error in ping: %s", err)
	}
	return &testDb{
		UserStore: db.NewMongoUserStore(client, testDbName),
	}
}

func TestInsertUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.tearDown(t)
	params := types.InsertUserParams{
		FirstName: "reza",
		LastName:  "chegeni",
		Email:     "rrr.reza18@gmail.com",
		Password:  "1234567",
	}
	app := fiber.New()
	userHandler := NewUserHandler(tdb.UserStore)
	app.Post("/", userHandler.HandleInsertUser)
	var user types.User
	correctRes, err := mockUser(app, &params)
	err = json.NewDecoder(correctRes.Body).Decode(&user)
	if err != nil {
		t.Error(err)
	}

	if err != nil {
		t.Error(err)
	}
	if correctRes.StatusCode != fiber.StatusCreated {
		t.Errorf("response status code wrong. Expected status code: %d Received: %d", fiber.StatusCreated, correctRes.StatusCode)
	}
	if len(user.ID) == 0 {
		t.Errorf("expecting a user id to be set")
	}
	if len(user.EncryptedPassword) > 0 {
		t.Errorf("expecting the EncryptedPassword not to be included in the json response")
	}
	if user.FirstName != params.FirstName {
		t.Errorf("expected firstname %s but got %s", params.FirstName, user.FirstName)
	}
	if user.LastName != params.LastName {
		t.Errorf("expected last name %s but got %s", params.LastName, user.LastName)
	}
	if user.Email != params.Email {
		t.Errorf("expected email %s but got %s", params.Email, user.Email)
	}

}

func mockUser(app *fiber.App, params *types.InsertUserParams) (*http.Response, error) {
	b, _ := json.Marshal(params)
	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	res, err := app.Test(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
