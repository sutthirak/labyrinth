package controllers

import (
	"io/ioutil"
	"strings"
	"strconv"
	"github.com/revel/revel"
	"github.com/sutthirak/labyrinth/data"
	"github.com/sutthirak/labyrinth/util"
	"github.com/sutthirak/labyrinth/renderer"
)

type SceneController struct {
	*revel.Controller
}

func (c SceneController) Scenes() revel.Result {
	scenes := data.GetSceneList()
	return c.Render(scenes)
}

func (c SceneController) Scene(hex string) revel.Result {
	scene := data.GetSceneByHex(hex)
	objects := data.GetObjectList()
	sprites := data.GetSpriteList()
	imageDirectory , _ := util.GetResourceImagePath()
	images , _ := ioutil.ReadDir(imageDirectory)
	return c.Render(scene,objects,sprites,images)
}

func (c SceneController) Preview(scene string) revel.Result {
	script := renderer.RenderSceneAndSetActive(data.GetSceneList(),scene,true,true)
	return c.Render(script)
}

func (c SceneController) AddScene(sceneName string) revel.Result {
	
	c.Validation.Required(sceneName).Message("Please enter your scene name")
	if c.Validation.HasErrors(){
        c.Validation.Keep()
        return c.Redirect("/admin/scenes")
    }

	scene := new(data.Scene)
	scene.Name = strings.Replace(sceneName," ","_",-1)
	scene.RefName = strings.Replace(sceneName," ","_",-1)

	//add default arc camera
	camera := data.Camera{}
	camera.Name = "default"
	camera.RefName = "default_0"
	camera.Kind = "FREE_CAMERA"

	camera.Properties = make(map[string]string)
	camera.Properties["x"] = "0"
	camera.Properties["y"] = "5"
	camera.Properties["z"] = "-1"
	camera.Properties["target"] = "new BABYLON.Vector3(0,0,0)"
	scene.Cameras = append(scene.Cameras,camera)

	scene.MainCamera = "default_0"
	data.SaveScene(scene)
	c.Flash.Success("Done!")
	return c.Redirect("/admin/scenes")
}

func (c SceneController) DeleteScene(hex string) revel.Result {

    data.DeleteScene(hex)
    message := "success"
    return c.RenderText(message)
}

func (c SceneController) AddObject(sceneHex string,
								   objectHex string) revel.Result {

	scene := data.GetSceneByHex(sceneHex)
	object := data.GetObjectByHex(objectHex)
	object.RefName = object.Name +"_"+ strconv.Itoa(len(scene.Objects))
	properties := object.Properties
	properties["x"] = "0"
	properties["y"] = "0"
	properties["z"] = "0"

	scene.Objects = append(scene.Objects,object)
	data.SaveScene(&scene)
	c.Flash.Success("Done!")
	return c.Redirect("/admin/scenes/%s?tab=object",scene.Id.Hex())
}

func (c SceneController) AddSprite(sceneHex string,
								   spriteHex string) revel.Result {

	scene := data.GetSceneByHex(sceneHex)
	sprite := data.GetSpriteByHex(spriteHex)
	sprite.RefName = sprite.Name +"_"+ strconv.Itoa(len(scene.Sprites))
	properties := sprite.Properties
	properties["x"] = "0"
	properties["y"] = "0"
	properties["z"] = "0"

	scene.Sprites = append(scene.Sprites,sprite)
	data.SaveScene(&scene)

	return c.Redirect("/admin/scenes/%s",scene.Id.Hex())
}

func (c SceneController) AddCamera(	sceneHex string,
									name string,
									kind string,
									x string,
									y string,
									z string,
									alpha string,
									beta string,
									redius string,
									target string) revel.Result{

	c.Validation.Required(name).Message("Please enter your camera name")
	c.Validation.Required(kind).Message("Please enter your camera kind")
	if kind == "FREE_CAMERA" {
		if !util.IsInt(x) {
			c.Validation.Error("Plase enter correct x value")
		}
		if !util.IsInt(y) {
			c.Validation.Error("Plase enter correct y value")
		}
		if !util.IsInt(x) {
			c.Validation.Error("Plase enter correct z value")
		}
	}
	if kind == "ARC_CAMERA" {
		if !util.IsFloat(alpha) {
			c.Validation.Error("Plase enter correct alpha value")
		}
		if !util.IsFloat(beta) {
			c.Validation.Error("Plase enter correct beta value")
		}
		if !util.IsFloat(redius) {
			c.Validation.Error("Plase enter correct redius value")
		}
	}

	if c.Validation.HasErrors(){
        c.Validation.Keep()
        return c.Redirect("/admin/scenes/%s?tab=camera",sceneHex)
    }


	scene := data.GetSceneByHex(sceneHex)
	camera := data.Camera{}
	camera.Name = strings.Replace(name," ","_",-1)
	camera.RefName = strings.Replace(name," ","_",-1) + "_" + strconv.Itoa(len(scene.Cameras))
	camera.Kind = kind
	camera.Properties = make(map[string]string)

	if kind == "FREE_CAMERA" {
		camera.Properties["x"] = x
		camera.Properties["y"] = y
		camera.Properties["z"] = z		
	}

	if kind == "ARC_CAMERA" {
		camera.Properties["alpha"] = alpha
		camera.Properties["beta"] = beta
		camera.Properties["redius"] = redius
	}

	if target != "" {
		camera.Properties["target"] = target
	}else{
		camera.Properties["target"] = "new BABYLON.Vector3(0,0,0)"
	}

	scene.Cameras = append(scene.Cameras,camera)
	data.SaveScene(&scene)
	c.Flash.Success("Done!")
	return c.Redirect("/admin/scenes/%s?tab=camera",scene.Id.Hex())
}

func (c SceneController) SetMainCamera(sceneHex string,
										cameraRefName string) revel.Result {

	scene := data.GetSceneByHex(sceneHex)
	scene.MainCamera = cameraRefName
	data.SaveScene(&scene)
	c.Flash.Success("Done!")
	return c.RenderText("success")	
}

func (c SceneController) UpdateObject(sceneHex string,
								   elementNumber int,
								   x string,
								   y string,
								   z string,
								   size string) revel.Result {

	c.Validation.Required(x).Message("Please enter x postion")
	c.Validation.Required(y).Message("Please enter y postion")
	c.Validation.Required(z).Message("Please enter z postion")
	c.Validation.Required(size).Message("Please enter size")
	if c.Validation.HasErrors(){
        c.Validation.Keep()
        return c.Redirect("/admin/scenes/%s?tab=object",sceneHex)
    }

	scene := data.GetSceneByHex(sceneHex)
	object := scene.Objects[elementNumber]

	properties := object.Properties
	properties["x"] = x
	properties["y"] = y
	properties["z"] = z
	properties["size"] = size

	s := data.SaveScene(&scene)
	c.Flash.Success("Updated!")
	return c.Redirect("/admin/scenes/%s?tab=object",s.Id.Hex())
}

func (c SceneController) UpdateSprite(sceneHex string,
								   elementNumber int,
								   x string,
								   y string,
								   z string,
								   size string) revel.Result {

	scene := data.GetSceneByHex(sceneHex)
	sprite := scene.Sprites[elementNumber]

	properties := sprite.Properties
	
	if x != "" {
		properties["x"] = x
	}

	if y != "" {
		properties["y"] = y
	}

	if z != "" {
		properties["z"] = z
	}

	if size != "" {
		properties["size"] = size
	}

	s := data.SaveScene(&scene)
	return c.Redirect("/admin/scenes/%s",s.Id.Hex())
}

func (c SceneController) UpdateCamera(sceneHex string,
										elementNumber int,
										x string,
										y string,
										z string,
										alpha string,
										beta string,
										redius string,
										target string) revel.Result {

	scene := data.GetSceneByHex(sceneHex)
	camera := scene.Cameras[elementNumber]
	if camera.Kind == "FREE_CAMERA" {
		camera.Properties["x"] = x
		camera.Properties["y"] = y
		camera.Properties["z"] = z		
	}

	if camera.Kind == "ARC_CAMERA" {
		camera.Properties["alpha"] = alpha
		camera.Properties["beta"] = beta
		camera.Properties["redius"] = redius
	}

	data.SaveScene(&scene)
	c.Flash.Success("Updated!")
	return c.Redirect("/admin/scenes/%s",scene.Id.Hex())
}

func (c SceneController) DeleteObject(	sceneHex string,
										elementNumber int) revel.Result {

	scene := data.GetSceneByHex(sceneHex)
	objs := make([]data.Object,0)

	for index , element := range scene.Objects {
		if index != elementNumber {
			objs = append(objs,element)
		}
	}

	scene.Objects = objs
	data.SaveScene(&scene)
	c.Flash.Success("Deleted!")

	return c.RenderText("success")
}

func (c SceneController) DeleteSprite(	sceneHex string,
										elementNumber int) revel.Result {

	scene := data.GetSceneByHex(sceneHex)
	sprites := make([]data.Sprite,0)

	for index , element := range scene.Sprites {
		if index != elementNumber {
			sprites = append(sprites,element)
		}
	}

	scene.Sprites = sprites
	data.SaveScene(&scene)
	return c.RenderText("success")

}

func (c SceneController) DeleteCamera(	sceneHex string,
										cameraRefName string) revel.Result {

	scene := data.GetSceneByHex(sceneHex)
	cameras := make([]data.Camera,0)

	for _ , element := range scene.Cameras {
		if element.RefName != cameraRefName || element.Name == "default" {
			cameras = append(cameras,element)
		}
	}

	//if camera list has only 1 child, set default camera
	if len(cameras) == 1 {
		scene.MainCamera = "default_0"
	}

	scene.Cameras = cameras
	data.SaveScene(&scene)
	c.Flash.Success("Deleted!")
	return c.RenderText("success")

}

func (c SceneController) Script(hex string) revel.Result {
	scenes := data.GetSceneList()
	scene := data.GetSceneByHex(hex)
	return c.Render(scenes,scene)
}

func (c SceneController) ScriptSave(hex string) revel.Result {
	scene := data.GetSceneByHex(hex)
	bytes, _ := ioutil.ReadAll(c.Request.Body)
	scene.RenderScript = string(bytes)
	data.SaveScene(&scene)
	return c.RenderText("success")
}

func (c SceneController) Action(sceneHex string,objectRefName string) revel.Result {
	scenes := data.GetSceneList()
	scene := data.GetSceneByHex(sceneHex)
	action := scene.Actions[objectRefName]
	if action != "" {
		c.Render(scenes,scene,objectRefName,action)	
	}
	return c.Render(scenes,scene,objectRefName)
}

func (c SceneController) ActionSave(sceneHex string,objectRefName string) revel.Result {
	scene := data.GetSceneByHex(sceneHex)
	bytes, _ := ioutil.ReadAll(c.Request.Body)
	action := string(bytes)
	scene.Actions[objectRefName] = action
	data.SaveScene(&scene)
	c.Flash.Success("Done!")
	return c.Redirect("/admin/scenes/%s/objects/%s/action",sceneHex,objectRefName)
}