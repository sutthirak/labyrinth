package controllers

import (
	"strings"
	"strconv"
	"github.com/revel/revel"
	"github.com/sutthirak/labyrinth/data"
)

type SceneLightController struct {
	*revel.Controller
}

func (c SceneLightController) AddLight(sceneHex string,
										name string,
										kind string,
										positionX string,
										positionY string,
										positionZ string,
										directionX string,
										directionY string,
										directionZ string,
										diffuseRed string,
										diffuseGreen string,
										diffuseBlue string,
										specularRed string,
										specularGreen string,
										specularBlue string,
										angle string,
										exponent string,
										intensity string) revel.Result {
	
	c.Validation.Required(name).Message("Please enter your light name")
	c.Validation.Required(kind).Message("Please enter your light kind")

	if c.Validation.HasErrors(){
        c.Validation.Keep()
        return c.Redirect("/admin/scenes/%s?tab=light",sceneHex)
    }

    scene := data.GetSceneByHex(sceneHex)

	light := data.Light{}
	light.Name = strings.Replace(name," ","_",-1)
	light.RefName = strings.Replace(name," ","_",-1) + "_" + strconv.Itoa(len(scene.Lights))
	light.Kind = kind
	light.Properties = make(map[string]string)
	light.Properties["x"] = positionX
	light.Properties["y"] = positionY
	light.Properties["z"] = positionZ
	light.Properties["diffusered"] = diffuseRed
	light.Properties["diffusegreen"] = diffuseGreen
	light.Properties["diffuseblue"] = diffuseBlue
	light.Properties["specularred"] = specularRed
	light.Properties["speculargreen"] = specularGreen
	light.Properties["specularblue"] = specularBlue
	light.Properties["intensity"] = intensity

	if kind == "SPOT_LIGHT" {
		light.Properties["directionx"] = directionX
		light.Properties["directiony"] = directionY
		light.Properties["directionz"] = directionZ
		light.Properties["angle"] = angle
		light.Properties["exponent"] = exponent
	}

	scene.Lights = append(scene.Lights,light)
	data.SaveScene(&scene)
	c.Flash.Success("Done!")
	return c.Redirect("/admin/scenes/%s?tab=light",scene.Id.Hex())
}

func (c SceneLightController) UpdateLight(sceneHex string,
									elementNumber int,
									positionX string,
									positionY string,
									positionZ string,
									directionX string,
									directionY string,
									directionZ string,
									diffuseRed string,
									diffuseGreen string,
									diffuseBlue string,
									specularRed string,
									specularGreen string,
									specularBlue string,
									angle string,
									exponent string,
									intensity string) revel.Result {
	
	scene := data.GetSceneByHex(sceneHex)
	light := scene.Lights[elementNumber]

	light.Properties["x"] = positionX
	light.Properties["y"] = positionY
	light.Properties["z"] = positionZ
	light.Properties["diffusered"] = diffuseRed
	light.Properties["diffusegreen"] = diffuseGreen
	light.Properties["diffuseblue"] = diffuseBlue
	light.Properties["specularred"] = specularRed
	light.Properties["speculargreen"] = specularGreen
	light.Properties["specularblue"] = specularBlue
	light.Properties["intensity"] = intensity

	if light.Kind == "SPOT_LIGHT" {
		light.Properties["directionx"] = directionX
		light.Properties["directiony"] = directionY
		light.Properties["directionz"] = directionZ
		light.Properties["angle"] = angle
		light.Properties["exponent"] = exponent
	}
	
	s := data.SaveScene(&scene)		
	c.Flash.Success("Done!")
	return c.Redirect("/admin/scenes/%s?tab=light",s.Id.Hex())
}

func (c SceneLightController) DeleteLight(sceneHex string,
									elementNumber int) revel.Result{

	scene := data.GetSceneByHex(sceneHex)
	lights := make([]data.Light,0)

	for index , element := range scene.Lights {
		if index != elementNumber {
			lights = append(lights,element)
		}
	}

	scene.Lights = lights
	data.SaveScene(&scene)
	c.Flash.Success("Done!")
	return c.RenderText("success")

}
