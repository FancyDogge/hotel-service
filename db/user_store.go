package db

import "github.com/FancyDogge/hotel-service/types"

//т.к. мы создаем интерфейс, потом можно будет сделать хоть postgres юзер стор, хоть mongodb, хоть in-memory
type UserStore interface {
	GetUserByID(string) (*types.User, error)
}

//Реализация интерфейса UserStore, заточенная под монго
type MongoUserStore struct {
}
