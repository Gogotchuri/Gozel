package renderer

import (
	"Gozel/renderer/render_types"
)

type renderCommand struct {
	rendererAPI render_types.RendererAPI
}

var RenderCommand = &renderCommand{rendererAPI: nil}

func (rc *renderCommand) Init() {
	rc.rendererAPI = CreateRendererAPI()
}

func (rc *renderCommand) SetViewport(x, y, width, height int32) {
	rc.rendererAPI.SetViewport(x, y, width, height)
}

func (rc *renderCommand) SetClearColor(r, g, b, a float32) {
	rc.rendererAPI.SetClearColor(r, g, b, a)
}

func (rc *renderCommand) Clear() {
	rc.rendererAPI.Clear()
}

func (rc *renderCommand) DrawIndexed(va render_types.VertexArray) {
	rc.rendererAPI.DrawVa(va)
}
