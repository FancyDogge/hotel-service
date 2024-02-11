package api

import (
	"context"
	"log"
	"testing"

	"github.com/FancyDogge/hotel-service/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const testMongoUri = "mongodb://admin:admin@localhost:27017/"

type testdb struct {
	db.UserStore
}

func setup(t *testing.T) *testdb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testMongoUri))
	if err != nil {
		log.Fatal(err)
	}

	return &testdb{
		UserStore: db.NewMongoUserStore(client),
	}
}

func TestPostUser(t *testing.T) {
	t.Fail()
}
