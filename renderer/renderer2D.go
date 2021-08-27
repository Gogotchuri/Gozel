package renderer

import (
	opengl "Gozel/renderer/openGL"
	"Gozel/renderer/render_types"
	"Gozel/renderer/shaders"
	"fmt"
	math32 "github.com/go-gl/mathgl/mgl32"
)

type renderer2D struct {
	VertexArray                render_types.VertexArray
	ColorShader, TextureShader render_types.Shader
}

var Renderer2D = renderer2D{}

func (r *renderer2D) Init() {
	RenderCommand.Init()
	var err error
	r.ColorShader, err = CreateShader("ColorShader", shaders.ColorVertexShaderSrc, shaders.ColorFragmentShaderSrc)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	r.TextureShader, err = CreateShader("TextureShader", shaders.TextureVertexShaderSrc, shaders.TextureFragmentShaderSrc)
	if err != nil {
		panic(err)
	}
	fmt.Println("Shaders created")
	vertices := []float32{
		0.5, 0.5, 0.0, 1.0, 1.0, // top right
		0.5, -0.5, 0.0, 1.0, 0.0, // bottom right
		-0.5, -0.5, 0.0, 0.0, 0.0, // bottom left
		-0.5, 0.5, 0.0, 0.0, 1.0, // top left
	}
	indices := []int32{
		0, 1, 3,
		1, 2, 3,
	}

	vb := CreateVertexBuffer(vertices, 4, 20, opengl.StaticDraw)

	vbLayout := render_types.CreateVertexBufferLayout()
	vbLayout.Push(render_types.CreateLayoutElement(render_types.Float3, "Vertex_coords", false))
	vbLayout.Push(render_types.CreateLayoutElement(render_types.Float2, "Texture_coords", false))
	vb.SetLayout(vbLayout)

	ib := CreateIndexBuffer(indices, 6, opengl.StaticDraw)

	r.VertexArray = CreateVertexArray()
	r.VertexArray.AddVertexBuffer(vb)
	r.VertexArray.SetIndexBuffer(ib)
	r.VertexArray.UnBind()


}

func (r *renderer2D) StartScene(camera OrthographicCamera) {
	viewProj := camera.GetViewProjectionMatrix()

	r.ColorShader.Bind()
	r.ColorShader.SetUniform("u_view_projection", render_types.Mat4, &(*viewProj)[0])

	r.TextureShader.Bind()
	r.TextureShader.SetUniform("u_view_projection", render_types.Mat4, &(*viewProj)[0])
}

func (r *renderer2D) EndScene() {}

func (r *renderer2D) DrawRect(position math32.Vec3, size math32.Vec2, color math32.Vec3,
	rotationAxis math32.Vec3, rotationAngle float32) {
	model := math32.Ident4()
	model = model.Mul4(math32.Translate3D(position[0], position[1], position[2]))
	model = model.Mul4(math32.Scale3D(size[0], size[1], 1.0))
	if rotationAngle != 0 {
		model = model.Mul4(math32.HomogRotate3D(math32.DegToRad(rotationAngle), rotationAxis))
	}

	r.ColorShader.Bind()
	r.ColorShader.SetUniform("u_model", render_types.Mat4, &model[0])
	r.ColorShader.SetUniform("u_color", render_types.Float3, &color[0])

	RenderCommand.DrawIndexed(r.VertexArray)
}
