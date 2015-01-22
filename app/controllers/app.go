package controllers

import (
	"net/http"
	"strings"
	"io/ioutil"
	"github.com/revel/revel"
	"github.com/sutthirak/labyrinth/data"
	"github.com/sutthirak/labyrinth/renderer"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {

	heavendoor , found := revel.Config.String("heavendoor.url")
	if found {

		script := renderer.RenderScene(data.GetSceneList())
		script = strings.Replace(script,"<script>","",-1)
		script = strings.Replace(script,"</script>","",-1)

		response , error := http.Post(heavendoor+"/minify","text/plain",strings.NewReader(script))
		defer response.Body.Close()

		if error == nil {
			body, err := ioutil.ReadAll(response.Body)

			if err == nil {
				script = "<script>" + string(body) + "</script>"
				return c.Render(script)
			}
		}
		
	}

	return c.Render()
}
