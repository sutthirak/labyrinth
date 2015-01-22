package controllers

import (
	"github.com/revel/revel"
	"github.com/sutthirak/labyrinth/data"
    "github.com/sutthirak/labyrinth/util"
	"gopkg.in/mgo.v2/bson"
)

type MaterialController struct {
	*revel.Controller
}

func (c MaterialController) Materials() revel.Result {
	materials := data.GetMaterialList()
	return c.Render(materials)
}

func (c MaterialController) Material(hex string) revel.Result {

    resources := util.GetResourceMaterial()

    if hex != "" {
        material := data.GetMaterialByHex(hex)
        return c.Render(material,resources)
    }

    //init new value to preview
    material := new (data.Material)
    color := new (data.Color)
    color.Red = "0.10"
    color.Green = "0.50"
    color.Blue = "0.50"
    material.Color = color
    material.Alpha = "1.00"
    return c.Render(material,resources)
}

func (c MaterialController) MaterialSave(
                            hex string,
                            name string,
                            red string,
                            green string,
                            blue string,
                            texture string,
                            alpha string) revel.Result {
    
    material := new (data.Material)

    //add validation
    c.Validation.Required(name).Message("Please enter your material name")
    if !util.IsFloat(red) {
        c.Validation.Error("Plase enter correct red value")
    }

    if !util.IsFloat(green) {
        c.Validation.Error("Plase enter correct green value")
    }

    if !util.IsFloat(blue) {
        c.Validation.Error("Plase enter correct blue value")
    }

    if !util.IsFloat(alpha) {
        c.Validation.Error("Plase enter correct alpha value")
    }

    if hex != ""{
        if bson.ObjectIdHex(hex).Valid(){
            material.Id = bson.ObjectIdHex(hex)
        }else{
            c.Validation.Error("Material Id is invalid")   
        }
    }

    if c.Validation.HasErrors(){
        c.Validation.Keep()
        c.FlashParams()
        return c.Redirect("/admin/material")
    }

    material.Name = name

    color := new (data.Color)
    color.Red = red
    color.Green = green
    color.Blue = blue
    material.Color = color

    material.Texture = texture
    material.Alpha = alpha

    material = data.SaveMaterial(material)
    c.Flash.Success("Done !")
	return c.Redirect("/admin/materials/%s",material.Id.Hex())
}

func (c MaterialController) MaterialDelete(hex string) revel.Result {
    revel.INFO.Printf("delete with material id = %s",hex)
    if hex != "" && bson.ObjectIdHex(hex).Valid() {
        message := data.DeleteMaterial(hex)
        return c.RenderText(message)
    }
    return c.RenderText("hex is invalid")
}