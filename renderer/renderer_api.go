package renderer

import (
	opengl "Gozel/renderer/openGL"
	"Gozel/renderer/render_types"
	"fmt"
	//math "github.com/go-gl/mathgl/mgl32"
)

func CreateRendererAPI() (render_types.RendererAPI, error){
	switch render_types.CurrentPlatform {
	case render_types.None:
		return nil, fmt.Errorf("platform none")
	case render_types.OpenGL:
		return opengl.CreateGLRendererAPI()
	}
	return nil, fmt.Errorf("platform is undefined")

}