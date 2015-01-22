package data

import (
	"log"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type Object struct {
    Id bson.ObjectId `bson:"_id"`
	Name string
    RefName string
	Kind string
    Properties map[string]string
    Material *Material
}

func GetObjectList() []Object{

	session, err := mgo.Dial("localhost")
    if err != nil {
        panic(err)
    }
    
    defer session.Close()
    connection := session.DB("labyrinth").C("object")

    objects := make([]Object,0)
    object := Object{}
    iter := connection.Find(nil).Skip(0).Iter()
    for iter.Next(&object) {
        objects = append(objects,object)
    }

    if err := iter.Close(); err != nil {
        panic("Error")
    }

    return objects
}

func GetObjectByHex(hex string) Object{
    objectId := bson.ObjectIdHex(hex)
    if objectId.Valid() {
        session, err := mgo.Dial("localhost")
        if err != nil {
            panic(err)
        }
        
        defer session.Close()
        connection := session.DB("labyrinth").C("object")
        object := Object{}
        connection.FindId(objectId).One(&object)
        return object        
    }
    return Object{}  
}

func SaveObject(obj *Object) *Object {

    id := obj.Id

    if !id.Valid() {
        obj.Id = bson.NewObjectId()
    }

	session, err := mgo.Dial("localhost")
    if err != nil {
        panic(err)
    }
    
    defer session.Close()

    connection := session.DB("labyrinth").C("object")
    _ , err = connection.UpsertId(obj.Id,obj)

    if err != nil {
         log.Fatal(err)
    }

    return obj
}

func DeleteObject(hex string) string {

    message := ""
    objectId := bson.ObjectIdHex(hex)

    if objectId.Valid() {
        session, err := mgo.Dial("localhost")
        if err != nil {
            panic(err)
            message = err.Error()
        }
    
        defer session.Close()

        connection := session.DB("labyrinth").C("object")
        err = connection.RemoveId(objectId)
        message = "success"

        if err != nil {
            log.Fatal(err)
            message = err.Error()

        }
    }

    return message
}