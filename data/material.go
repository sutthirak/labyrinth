package data

import (
	"log"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type Color struct{
    Red string
    Green string
    Blue string
}

type Material struct {
	Id bson.ObjectId `bson:"_id"`
	Name string
    Color *Color
    Texture string
    Alpha string
}

func GetMaterialList() []Material{
	session, err := mgo.Dial("localhost")
    if err != nil {
        panic(err)
    }
    
    defer session.Close()
    connection := session.DB("labyrinth").C("material")
    materials := make([]Material,0)
    material := Material{}
    iter := connection.Find(nil).Skip(0).Iter()
    for iter.Next(&material) {
        materials = append(materials,material)
    }

    if err := iter.Close(); err != nil {
        panic("Error")
    }

    return materials
}

func GetMaterialByHex(hex string) Material{
	objectId := bson.ObjectIdHex(hex)
	if objectId.Valid() {
		session, err := mgo.Dial("localhost")
        if err != nil {
            panic(err)
        }
        
        defer session.Close()
        connection := session.DB("labyrinth").C("material")
       	material := Material{}
        connection.FindId(objectId).One(&material)
        return material        
	}
	return Material{}	
}

func SaveMaterial(material *Material) *Material {

	id := material.Id
    if !id.Valid() {
        material.Id = bson.NewObjectId()
    }

	session, err := mgo.Dial("localhost")
    if err != nil {
        panic(err)
    }
    
    defer session.Close()

    connection := session.DB("labyrinth").C("material")
    _ , err = connection.UpsertId(material.Id,material)

    if err != nil {
    	log.Fatal(err)
    }

    return material
}

func DeleteMaterial(hex string) string{
    
    message := ""
    material := GetMaterialByHex(hex)

    if material.Id.Valid() {
        session, err := mgo.Dial("localhost")
        if err != nil {
            panic(err)
            message = err.Error()
        }
    
        defer session.Close()

        connection := session.DB("labyrinth").C("material")
        err = connection.Remove(material)
        message = "success"

        if err != nil {
            log.Fatal(err)
            message = err.Error()

        }
    }

    return message

}