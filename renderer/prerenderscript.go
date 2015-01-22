package renderer

import (
	"bytes"
)

func RenderPreRenderScript(script string) string{

	var buffer bytes.Buffer

	buffer.WriteString("scene.registerBeforeRender(function () {\n")
		buffer.WriteString(script)
	buffer.WriteString("});\n")

	return buffer.String()
}