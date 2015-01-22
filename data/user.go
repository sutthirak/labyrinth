package data

import (
	"log"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "golang.org/x/crypto/bcrypt"
)

type User struct {
	Id bson.ObjectId `bson:"_id"`
	Email string
	Password string
	FirstName string 
	LastName string
}

func CountUser() int {
    session, err := mgo.Dial("localhost")
    if err != nil {
        panic(err)
    }
    
    defer session.Close()
    connection := session.DB("labyrinth").C("user")
    count , err := connection.Count();

    if err != nil {
        panic(err)
    }

    return count
}

func Auth(email string,password string) bool{

    hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 12)

	session, err := mgo.Dial("localhost")
    if err != nil {
        panic(err)
    }
    
    defer session.Close()
    connection := session.DB("labyrinth").C("user")
    count , err := connection.Find(bson.M{"email":email}).Count();

    if err != nil {
    	panic(err)
    }

    if(count == 1){
        user := User{}
        connection.Find(bson.M{"email":email}).One(&user)
        err = bcrypt.CompareHashAndPassword(hashed, []byte(password))
        if err == nil {
    	   return true
        }
    }
    return false        	
}

func SaveUser(user *User) *User {

	id := user.Id
    if !id.Valid() {
        user.Id = bson.NewObjectId()
    }

	session, err := mgo.Dial("localhost")
    if err != nil {
        panic(err)
    }
    
    defer session.Close()

    connection := session.DB("labyrinth").C("user")
    _ , err = connection.UpsertId(user.Id,user)

    if err != nil {
    	log.Fatal(err)
    }

    return user
}