package render_types


type RenderPlatform int

//TODO move me away

var CurrentPlatform RenderPlatform = OpenGL
const (
	None RenderPlatform = iota
	OpenGL
)

type RendererAPI interface {
	Init()
	Destroy()

	SetViewport(x, y, width, height int32)
	SetClearColor(r, g, b, a float32)

	Clear()

	DrawVa(va VertexArray)
	GetAPI() RenderPlatform
}
