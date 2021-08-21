package render_types


type VertexBuffer interface {
	GetLayout() *VertexBufferLayout
	SetLayout(vbi VertexBufferLayout)

	GetCount() int32

	Bind()
	UnBind()

	Destroy()
}


type IndexBuffer interface {
	GetCount() int32

	Bind()
	UnBind()

	Destroy()
}

