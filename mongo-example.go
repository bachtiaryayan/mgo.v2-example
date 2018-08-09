package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/icrowley/fake"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var db *mgo.Database

// User is an object representing the database table.
type User struct {
	Fullname         string    `json:"fullname" bson:"fullname" toml:"fullname" yaml:"fullname"`
	Username         string    `json:"username" bson:"username" toml:"username" yaml:"username"`
	Password         string    `json:"password" bson:"password" toml:"password" yaml:"password"`
	Email            string    `json:"email" bson:"email" toml:"email" yaml:"email"`
	CreatedAt        time.Time `json:"created_at" bson:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" bson:"updated_at" toml:"updated_at" yaml:"updated_at"`
	Gender           string    `json:"gender" bson:"gender,omitempty" toml:"gender" yaml:"gender,omitempty"`
	DeletedAt        time.Time `json:"deleted_at" bson:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	Phone            string    `json:"phone" bson:"phone,omitempty" toml:"phone" yaml:"phone,omitempty"`
	PhoneVerified    int8      `json:"phone_verified" bson:"phone_verified,omitempty" toml:"phone_verified" yaml:"phone_verified,omitempty"`
	EmailVerified    int8      `json:"email_verified" bson:"email_verified,omitempty" toml:"email_verified" yaml:"email_verified,omitempty"`
	ConfirmationCode string    `json:"confirmation_code" bson:"confirmation_code,omitempty" toml:"confirmation_code" yaml:"confirmation_code,omitempty"`
}

func main() {

	session, err := mgo.Dial("localhost/collection-name")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("connected!")

	db = session.DB("collection-name")

	res := []User{}

	if err := db.C("users").Find(nil).All(&res); err != nil {
		log.Fatalf("errror %v", err)
	}
	fmt.Println(res)
	// ress := User{}
	if err := db.C("users").Find(bson.M{"username": bson.RegEx{Pattern: "yay"}}).All(&res); err != nil {
		log.Fatalf("errror %v", err)
	}
	username := fake.UserName()
	err = db.C("users").Insert(
		&User{
			Fullname:      fake.FullName(),
			Username:      username,
			Phone:         fake.Phone(),
			PhoneVerified: 1,
			Email:         fake.EmailAddress(),
			CreatedAt:     time.Now(),
		})

	if err != nil {
		panic(err)
	}

	colQuerier := bson.M{"username": username}
	change := bson.M{"$set": bson.M{"phone": fake.Phone(), "timestamp": time.Now()}}
	err = db.C("users").Update(colQuerier, change)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
	err = db.C("users").Remove(bson.M{"username": username})
	if err != nil {
		fmt.Printf("remove fail %v\n", err)
		os.Exit(1)
	}

	fmt.Println(username)
	fmt.Println("done")

}
