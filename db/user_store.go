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
	GetUsers(context.Context) ([]*types.User, error)
	CreateUser(context.Context, *types.User) (*types.User, error)
	DeleteUser(context.Context, string) (*types.User, error)
	UpdateUser(c context.Context, filter bson.M, params types.UpdateUserParams) error
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

func (s *MongoUserStore) CreateUser(ctx context.Context, user *types.User) (*types.User, error) {
	res, err := s.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = res.InsertedID.(primitive.ObjectID) //manipulate the pointer and retern it. So every func using this user pointer will see the changes
	return user, nil
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

func (s *MongoUserStore) UpdateUser(ctx context.Context, filter bson.M, params types.UpdateUserParams) error {
	update := bson.D{
		{
			"$set", params.ToBSON(),
		},
	}

	_, err := s.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (s *MongoUserStore) DeleteUser(ctx context.Context, id string) (*types.User, error) {
	//validate id
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// пока что не нашел способ удалить юзера и вернуть его данные без 2х запросов к коллекции
	var user types.User
	if err := s.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&user); err != nil {
		return nil, err
	}

	_, err = s.coll.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *MongoUserStore) GetUsers(ctx context.Context) ([]*types.User, error) {
	cur, err := s.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var users []*types.User
	if err := cur.All(ctx, &users); err != nil {
		return []*types.User{}, nil
	}
	return users, nil
}
