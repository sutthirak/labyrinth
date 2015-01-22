package data

import (
	"log"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type Scene struct {
	Id bson.ObjectId `bson:"_id"`
	Name string
    RefName string
    MainCamera string
    Objects []Object
    Cameras []Camera
    Lights []Light
    Sprites []Sprite
    Images []Image
    Texts []Text
    Animations []Animation
    RenderScript string
    Actions map[string]string
}

type Action struct {
    Target string
    ActionScript string
}

func GetSceneList() []Scene{
	session, _ := mgo.Dial("localhost")
    
    defer session.Close()
    connection := session.DB("labyrinth").C("scene")
    scenes := make([]Scene,0)
    scene := Scene{}
    iter := connection.Find(nil).Skip(0).Iter()
    for iter.Next(&scene) {
        scenes = append(scenes,scene)
    }

    defer iter.Close();

    return scenes
}

func GetSceneByHex(hex string) Scene{
	objectId := bson.ObjectIdHex(hex)
	if objectId.Valid() {
		session, err := mgo.Dial("localhost")
        if err != nil {
            panic(err)
        }
        
        defer session.Close()
        connection := session.DB("labyrinth").C("scene")
       	scene := Scene{}
        connection.FindId(objectId).One(&scene)
        return scene        
	}
	return Scene{}	
}

func SaveScene(scene *Scene) *Scene {

	id := scene.Id
    if !id.Valid() {
        scene.Id = bson.NewObjectId()
    }

	session, err := mgo.Dial("localhost")
    
    defer session.Close()

    connection := session.DB("labyrinth").C("scene")
    _ , err = connection.UpsertId(scene.Id,scene)

    if err != nil {
    	log.Fatal(err)
    }

    return scene
}

func DeleteScene(hex string) string{
    
    message := ""
    objectId := bson.ObjectIdHex(hex)

    if objectId.Valid() {

        session, err := mgo.Dial("localhost")
        if err != nil {
            message = err.Error()
        }
    
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        
        connection := session.DB("labyrinth").C("scene")
        err = connection.RemoveId(objectId)
        message = "success"

        if err != nil {
            log.Fatal(err)
            message = err.Error()
        }
    }

    return message

}