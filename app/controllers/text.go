package controllers

import (
	"io/ioutil"
	"strings"
	"strconv"
	"github.com/revel/revel"
	"github.com/sutthirak/labyrinth/data"
)

type TextController struct {
	*revel.Controller
}

func (c TextController) TextSave(sceneHex string,
								name string,
								content string,
								size string,
								x string,
								y string,
								height string,
								width string,
								color string,
								stroke string) revel.Result {

	scene := data.GetSceneByHex(sceneHex)

	text := data.Text{}
	text.Name = strings.Replace(name," ","_",-1)
	text.RefName = strings.Replace(name," ","_",-1) +"_"+ strconv.Itoa(len(scene.Texts))
	properties := make(map[string]string)
	properties["content"] = content
	properties["x"] = x
	properties["y"] = y
	properties["height"] = height
	properties["width"] = width
	properties["size"] = size
	properties["color"] = color
	properties["stroke"] = stroke
	text.Properties = properties

	scene.Texts = append(scene.Texts,text)
	s := data.SaveScene(&scene)

	return c.Redirect("/admin/scenes/%s?tab=text",s.Id.Hex())

}

func (c TextController) TextUpdate(sceneHex string,
									elementNumber int,
									content string,
									size string,
									x string,
									y string,
									height string,
									width string,
									color string,
									stroke string) revel.Result {
	scene := data.GetSceneByHex(sceneHex)
	text := scene.Texts[elementNumber]
	text.Properties["content"] = content
	text.Properties["x"] = x
	text.Properties["y"] = y
	text.Properties["height"] = height
	text.Properties["width"] = width
	text.Properties["size"] = size
	text.Properties["color"] = color
	text.Properties["stroke"] = stroke
	s := data.SaveScene(&scene)		
	c.Flash.Success("Updated!")
	return c.Redirect("/admin/scenes/%s?tab=text",s.Id.Hex())
}

func (c TextController) TextDelete(sceneHex string,
									elementNumber int) revel.Result{

	scene := data.GetSceneByHex(sceneHex)
	texts := make([]data.Text,0)

	for index , element := range scene.Texts {
		if index != elementNumber {
			texts = append(texts,element)
		}
	}

	scene.Texts = texts
	data.SaveScene(&scene)
	return c.RenderText("success")

}

func (c TextController) Action(sceneHex string,textRefName string) revel.Result {
	scenes := data.GetSceneList()
	scene := data.GetSceneByHex(sceneHex)
	for _ , text := range scene.Texts {
		if text.RefName == textRefName && text.Properties["action"] != "" {
			action := text.Properties["action"]
			c.Render(scenes,scene,textRefName,action)	
		}
	}
	return c.Render(scenes,scene,textRefName)
}

func (c TextController) ActionSave(sceneHex string,textRefName string) revel.Result {
	scene := data.GetSceneByHex(sceneHex)
	bytes, _ := ioutil.ReadAll(c.Request.Body)
	action := string(bytes)

	for _ , text := range scene.Texts {
		if text.RefName == textRefName && action != "" {
			text.Properties["action"] = action
			data.SaveScene(&scene)
		}
	}

	return c.RenderText("success")
}