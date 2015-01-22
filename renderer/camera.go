package renderer

import (
	"bytes"
	"github.com/sutthirak/labyrinth/data"
)

func RenderCamera(cameras []data.Camera,objects []data.Object) string{

	var buffer bytes.Buffer
	for _ , camera := range cameras {

		target := "new BABYLON.Vector3(0,0,0)"
		if camera.Properties["target"] != "new BABYLON.Vector3(0,0,0)" {
			for _ , object := range objects {
				if object.RefName == camera.Properties["target"] {
					target = "new BABYLON.Vector3("+object.Properties["x"]+","+object.Properties["y"]+","+object.Properties["z"]+")"
				}
			} 
		}

		if camera.Kind == "FREE_CAMERA" {
			buffer.WriteString("var "+camera.RefName+" = new BABYLON.FreeCamera(\""+camera.RefName+"\", new BABYLON.Vector3("+camera.Properties["x"]+","+camera.Properties["y"]+","+camera.Properties["z"]+"), scene);\n")
			buffer.WriteString(camera.RefName+".setTarget("+target+");\n")
		}

		if camera.Kind == "ARC_CAMERA" {
			buffer.WriteString("var "+camera.RefName+" = new BABYLON.ArcRotateCamera(\""+camera.RefName+"\","+camera.Properties["alpha"]+","+camera.Properties["beta"]+","+camera.Properties["redius"]+", "+target+", scene);\n");
		}

	}

	return buffer.String()
}