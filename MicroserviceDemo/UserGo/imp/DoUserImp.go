package imp

import (
	"MicroserviceDemo/UserGo/db"
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io"
	"strconv"
	"time"
)

type DoUserImp struct {
}

type Card struct {
	LongNum string        `json:"longNum" bson:"longNum, omitempty"`
	Expires string        `json:"expires" bson:"expires, omitempty"`
	CCV     string        `json:"ccv" bson:"ccv,omitempty"`
	ID      bson.ObjectId `json:"_id" bson:"_id"`
}
type CardDetail struct {
	LongNum string        `json:"longNum" bson:"longNum"`
	Expires string        `json:"expires" bson:"expires"`
	CCV     string        `json:"ccv" bson:"ccv"`
	ID      bson.ObjectId `json:"_id" bson:"_id"`
}
type Address struct {
	Street   string        `json:"street" bson:"street,omitempty"`
	Number   string        `json:"number" bson:"number,omitempty"`
	Country  string        `json:"country" bson:"country,omitempty"`
	City     string        `json:"city" bson:"city,omitempty"`
	PostCode string        `json:"postcode" bson:"postcode,omitempty"`
	ID       bson.ObjectId `json:"_id" bson:"_id"`
}
type AddressDetail struct {
	Street   string        `json:"street" bson:"street"`
	Number   string        `json:"number" bson:"number"`
	Country  string        `json:"country" bson:"country"`
	City     string        `json:"city" bson:"city"`
	PostCode string        `json:"postcode" bson:"postcode"`
	ID       bson.ObjectId `json:"_id" bson:"_id"`
}
type User struct {
	FirstName string        `json:"firstName" bson:"firstName"`
	LastName  string        `json:"lastName" bson:"lastName"`
	Email     string        `json:"-" bson:"email"`
	Username  string        `json:"username" bson:"username"`
	Password  string        `json:"-" bson:"password,omitempty"`
	Addresses []Address     `json:"-,omitempty" bson:"-"`
	Cards     []Card        `json:"-,omitempty" bson:"-"`
	UserID    bson.ObjectId `json:"_id" bson:"_id"`
	Salt      string        `json:"-" bson:"salt"`
}
type UserDetail struct {
	FirstName string          `json:"firstName" bson:"firstName"`
	LastName  string          `json:"lastName" bson:"lastName"`
	Email     string          `json:"-" bson:"email"`
	Username  string          `json:"username" bson:"username"`
	Password  string          `json:"-" bson:"password,omitempty"`
	Addresses []bson.ObjectId `json:"addresses" bson:"addresses"`
	Cards     []bson.ObjectId `json:"cards" bson:"cards"`
	UserID    bson.ObjectId   `json:"_id" bson:"_id"`
	Salt      string          `json:"-" bson:"salt"`
}

var (
	ErrUnauthorized = errors.New("Unauthorized")
)

type Customerror struct {
	infoa string
	infob string
	Err   error
}

func (cerr Customerror) Error() string {
	errorinfo := fmt.Sprintf("infoa : %s , infob : %s , original err info : %s ", cerr.infoa, cerr.infob, cerr.Err.Error())
	return errorinfo
}

func (imp *DoUserImp) EchoHello(name string, greeting *string) (int32, error) {
	*greeting = "hello " + name
	return 0, nil
}
func (u *User) NewSalt() {
	h := sha1.New()
	io.WriteString(h, strconv.Itoa(int(time.Now().UnixNano())))
	u.Salt = fmt.Sprintf("%x", h.Sum(nil))
}
func New() User {
	u := User{Addresses: make([]Address, 0), Cards: make([]Card, 0)}
	u.NewSalt()
	return u
}
func calculatePassHash(pass, salt string) string {
	h := sha1.New()
	io.WriteString(h, salt)
	io.WriteString(h, pass)
	return fmt.Sprintf("%x", h.Sum(nil))
}
func (imp *DoUserImp) Login(username string, password string) (string, error) {
	session, sessionErr := db.Init()
	fmt.Sprintln(session)
	if sessionErr != nil {
		fmt.Println(sessionErr)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("users").C("customers")
	var user User
	customersError := c.Find(bson.M{"username": username}).One(&user)
	if customersError != nil {
		fmt.Println(customersError)
	}
	if user.Password != calculatePassHash(password, user.Salt) {
		return ErrUnauthorized.Error(), nil
	}
	userjson, _ := json.Marshal(user)
	return string(userjson), nil
}
func (imp *DoUserImp) Register(username string, password string, email string, first string, last string) (string, error) {
	session, sessionErr := db.Init()
	fmt.Sprintln(session)
	if sessionErr != nil {
		fmt.Println(sessionErr)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("users").C("customers")

	user := User{Addresses: make([]Address, 0), Cards: make([]Card, 0)}
	user.NewSalt()
	user.Username = username
	user.Password = calculatePassHash(password, user.Salt)
	user.Email = email
	user.FirstName = first
	user.LastName = last
	user.UserID = bson.NewObjectId()
	insertErr := c.Insert(&user)
	////err := db.CreateUser(&u)
	//customersError := c.Find(bson.M{"username": username}).One(&user)
	if insertErr != nil {
		fmt.Println(insertErr)
	}
	//if user.Password != calculatePassHash(password, user.Salt) {
	//	return ErrUnauthorized.Error(), nil
	//}
	//fmt.Println(calculatePassHash(password, user.Salt))
	userjson, _ := json.Marshal(user.UserID)
	return "{\"id\" :" + string(userjson) + "}", nil
}

func (imp *DoUserImp) Customers() (string, error) {
	session, sessionErr := db.Init()
	fmt.Sprintln(session)
	if sessionErr != nil {
		fmt.Println(sessionErr)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("users").C("customers")

	var users []User
	customersError := c.Find(nil).All(&users)
	if customersError != nil {
		fmt.Println(customersError)
	}
	usersjson, _ := json.Marshal(users)
	return string(usersjson), nil
}
func (imp *DoUserImp) DeleteCustomerById(id string) (string, error) {
	session, sessionErr := db.Init()
	fmt.Sprintln(session)
	if sessionErr != nil {
		fmt.Println(sessionErr)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("users").C("customers")

	deleteError := c.RemoveId(bson.ObjectIdHex(id))
	if deleteError != nil {
		return deleteError.Error(), nil
	}
	return "{\"status\": true}", nil
}
func (imp *DoUserImp) FindCustomerById(id string) (string, error) {
	session, sessionErr := db.Init()
	fmt.Sprintln(session)
	if sessionErr != nil {
		fmt.Println(sessionErr)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("users").C("customers")

	var customer User
	findError := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&customer)
	if findError != nil {
		return findError.Error(), nil
	}
	customerjson, _ := json.Marshal(customer)
	return string(customerjson), nil
}
func (imp *DoUserImp) FindCustomerCardById(id string) (string, error) {
	session, sessionErr := db.Init()
	fmt.Sprintln(session)
	if sessionErr != nil {
		fmt.Println(sessionErr)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("users").C("customers")

	var customer UserDetail
	findCustomerError := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&customer)
	if findCustomerError != nil {
		return findCustomerError.Error(), nil
	}

	var cards []CardDetail
	for _, s := range customer.Cards {
		var card CardDetail
		cardDB := session.DB("users").C("cards")
		fmt.Println(s)
		findCardError := cardDB.Find(bson.M{"_id": s}).One(&card)
		fmt.Println(findCardError)
		fmt.Println(cards)
		cards = append(cards, card)
	}
	cardsjson, _ := json.Marshal(cards)
	return string(cardsjson), nil
}
func (imp *DoUserImp) FindCustomerAddressById(id string) (string, error) {
	session, sessionErr := db.Init()
	fmt.Sprintln(session)
	if sessionErr != nil {
		fmt.Println(sessionErr)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("users").C("customers")

	var customer UserDetail
	findCustomerError := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&customer)
	if findCustomerError != nil {
		return findCustomerError.Error(), nil
	}

	var cards []AddressDetail
	for _, s := range customer.Addresses {
		var card AddressDetail
		cardDB := session.DB("users").C("addresses")
		fmt.Println(s)
		findCardError := cardDB.Find(bson.M{"_id": s}).One(&card)
		fmt.Println(findCardError)
		fmt.Println(cards)
		cards = append(cards, card)
	}
	cardsjson, _ := json.Marshal(cards)
	return string(cardsjson), nil
}

func (imp *DoUserImp) Cards() (string, error) {
	session, sessionErr := db.Init()
	fmt.Sprintln(session)
	if sessionErr != nil {
		fmt.Println(sessionErr)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("users").C("cards")

	var users []CardDetail
	customersError := c.Find(nil).All(&users)
	if customersError != nil {
		fmt.Println(customersError)
	}
	usersjson, _ := json.Marshal(users)
	return string(usersjson), nil
}
func (imp *DoUserImp) AddCard(longNum string, expires string, ccv string, userID string) (string, error) {
	session, sessionErr := db.Init()
	fmt.Sprintln(session)
	if sessionErr != nil {
		fmt.Println(sessionErr)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("users").C("cards")
	update := session.DB("users").C("customers")
	cardId := bson.NewObjectId()
	card := CardDetail{longNum, expires, ccv, cardId}
	insertErr := c.Insert(&card)
	if insertErr != nil {
		return insertErr.Error(), nil
	}

	selector := bson.M{"_id": bson.ObjectIdHex(userID)}
	data := bson.M{"$addToSet": bson.M{"cards": cardId}}
	updateError := update.Update(selector, data)
	if updateError != nil {
		return updateError.Error(), nil
	}
	cardIdJson, _ := json.Marshal(cardId)
	return "{\"id\" :" + string(cardIdJson) + "}", nil
}
func (imp *DoUserImp) DeleteCardById(id string) (string, error) {
	session, sessionErr := db.Init()
	if sessionErr != nil {
		fmt.Println(sessionErr)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("users").C("cards")

	deleteError := c.RemoveId(bson.ObjectIdHex(id))
	if deleteError != nil {
		return deleteError.Error(), nil
	}
	return "{\"status\": true}", nil
}
func (imp *DoUserImp) FindCardById(id string) (string, error) {
	session, sessionErr := db.Init()
	fmt.Sprintln(session)
	if sessionErr != nil {
		fmt.Println(sessionErr)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("users").C("cards")

	var card CardDetail
	findCustomerError := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&card)
	if findCustomerError != nil {
		return findCustomerError.Error(), nil
	}
	cardjson, _ := json.Marshal(card)
	return string(cardjson), nil
}

func (imp *DoUserImp) Addresses() (string, error) {
	session, sessionErr := db.Init()
	fmt.Sprintln(session)
	if sessionErr != nil {
		fmt.Println(sessionErr)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("users").C("addresses")

	var addresses []AddressDetail
	customersError := c.Find(nil).All(&addresses)
	if customersError != nil {
		fmt.Println(customersError)
	}
	addressesjson, _ := json.Marshal(addresses)
	return string(addressesjson), nil
}
func (imp *DoUserImp) AddAddress(street string, number string, country string, city string, postcode string, userID string) (string, error) {
	session, sessionErr := db.Init()
	fmt.Sprintln(session)
	if sessionErr != nil {
		fmt.Println(sessionErr)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("users").C("addresses")
	update := session.DB("users").C("customers")
	addressesId := bson.NewObjectId()
	addresses := AddressDetail{street , number , country , city , postcode,addressesId }
	insertErr := c.Insert(&addresses)
	if insertErr != nil {
		return insertErr.Error(), nil
	}

	selector := bson.M{"_id": bson.ObjectIdHex(userID)}
	data := bson.M{"$addToSet": bson.M{"addresses": addressesId}}
	updateError := update.Update(selector, data)
	if updateError != nil {
		return updateError.Error(), nil
	}
	addressesIdJson, _ := json.Marshal(addressesId)
	return "{\"id\" :" + string(addressesIdJson) + "}", nil
}
func (imp *DoUserImp) DeleteAddressById(id string) (string, error) {
	session, sessionErr := db.Init()
	if sessionErr != nil {
		fmt.Println(sessionErr)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("users").C("addresses")

	deleteError := c.RemoveId(bson.ObjectIdHex(id))
	if deleteError != nil {
		return deleteError.Error(), nil
	}
	return "{\"status\": true}", nil
}
func (imp *DoUserImp) FindAddressById(id string) (string, error) {
	session, sessionErr := db.Init()
	fmt.Sprintln(session)
	if sessionErr != nil {
		fmt.Println(sessionErr)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("users").C("addresses")

	var card AddressDetail
	findCustomerError := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&card)
	if findCustomerError != nil {
		return findCustomerError.Error(), nil
	}
	cardjson, _ := json.Marshal(card)
	return string(cardjson), nil
}

