package renderer

import (
	"bytes"
	"github.com/sutthirak/labyrinth/data"
)

func RenderLight(light data.Light) string{

	var buffer bytes.Buffer

	if light.Kind == "POINT_LIGHT" {
		buffer.WriteString("var light_"+light.RefName+" = new BABYLON.PointLight(\""+light.RefName+"\", new BABYLON.Vector3("+light.Properties["x"]+", "+light.Properties["y"]+", "+light.Properties["z"]+"), scene);\n")
	}

	if light.Kind == "SPOT_LIGHT" {
		buffer.WriteString("var light_"+light.RefName+" = new BABYLON.SpotLight("+light.RefName+", new BABYLON.Vector3("+light.Properties["x"]+","+light.Properties["y"]+","+light.Properties["z"]+"), new BABYLON.Vector3("+light.Properties["directtionx"]+","+light.Properties["directtiony"]+", "+light.Properties["directionz"]+"), "+light.Properties["angle"]+", "+light.Properties["exponent"]+", scene);\n")
	}

	buffer.WriteString("light_"+light.RefName+".diffuse = new BABYLON.Color3("+light.Properties["diffusered"]+","+light.Properties["diffuseblue"]+","+light.Properties["diffusegreen"]+");\n")
	buffer.WriteString("light_"+light.RefName+".specular = new BABYLON.Color3("+light.Properties["specularred"]+","+light.Properties["specularblue"]+","+light.Properties["speculargreen"]+");\n")
	
	buffer.WriteString("light_"+light.RefName+".intensity = "+light.Properties["intensity"]+";\n")

	return buffer.String()
}