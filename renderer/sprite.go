package renderer

import (
	"bytes"
	"github.com/sutthirak/labyrinth/data"
)

func RenderSprite(sprites []data.Sprite) string{

	var buffer bytes.Buffer

	for _ , sprite := range sprites {
		buffer.WriteString("var spriteManager_"+sprite.RefName+" = new BABYLON.SpriteManager(\""+sprite.RefName+"Manager\", \"/public/asset/"+sprite.AssetName+"\", "+sprite.Properties["instance"]+", "+sprite.Properties["size"]+", scene);\n")
		buffer.WriteString("var sprite_"+sprite.RefName+" = new BABYLON.Sprite(\""+sprite.RefName+"\", spriteManager_"+sprite.RefName+");\n")
		buffer.WriteString("sprite_"+sprite.RefName+".position.x = "+sprite.Properties["x"]+";\n")
		buffer.WriteString("sprite_"+sprite.RefName+".position.y = "+sprite.Properties["y"]+";\n")
	}

	return buffer.String()
}