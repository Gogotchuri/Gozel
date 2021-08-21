package renderer

import (
	opengl "Gozel/renderer/openGL"
	"Gozel/renderer/render_types"
	"fmt"
)

func CreateVertexArray() (render_types.VertexArray, error) {
	switch render_types.CurrentPlatform {
	case render_types.None:
		return nil, fmt.Errorf("platform none")
	case render_types.OpenGL:
		return opengl.CreateGLVertexArray()
	}
	return nil, fmt.Errorf("platform is undefined")
}