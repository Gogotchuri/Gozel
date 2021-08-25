package renderer

import (
	opengl "Gozel/renderer/openGL"
	"Gozel/renderer/render_types"
	"fmt"
)

func CreateVertexArray() render_types.VertexArray {
	switch render_types.CurrentPlatform {
	case render_types.None:
		fmt.Println("platform none")
		return nil //TODO null safety?

	case render_types.OpenGL:
		return opengl.CreateGLVertexArray()
	}
	fmt.Println("platform is undefined")
	return nil
}
