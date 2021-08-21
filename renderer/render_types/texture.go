package render_types


type Texture interface {
	Bind(slot uint32)
	UnBind()

	GetWidth() int
	GetHeight() int

	Destroy()
}