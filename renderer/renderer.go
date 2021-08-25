package renderer

import (
	"Gozel/renderer/render_types"
	math32 "github.com/go-gl/mathgl/mgl32"
)

type sceneData struct {
	viewProjectionMatrix math32.Mat4
}

type renderer struct {
	scene sceneData
}

var Renderer = &renderer{}


func (r *renderer) Init()  {
	RenderCommand.Init()
}

func (r *renderer) Shutdown()  {

}


func (r *renderer) BeginScene()  {
	//TODO camera
}

func (r *renderer) EndScene()  {

}

func (r *renderer) Submit(shader render_types.Shader, va render_types.VertexArray, transform * math32.Mat4)  {
	if transform == nil {
		t := math32.Ident4()
		transform = &t
	}
	shader.Bind()
	//TODO Fix those in gl shader
	shader.SetUniform("u_ViewProjection", render_types.Mat4, &r.scene.viewProjectionMatrix[0])
	shader.SetUniform("u_Transform", render_types.Mat4, &transform[0])

	va.Bind()
	RenderCommand.DrawIndexed(va)
}
