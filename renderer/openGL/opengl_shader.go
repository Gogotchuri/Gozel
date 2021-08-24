package opengl

import (
	"Gozel/renderer/render_types"
	"fmt"
	"github.com/go-gl/gl/all-core/gl"
	"strings"
)

var _ render_types.Shader = &GLShader{}

type GLShader struct {
	Name, VertexSrc, FragmentSrc string
	GLProgramID uint32
}

func CreateGLShader(name, vertexSrc, fragmentSrc string) (*GLShader, error) {
	programID := gl.CreateProgram()

	vertexShaderID, err := compileShader(render_types.VertexShader, vertexSrc)
	if err != nil {
		return nil, err
	}
	fragmentShaderID, err := compileShader(render_types.FragmentShader, fragmentSrc)
	if err != nil {
		return nil, err
	}

	gl.AttachShader(programID, vertexShaderID)
	gl.AttachShader(programID, fragmentShaderID)
	gl.LinkProgram(programID)

	var status int32
	gl.GetProgramiv(programID, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		gl.GetProgramiv(programID, gl.INFO_LOG_LENGTH, &status)
		log := strings.Repeat("\x00", int(status+1)) //Buff it up
		gl.GetProgramInfoLog(programID, status, nil, gl.Str(log))
		gl.DeleteProgram(programID)
		return nil, fmt.Errorf("program creation failed: \n %v", log)
	}

	gl.DeleteShader(vertexShaderID)
	gl.DeleteShader(fragmentShaderID)

	return &GLShader{
		Name:        name,
		VertexSrc:   vertexSrc,
		FragmentSrc: fragmentSrc,
		GLProgramID: programID,
	}, nil
}

func CreateGLShaderFromFile(f string) (*GLShader, error) {
	panic("undefined")
}

func (s *GLShader) Bind() {
	gl.UseProgram(s.GLProgramID)
}

func (s *GLShader) UnBind() {
	gl.UseProgram(0)
}

func (s *GLShader) GetID() uint32 {
	return s.GLProgramID
}

func (s *GLShader) GetName() string {
	return s.Name
}

func (s *GLShader) SetUniform(name string, uType render_types.ShaderDataType, data interface{}) {
	uniformLoc := gl.GetUniformLocation(s.GLProgramID, gl.Str(name))
	switch uType {
	case render_types.Float1:
		td := data.(float32)
		gl.Uniform1f(uniformLoc, td)
	case render_types.Float2:
		td := data.(*float32)
		gl.Uniform2fv(uniformLoc, 1, td)
	case render_types.Float3:
		td := data.(*float32)
		gl.Uniform3fv(uniformLoc, 1, td)
	case render_types.Float4:
		td := data.(*float32)
		gl.Uniform4fv(uniformLoc, 1, td)
	case render_types.Int1:
		td := data.(int32)
		gl.Uniform1i(uniformLoc, td)
	case render_types.Int2:
		td := data.(*int32)
		gl.Uniform2iv(uniformLoc, 1, td)
	case render_types.Int3:
		td := data.(*int32)
		gl.Uniform3iv(uniformLoc, 1, td)
	case render_types.Int4:
		td := data.(*int32)
		gl.Uniform4iv(uniformLoc, 1, td)
	case render_types.Mat2:
		td := data.(*float32)
		gl.UniformMatrix2fv(uniformLoc, 1, false, td)
	case render_types.Mat3:
		td := data.(*float32)
		gl.UniformMatrix3fv(uniformLoc, 1, false, td)
	case render_types.Mat4:
		td := data.(*float32)
		gl.UniformMatrix4fv(uniformLoc, 1, false, td)
	}
}

func (s *GLShader) Destroy() {
	gl.DeleteProgram(s.GLProgramID)
}

func compileShader(shaderType render_types.ShaderType, src string) (uint32, error) {
	shaderID := gl.CreateShader(uint32(shaderType))
	cSrc, free := gl.Strs(src)

	gl.ShaderSource(shaderID, 1, cSrc, nil)
	free()
	gl.CompileShader(shaderID)

	var status int32
	gl.GetShaderiv(shaderID, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		gl.GetShaderiv(shaderID, gl.INFO_LOG_LENGTH, &status)

		log := strings.Repeat("\x00", int(status+1)) //Buff it up
		gl.GetShaderInfoLog(shaderID, status, nil, gl.Str(log))
		gl.DeleteShader(shaderID)
		return 0, fmt.Errorf("faile to compile shader: \n %v \n %v", src, log)
	}

	return shaderID, nil
}