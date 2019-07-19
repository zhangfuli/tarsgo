package db

import (
	"fmt"
	"github.com/microservices-demo/user/users"
	"gopkg.in/mgo.v2"
)

// Database represents a simple interface so we can switch to a new system easily
// this is just basic and specific to this microservice
type Database interface {
	Init() error
	GetUserByName(string) (users.User, error)
	GetUser(string) (users.User, error)
	GetUsers() ([]users.User, error)
	CreateUser(*users.User) error
	GetUserAttributes(*users.User) error
	GetAddress(string) (users.Address, error)
	GetAddresses() ([]users.Address, error)
	CreateAddress(*users.Address, string) error
	GetCard(string) (users.Card, error)
	GetCards() ([]users.Card, error)
	Delete(string, string) error
	CreateCard(*users.Card, string) error
	Ping() error
}


//Init inits the selected DB in DefaultDb
func Init()(*mgo.Session,error) {
	session,err:=mgo.Dial("mongodb://localhost:27017")
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println("connect success")
	}
	return session, err
}

