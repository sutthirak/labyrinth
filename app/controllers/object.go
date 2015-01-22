package controllers

import (
	"github.com/revel/revel"
    "github.com/sutthirak/labyrinth/data"
    "github.com/sutthirak/labyrinth/util"
    "github.com/sutthirak/labyrinth/renderer"
    "gopkg.in/mgo.v2/bson"
)

type ObjectController struct {
	*revel.Controller
}

func (c ObjectController) Objects() revel.Result {
    objects := data.GetObjectList()
	return c.Render(objects)
}

func (c ObjectController) Object(hex string) revel.Result {

    materials := data.GetMaterialList()
    babylons := util.GetResourceBabylon()

    if hex != "" {
        object := data.GetObjectByHex(hex)

        if object.Material == nil {
            material := data.Material{}
            material.Color = &data.Color{}
            object.Material = &material 
        }
        preview := renderer.PreviewObject(object)
        return c.Render(object,materials,babylons,preview)
    }

    object := data.Object{}
    preview := renderer.PreviewObject(object)
    return c.Render(materials,babylons,preview)
}

func (c ObjectController) ObjectSave(
                            hex string,
                            name string,
                            kind string,
                            size string,
                            segment string,
                            diameter string,
                            materialHex string,
                            file string,
                            mesh string) revel.Result {
    
    c.Validation.Required(kind).Message("Please enter kind of an object")
    c.Validation.Required(name).Message("Please enter name")

    if kind == "Box" || kind == "Plane" {
        if !util.IsInt(size) {
            c.Validation.Error("Plase enter correct size value")
        }
    }

    if kind == "Sphere" {
        if !util.IsInt(segment) {
            c.Validation.Error("Plase enter correct segment value")
        }
        if !util.IsInt(diameter) {
            c.Validation.Error("Plase enter correct diameter value")
        }
    }

    if kind == "Babylon" {
        c.Validation.Required(mesh).Message("Please enter mesh value")
        c.Validation.Required(file).Message("Please choose babylon file")
    }

    if c.Validation.HasErrors(){
        c.Validation.Keep()
        c.FlashParams()
        if hex == ""{
            c.Redirect("/admin/objects")
        }
        return c.Redirect("/admin/objects/%s",hex)
    }

    obj := new (data.Object)
    if hex != "" && bson.ObjectIdHex(hex).Valid() {
        obj.Id = bson.ObjectIdHex(hex)
    }

    obj.Name = name
    obj.Kind = kind

    var properties map[string]string
    properties = make(map[string]string)

    if kind == "Box" || kind == "Plane" {
        properties["size"] = size
    }

    if kind == "Sphere" {
        properties["segment"] = segment   
        properties["diameter"] = diameter   
    }

    if kind == "Babylon" {
        properties["file"] = file   
        properties["mesh"] = mesh   
    }     

    obj.Properties = properties

    material := data.Material{}
    if materialHex != "" && bson.ObjectIdHex(materialHex).Valid() {
        material = data.GetMaterialByHex(materialHex)
    }

    obj.Material = &material
    obj = data.SaveObject(obj)
    c.Flash.Success("Done !")
	return c.Redirect("/admin/objects/%s",obj.Id.Hex())
}

func (c ObjectController) ObjectDelete(hex string) revel.Result {
    revel.INFO.Printf("delete with object id = %s",hex)
    message := data.DeleteObject(hex)
    return c.RenderText(message)
}