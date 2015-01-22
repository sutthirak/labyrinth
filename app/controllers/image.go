package controllers

import (
	"io/ioutil"
	"strings"
	"strconv"
	"github.com/revel/revel"
	"github.com/sutthirak/labyrinth/data"
)

type ImageController struct {
	*revel.Controller
}

func (c ImageController) ImagesSave(sceneHex string,
											name string,
											source string,
											width string,
											height string,
											x string,
											y string) revel.Result {

	scene := data.GetSceneByHex(sceneHex)

	image := data.Image{}
	image.Name = strings.Replace(name," ","_",-1)
	image.RefName = strings.Replace(name," ","_",-1) +"_"+ strconv.Itoa(len(scene.Images))
	image.Source = source
	properties := make(map[string]string)
	properties["width"] = width
	properties["height"] = height
	properties["x"] = x
	properties["y"] = y
	image.Properties = properties

	scene.Images = append(scene.Images,image)
	s := data.SaveScene(&scene)

	return c.Redirect("/admin/scenes/%s?tab=image",s.Id.Hex())

}

func (c ImageController) ImagesUpdate(sceneHex string,
										elementNumber int,
										width string,
										height string,
										x string,
										y string) revel.Result {

	scene := data.GetSceneByHex(sceneHex)
	image := scene.Images[elementNumber]
	image.Properties["width"] = width
	image.Properties["height"] = height
	image.Properties["x"] = x
	image.Properties["y"] = y
	s := data.SaveScene(&scene)		
	c.Flash.Success("Updated!")
	return c.Redirect("/admin/scenes/%s?tab=image",s.Id.Hex())
}

func (c ImageController) ImagesDelete(sceneHex string,
										elementNumber int) revel.Result{

	scene := data.GetSceneByHex(sceneHex)
	images := make([]data.Image,0)

	for index , element := range scene.Images {
		if index != elementNumber {
			images = append(images,element)
		}
	}

	scene.Images = images
	data.SaveScene(&scene)
	return c.RenderText("success")

}

func (c ImageController) Action(sceneHex string,imageRefName string) revel.Result {
	scenes := data.GetSceneList()
	scene := data.GetSceneByHex(sceneHex)
	for _ , image := range scene.Images {
		if image.RefName == imageRefName && image.Properties["action"] != "" {
			action := image.Properties["action"]
			c.Render(scenes,scene,imageRefName,action)	
		}
	}
	return c.Render(scenes,scene,imageRefName)
}

func (c ImageController) ActionSave(sceneHex string,imageRefName string) revel.Result {
	scene := data.GetSceneByHex(sceneHex)
	bytes, _ := ioutil.ReadAll(c.Request.Body)
	action := string(bytes)

	for _ , image := range scene.Images {
		if image.RefName == imageRefName && action != "" {
			image.Properties["action"] = action
			data.SaveScene(&scene)
		}
	}

	return c.RenderText("success")
}