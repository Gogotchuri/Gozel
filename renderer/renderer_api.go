package renderer

import (
	opengl "Gozel/renderer/openGL"
	"Gozel/renderer/render_types"
	"fmt"
	//math "github.com/go-gl/mathgl/mgl32"
)

func CreateRendererAPI() render_types.RendererAPI {
	switch render_types.CurrentPlatform {
	case render_types.None:
		fmt.Println("platform none")
		return nil
	case render_types.OpenGL:
		return opengl.CreateGLRendererAPI()
	}
	fmt.Println("platform undefined")
	return nil

}
