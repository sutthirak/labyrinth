package controllers

import (
	"strconv"
	"github.com/revel/revel"
    "github.com/sutthirak/labyrinth/data"
    "github.com/sutthirak/labyrinth/util"
    "gopkg.in/mgo.v2/bson"
)

type SpriteController struct {
	*revel.Controller
}

func (c SpriteController) Sprites() revel.Result {
    sprites := data.GetSpriteList()
	return c.Render(sprites)
}

func (c SpriteController) Sprite(hex string) revel.Result {

    resources := util.GetResourceImage()

    if hex != "" {
        sprite := data.GetSpriteByHex(hex)
        return c.Render(sprite,resources)
    }

    return c.Render(resources)
}

func (c SpriteController) SpriteSave(
                            hex string,
                            name string,
                            asset string,
                            instance int,
                            size int) revel.Result {
    
    sprite := new (data.Sprite)
    
    if hex != "" && bson.ObjectIdHex(hex).Valid() {
        sprite.Id = bson.ObjectIdHex(hex)
    }

    sprite.Name = name
    sprite.AssetName = asset

    if sprite.Properties == nil {
    	sprite.Properties = make(map[string]string)
    }

	sprite.Properties["size"] = strconv.Itoa(size)
	sprite.Properties["instance"] = strconv.Itoa(instance)

    sprite = data.SaveSprite(sprite)
	return c.Redirect("/admin/sprites/%s",sprite.Id.Hex())
}

func (c SpriteController) SpriteDelete(hex string) revel.Result {
    revel.INFO.Printf("delete with sprite id = %s",hex)
    message := data.DeleteSprite(hex)
    return c.RenderText(message)
}