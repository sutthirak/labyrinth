package renderer

import (
	"bytes"
	//"strconv"
	"github.com/sutthirak/labyrinth/data"
)

func PreviewObject(object data.Object) string {

	var buffer bytes.Buffer
	buffer.WriteString("<script>\n")
	buffer.WriteString("var canvas = document.getElementById(\"previewCanvas\");\n")
	buffer.WriteString("var engine = new BABYLON.Engine(canvas, true);\n")
	buffer.WriteString("var createScene = function () {\n")
		buffer.WriteString("var scene = new BABYLON.Scene(engine);\n")
		buffer.WriteString("var camera = new BABYLON.FreeCamera(\"camera1\", new BABYLON.Vector3(0, 5, -1), scene);\n")
		buffer.WriteString("camera.setTarget(new BABYLON.Vector3.Zero());\n")
		buffer.WriteString("camera.attachControl(canvas, false);\n")
		buffer.WriteString("var light = new BABYLON.HemisphericLight(\"light1\", new BABYLON.Vector3(0, 1, 0), scene);\n")
		buffer.WriteString("light.intensity = .5;\n")
		
		//create mock properties to preview
		object.RefName = "object1"
		if object.Properties == nil {
			object.Properties = make(map[string]string)
		}
		
		object.Properties["x"] = "0"
		object.Properties["y"] = "0"
		object.Properties["z"] = "0"
		buffer.WriteString(RenderObject(object))

		buffer.WriteString("return scene;\n")
	buffer.WriteString("};\n")
	
	buffer.WriteString("var scene = createScene();\n")

	buffer.WriteString("engine.runRenderLoop(function () {\n")
		buffer.WriteString("scene.render();\n")
	buffer.WriteString("});\n")

	buffer.WriteString("window.addEventListener(\"resize\", function () {\n")
		buffer.WriteString("engine.resize();\n")
	buffer.WriteString("});\n")
	buffer.WriteString("</script>\n")
	return buffer.String()
}