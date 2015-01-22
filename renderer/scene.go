package renderer

import (
	"bytes"
	"github.com/sutthirak/labyrinth/data"
)

func RenderScene(scenes []data.Scene) string{
	return RenderSceneAndSetActive(scenes, scenes[0].Id.Hex(),false,true)
}

func RenderSceneAndSetActive(scenes []data.Scene, activeSceneHex string, isControlCamera bool,is2dEnable bool) string {
	var buffer bytes.Buffer
	buffer.WriteString("<script>\n")
		buffer.WriteString("var canvas = document.getElementById(\"renderCanvas\");\n")
		buffer.WriteString("var engine = new BABYLON.Engine(canvas, true);\n")
		buffer.WriteString("var movie = {};\n")
		buffer.WriteString("var movie2d = {};\n")
		buffer.WriteString("var movie2dImage = {};\n")
		buffer.WriteString("var movie2dText = {};\n")

		for index , scene := range scenes {

			//2d function
			if is2dEnable {
				buffer.WriteString("movie2d[\""+scene.RefName+"\"] = function(){\n")

					//render 2d image
					for _ , image := range scene.Images {
						buffer.WriteString("var "+image.RefName+" = Raphael("+image.Properties["x"]+","+image.Properties["y"]+","+image.Properties["height"]+","+image.Properties["width"]+");\n")
						buffer.WriteString("var obj_"+image.RefName+" = "+image.RefName+".image(\"/public/asset/image/"+image.Source+"\","+image.Properties["x"]+","+image.Properties["y"]+","+image.Properties["width"]+","+image.Properties["height"]+");\n")	
						buffer.WriteString("obj_"+image.RefName+".click(function(){"+image.Properties["action"]+"});\n")
					}

					buffer.WriteString("movie2dImage[\""+scene.RefName+"\"] = [")
					for index , image := range scene.Images {
						if index == 0 {
							buffer.WriteString(image.RefName)
						}else{
							buffer.WriteString(","+image.RefName)
						}
					}
					buffer.WriteString("];\n")

					//render 2d text
					for _ , text := range scene.Texts {
						buffer.WriteString("var "+text.RefName+" = Raphael("+text.Properties["x"]+","+text.Properties["y"]+","+text.Properties["height"]+","+text.Properties["width"]+");\n")
						buffer.WriteString("var obj_"+text.RefName+" = "+text.RefName+".text("+text.Properties["x"]+","+text.Properties["y"]+", \""+text.Properties["content"]+"\");\n")
						buffer.WriteString("obj_"+text.RefName+".attr({ \"font-size\": "+text.Properties["size"]+", \"font-family\": \"Roboto\", \"fill\":\""+text.Properties["color"]+"\" , \"stroke\":\""+text.Properties["stroke"]+"\" });\n")
						buffer.WriteString("obj_"+text.RefName+".click(function(){"+text.Properties["action"]+"});\n")
					}

					buffer.WriteString("movie2dText[\""+scene.RefName+"\"] = [")
					for index , text := range scene.Texts {
						if index == 0 {
							buffer.WriteString(text.RefName)
						}else{
							buffer.WriteString(","+text.RefName)
						}
					}
					buffer.WriteString("];\n")

				buffer.WriteString("};\n")
			}

			buffer.WriteString("var "+scene.RefName+" = function () {\n")

				//render 2d element
				if is2dEnable && index == 0{
					buffer.WriteString("movie2d[\""+scene.RefName+"\"];\n")
				}

				//create scene
				buffer.WriteString("var scene = new BABYLON.Scene(engine);\n")

				//create light
				for _ , light := range scene.Lights {
					buffer.WriteString(RenderLight(light))
				}

				//loop for create object
				for _ , object := range scene.Objects {
					buffer.WriteString(RenderObject(object))
				}

				//create camera
				buffer.WriteString(RenderCamera(scene.Cameras,scene.Objects))
	   			buffer.WriteString("scene.activeCamera = "+scene.MainCamera+";\n")

	   			//prevent mouse control
	   			if isControlCamera {
	   				buffer.WriteString("scene.activeCamera.attachControl(canvas,false);\n");
	   			}

				//create animation
				for _ , animation := range scene.Animations {
					name := "animate_"+animation.RefName
					buffer.WriteString("var "+name+" = new BABYLON.Animation(\""+animation.Name+"\",\""+animation.Kind+"\","+animation.FrameRate+","+animation.FrameValueType+","+animation.Type+");\n")

					//loop for pushing animation
					if animation.Frames != nil && len(animation.Frames) > 0{
						buffer.WriteString("var "+name+"_key = [];\n")

						//this type don't use parse float function in javascript	
						if animation.Type == "BABYLON.Animation.ANIMATIONLOOPMODE_RELATIVE" {
							for _, frame := range animation.Frames {
						    	buffer.WriteString(name+"_key.push({frame: \""+frame.Frame+"\",value: "+frame.Value+"});\n")
						    }
						}else{
							for _, frame := range animation.Frames {
						    	buffer.WriteString(name+"_key.push({frame: \""+frame.Frame+"\",value: parseFloat(\""+frame.Value+"\")});\n")
						    }
						}

					   	buffer.WriteString(name+".setKeys("+name+"_key);\n");
				    	buffer.WriteString(animation.Target+".animations.push("+name+");\n")

					}
			    
			    	buffer.WriteString("scene.beginAnimation("+animation.Target+", 0, 1000, true);\n")
				}

				//render sprite
			    buffer.WriteString(RenderSprite(scene.Sprites))

				//add pre-render script
		    	buffer.WriteString(RenderPreRenderScript(scene.RenderScript))

				buffer.WriteString("return scene;\n")
			buffer.WriteString("};\n")

			//add scene to movie
			buffer.WriteString("movie[\""+scene.RefName+"\"] = "+scene.RefName+"();\n")
		}

	//init create scene function
	for _ , element := range scenes {
		if element.Id.Hex() == activeSceneHex {
			buffer.WriteString("var labyrinth = "+element.RefName+"();\n")
			buffer.WriteString("var state = \""+element.RefName+"\";\n")
			if is2dEnable {
				buffer.WriteString("movie2d[state]();\n")
			}
		}
	}

	buffer.WriteString("engine.runRenderLoop(function () {\n")
		buffer.WriteString("labyrinth.render();\n")
		//buffer.WriteString("sun_0.rotate.y = sun_0.rotate.y + 1;\n")
	buffer.WriteString("});\n")

	buffer.WriteString("window.addEventListener(\"resize\", function () {\n")
		buffer.WriteString("engine.resize();\n")
	buffer.WriteString("});\n")

	buffer.WriteString("function redirect(scene) {\n")

		//clear 2d in old state
		buffer.WriteString("var images = movie2dImage[state];\n")
		buffer.WriteString("for(var index in images){\n")
		buffer.WriteString("images[index].clear();\n")
		buffer.WriteString("}\n")

		buffer.WriteString("var texts = movie2dText[state];\n")
		buffer.WriteString("for(var index in texts){\n")
		buffer.WriteString("texts[index].clear();\n")
		buffer.WriteString("}\n")

		//set new stage
		buffer.WriteString("var renderer = movie[scene];\n")
		buffer.WriteString("labyrinth = renderer;\n")
		buffer.WriteString("state = scene;\n")

		//draw new element
		if is2dEnable {
			buffer.WriteString("movie2d[state].call();\n")
		}

	buffer.WriteString("};\n")

	//event listener
	buffer.WriteString(RenderAction(scenes))

	buffer.WriteString("</script>")
	
	return buffer.String()	
}