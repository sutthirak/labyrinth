package controllers

import (
	"log"
	"strings"
	"strconv"
	"github.com/revel/revel"
	"github.com/sutthirak/labyrinth/data"
	"github.com/sutthirak/labyrinth/util"
)

type AnimationController struct {
	*revel.Controller
}

func (c AnimationController) Animations(sceneHex string) revel.Result {
	scene := data.GetSceneByHex(sceneHex)
	return c.Render(scene)
}

func (c AnimationController) AnimationsSave(sceneHex string,
											name string,
											target string,
											kind string,
											animationtype string) revel.Result {
	
	scene := data.GetSceneByHex(sceneHex)

	c.Validation.Required(name).Message("Please enter your animation name")
	if c.Validation.HasErrors(){
        c.Validation.Keep()
        return c.Redirect("/admin/scenes/%s/animations",sceneHex)
    }

	animation := data.Animation{}
	animation.Name = strings.Replace(name," ","_",-1)
	animation.RefName = strings.Replace(name," ","_",-1) +"_"+ strconv.Itoa(len(scene.Animations))
	animation.Target = target
	animation.Kind = kind
	animation.Type = animationtype
	animation.FrameRate = "60"
	animation.FrameValueType = "BABYLON.Animation.ANIMATIONTYPE_FLOAT"
	scene.Animations = append(scene.Animations,animation)
	s := data.SaveScene(&scene)


	return c.Redirect("/admin/scenes/%s/animations",s.Id.Hex())
}

func (c AnimationController) AnimationsDelete(sceneHex string,
											elementNumber int) revel.Result{

	scene := data.GetSceneByHex(sceneHex)
	animations := make([]data.Animation,0)

	for index , element := range scene.Animations {
		if index != elementNumber {
			animations = append(animations,element)
		}
	}

	scene.Animations = animations
	data.SaveScene(&scene)
	return c.RenderText("success")

}

func (c AnimationController) AnimationFrame(sceneHex string,
											elementNumber int) revel.Result{

	scene := data.GetSceneByHex(sceneHex)
	object := data.Object{}
	animation := scene.Animations[elementNumber]
	for _ , o := range scene.Objects {
		if o.RefName == animation.Target {
			object = o
		}
	}
	return c.Render(scene,object,elementNumber,animation)
}

func (c AnimationController) AnimationFrameSave(sceneHex string,
											increment string,
											elementNumber int,
											frame string,
											value string) revel.Result{

	if !util.IsInt(frame) {
		c.Validation.Error("Plase enter correct frame value")
	}
	if !util.IsFloat(value) {
		c.Validation.Error("Plase enter correct value")
	}

	if c.Validation.HasErrors(){
        c.Validation.Keep()
        return c.Redirect("/admin/scenes/%s/animations/%b/frames",sceneHex,elementNumber)
    }

	scene := data.GetSceneByHex(sceneHex)
	animation := scene.Animations[elementNumber]
	animationFrame := data.AnimationFrame{}
	animationFrame.Frame = frame
	animationFrame.Value = value
	animation.Frames = append(animation.Frames,animationFrame)
	scene.Animations[elementNumber] = animation
	s := data.SaveScene(&scene)
	return c.Redirect("/admin/scenes/%s/animations/%b/frames",s.Id.Hex(),elementNumber)
}
		
func (c AnimationController) AnimationsFrameDelete(sceneHex string,
											elementNumber int,
											frameNumber int) revel.Result{

	log.Println("frame index = ",frameNumber)
	scene := data.GetSceneByHex(sceneHex)
	animation := scene.Animations[elementNumber]

	frames := make([]data.AnimationFrame,0)

	for index , frame := range animation.Frames {
		log.Println("index = ",index)
		if index != frameNumber {
			log.Println("not equal")
			frames = append(frames,frame)
		}
	}

	scene.Animations[elementNumber].Frames = frames
	data.SaveScene(&scene)
	return c.RenderText("success")

}
