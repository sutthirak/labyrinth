package data

import (
	"log"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type Sprite struct {
	Id bson.ObjectId `bson:"_id"`
	Name string
	RefName string
	AssetName string
	Properties map[string]string
}

func GetSpriteList() []Sprite{

	session, err := mgo.Dial("localhost")
    if err != nil {
        panic(err)
    }
    
    defer session.Close()
    connection := session.DB("labyrinth").C("sprite")

    sprites := make([]Sprite,0)
    sprite := Sprite{}
    iter := connection.Find(nil).Skip(0).Iter()
    for iter.Next(&sprite) {
        sprites = append(sprites,sprite)
    }

    if err := iter.Close(); err != nil {
        panic("Error")
    }

    return sprites
}

func GetSpriteByHex(hex string) Sprite{
	objectId := bson.ObjectIdHex(hex)
	if objectId.Valid() {
		session, err := mgo.Dial("localhost")
        if err != nil {
            panic(err)
        }
        
        defer session.Close()
        connection := session.DB("labyrinth").C("sprite")
       	sprite := Sprite{}
        connection.FindId(objectId).One(&sprite)
        return sprite        
	}
	return Sprite{}	
}

func SaveSprite(sprite *Sprite) *Sprite {

    id := sprite.Id

    if !id.Valid() {
        sprite.Id = bson.NewObjectId()
    }

	session, err := mgo.Dial("localhost")
    if err != nil {
        panic(err)
    }
    
    defer session.Close()

    connection := session.DB("labyrinth").C("sprite")
    _ , err = connection.UpsertId(sprite.Id,sprite)

    if err != nil {
         log.Fatal(err)
    }

    return sprite
}

func DeleteSprite(hex string) string{
    
    message := ""
    sprite := GetSpriteByHex(hex)

    if sprite.Id.Valid() {
        session, err := mgo.Dial("localhost")
        if err != nil {
            panic(err)
            message = err.Error()
        }
    
        defer session.Close()

        connection := session.DB("labyrinth").C("sprite")
        err = connection.Remove(sprite)
        message = "success"

        if err != nil {
            log.Fatal(err)
            message = err.Error()

        }
    }

    return message

}