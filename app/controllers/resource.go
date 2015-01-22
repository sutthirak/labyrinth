package controllers

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
	"github.com/revel/revel"
	"github.com/sutthirak/labyrinth/util"
)

type ResourceController struct {
	*revel.Controller
}

func (c ResourceController) Resources() revel.Result {

	materialDirectory , _ := util.GetResourceMaterialPath()
	materials , _ := ioutil.ReadDir(materialDirectory)

	babylonDirectory , _ := util.GetResourceBabylonPath()
	babylons , _ := ioutil.ReadDir(babylonDirectory)

	imageDirectory , _ := util.GetResourceImagePath()
	images , _ := ioutil.ReadDir(imageDirectory)

	return c.Render(materials,babylons,images)	

}

func (c ResourceController) Resource(data io.Reader,
									kind string) revel.Result {
	
	headers, _ := c.Params.Files["data"]
	
	revel.INFO.Printf("upload file name : %s",strings.ToLower(headers[0].Filename))
	bytes , _ := ioutil.ReadAll(data)
	fileName := strings.ToLower(headers[0].Filename)

	if kind == "Material" {
 		resource , found := util.GetResourceMaterialPath()
		if found {		
			ioutil.WriteFile(resource + fileName,bytes,0666)
			revel.INFO.Printf("Saved file to folder %s ", resource)
			c.Flash.Success("Done !")
		}
	}

	if kind == "Babylon" {
 		resource , found := util.GetResourceBabylonPath()
		if found {		
			ioutil.WriteFile(resource + fileName,bytes,0666)
			revel.INFO.Printf("Saved file to folder %s ", resource)
			c.Flash.Success("Done !")
		}
	}

	if kind == "Image" {
		resource , found := util.GetResourceImagePath()
		if found {		
			ioutil.WriteFile(resource + fileName,bytes,0666)
			revel.INFO.Printf("Saved file to folder %s ", resource)
			c.Flash.Success("Done !")
		}
	}

	return c.Redirect("/admin/resources/")
}

func (c ResourceController) ResourceDelete(kind string,
											fileName string) revel.Result {

	message := "file not found"
	revel.INFO.Printf("remove file name = %s",fileName)

	if kind == "Material" {
		resource , found := util.GetResourceMaterialPath()
		if found {
			error := os.Remove(resource + fileName)
			if error == nil {
				message = "success"
			} else {
				message = error.Error()
				revel.INFO.Printf(error.Error())
			}
		}
	}

	if kind == "Babylon" {
		resource , found := util.GetResourceBabylonPath()
		if found {
			error := os.Remove(resource + fileName)
			if error == nil {
				message = "success"
			} else {
				message = error.Error()
				revel.INFO.Printf(error.Error())
			}
		}
	}

	if kind == "Image" {
		resource , found := util.GetResourceImagePath()
		if found {
			error := os.Remove(resource + fileName)
			if error == nil {
				message = "success"
			} else {
				message = error.Error()
				revel.INFO.Printf(error.Error())
			}
		}
	}

	return c.RenderText(message)

}