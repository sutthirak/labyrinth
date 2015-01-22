package renderer

import (
	"bytes"
	//"strconv"
	"github.com/sutthirak/labyrinth/data"
)

func RenderObject(object data.Object) string{

	var buffer bytes.Buffer

	if object.RefName != "" {
		name := object.RefName
		buffer.WriteString("var "+name+";\n")

		//check kind of object
		if object.Kind == "Box" {
			buffer.WriteString(name+" = BABYLON.Mesh.CreateBox(\""+name+"\","+ object.Properties["size"] +", scene);\n")
		}else if object.Kind == "Sphere"{
			buffer.WriteString(name+"= BABYLON.Mesh.CreateSphere(\""+name+"\","+object.Properties["segment"]+","+object.Properties["diameter"]+", scene);\n")
		}else if object.Kind == "Plane"{
			buffer.WriteString(name+" = BABYLON.Mesh.CreatePlane(\""+name+"\","+object.Properties["size"]+", scene);\n")
		}else if object.Kind == "Babylon"{
			buffer.WriteString("BABYLON.SceneLoader.ImportMesh(\""+object.Properties["mesh"]+"\",\"http://localhost:9000/public/asset/babylon/\",\""+object.Properties["file"]+"\", scene, function (newMeshes, particleSystems) {\n")
		        buffer.WriteString("for(var index in newMeshes){\n")
		        	buffer.WriteString("if(newMeshes[index].name == \""+name+"\"){\n")
			            buffer.WriteString(name +" = newMeshes[0];\n")
			            buffer.WriteString(name+".position.x = "+object.Properties["x"]+";\n")
			            buffer.WriteString(name+".position.y = "+object.Properties["y"]+";\n")
			            buffer.WriteString(name+".position.z = "+object.Properties["z"]+";\n")
			        buffer.WriteString("}\n")
		        buffer.WriteString("}\n")
	        buffer.WriteString("});\n")
		}

		//set object position
		if object.Kind != "" && object.Kind != "Babylon" {
			buffer.WriteString(name+".position = new BABYLON.Vector3("+object.Properties["x"]+","+object.Properties["y"]+","+object.Properties["z"]+");\n") 
		}

		//set material
		if object.Material != nil{
			materialObjectName := "material_"+object.RefName
			buffer.WriteString("var "+materialObjectName+" = new BABYLON.StandardMaterial(\""+materialObjectName+"\", scene)\n");
			
			if object.Material.Color != nil{
				buffer.WriteString(materialObjectName+".diffuseColor = new BABYLON.Color3("+object.Material.Color.Red+","+object.Material.Color.Green+","+object.Material.Color.Blue+");\n")
			}

			if object.Material.Alpha != "" {			
				buffer.WriteString(materialObjectName+".alpha = "+object.Material.Alpha+";\n")
			}

			if object.Material.Texture != "" {
				buffer.WriteString(materialObjectName+".diffuseTexture = new BABYLON.Texture(\"/public/asset/material/"+object.Material.Texture+"\", scene);\n")
			}

			buffer.WriteString(name+".material = "+materialObjectName+";\n")
		}
	}

	return buffer.String()
}