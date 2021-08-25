package opengl

import (
	"Gozel/renderer/render_types"
	"github.com/go-gl/gl/all-core/gl"
)

var _ render_types.RendererAPI = &GLRendererAPI{}

type GLRendererAPI struct{}

func CreateGLRendererAPI() *GLRendererAPI {
	return &GLRendererAPI{}
}

func (renderer *GLRendererAPI) Init() {
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Enable(gl.DEPTH_TEST)
}

func (renderer *GLRendererAPI) Destroy() {}

func (renderer *GLRendererAPI) SetViewport(x, y, width, height int32) {
	gl.Viewport(x, y, width, height)
}

func (renderer *GLRendererAPI) SetClearColor(r, g, b, a float32) {
	gl.ClearColor(r, g, b, a)
}

func (renderer *GLRendererAPI) Clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (renderer *GLRendererAPI) DrawVa(va render_types.VertexArray) {
	gl.DrawElements(gl.TRIANGLES, va.GetIndexBuffer().GetCount(), gl.UNSIGNED_INT, gl.Ptr(nil))
	gl.BindTexture(gl.TEXTURE_2D, 0)
}

func (renderer *GLRendererAPI) GetAPI() render_types.RenderPlatform {
	return render_types.OpenGL
}
