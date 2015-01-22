package renderer

import (
	"bytes"
	"github.com/sutthirak/labyrinth/data"
)

func RenderAction(scenes []data.Scene) string{

	var buffer bytes.Buffer

	buffer.WriteString("window.addEventListener(\"click\", function () {\n")

		buffer.WriteString("var pickResult = labyrinth.pick(labyrinth.pointerX, labyrinth.pointerY);\n")
		buffer.WriteString("if (pickResult.hit) {\n")

			for _ , scene := range scenes {
				for key , value := range scene.Actions {
					buffer.WriteString("if(state === \""+scene.RefName+"\" && pickResult.pickedMesh.id === \""+key+"\"){\n")
					buffer.WriteString(value+"\n")
					buffer.WriteString("}\n")	
				}
			}

			// buffer.WriteString("console.log(pickResult.pickedMesh);")
			// buffer.WriteString("console.log(pickResult.pickedMesh.name);")
			// buffer.WriteString("console.log(pickResult.pickedMesh.id);")

		buffer.WriteString("}")
	buffer.WriteString("});")

	return buffer.String()
}