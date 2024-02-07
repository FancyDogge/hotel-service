package db

import (
	"context"

	"github.com/FancyDogge/hotel-service/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const userColl = "users"

// т.к. мы создаем интерфейс, потом можно будет сделать хоть postgres юзер стор, хоть mongodb, хоть in-memory
type UserStore interface {
	GetUserByID(context.Context, string) (*types.User, error)
}

// Реализация интерфейса UserStore, заточенная под монго
// Содержит клиент собственно монгодб
// Store == Table*
type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

// Реализация/активация MongoUserStore, чтобы создать его и назначить в переменную. возвращает поинтер к монгобзерстор
func NewMongoUserStore(client *mongo.Client) *MongoUserStore {

	return &MongoUserStore{
		client: client,
		coll:   client.Database(DBNAME).Collection(userColl),
	}
}

// UserStore implementation in MongoUserStore
func (s *MongoUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	//validate id
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user types.User
	if err := s.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&user); err != nil { //user as a pointer, чтобы после прогона ф-ции были изменения в юзере
		return nil, err //empty pointer == nil
	}
	return &user, nil
}
