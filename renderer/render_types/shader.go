package render_types


type Shader interface {
	Bind()
	UnBind()

	GetID() uint32
	GetName() string

	SetUniform(name string, uType ShaderDataType, data interface{})

	Destroy()
}

