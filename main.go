package main

import (
	"Gozel/renderer"
	"fmt"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	math32 "github.com/go-gl/mathgl/mgl32"
	"time"
)

// func init() {
// 	// This is needed to arrange that main() runs on main thread.
// 	// See documentation for functions that are only allowed to be called from the main thread.
// 	runtime.LockOSThread()
// }

func main() {
	HelloSquare()
	//OurInterface()
	//native()
}

func HelloSquare() {
	window, err := renderer.CreateWindow(600, 800, "WW")
	if err != nil {
		panic(err)
	}
	window.SetVSync(true)
	renderer.Renderer2D.Init()
	oCamera := renderer.CreateOrthographicCamera(0, 600, 0, 800)
	oCamera.SetPosition(math32.Vec3{100, 100, 0})
	var dx, dy, dz, rotation float32 = 300, 400, 0, 0
	for window.IsOpen() {
		start := time.Now()
		renderer.RenderCommand.SetClearColor(0.2, 0.3, 0.3, 1.0)
		renderer.RenderCommand.Clear()

		//var dx, dy, dz, rotation float32
		if window.GetBaseWindow().(*glfw.Window).GetKey(glfw.KeyLeft) == glfw.Press {
			dx += 10
		}
		if window.GetBaseWindow().(*glfw.Window).GetKey(glfw.KeyRight) == glfw.Press{
			dx -= 10
		}
		if window.GetBaseWindow().(*glfw.Window).GetKey(glfw.KeyUp) == glfw.Press {
			dy += 10
		}
		if window.GetBaseWindow().(*glfw.Window).GetKey(glfw.KeyDown) == glfw.Press {
			dy -= 10
		}
		if window.GetBaseWindow().(*glfw.Window).GetKey(glfw.KeySpace) == glfw.Press {
			rotation = 45
		}
		//oCamera.Move(dx, dy, dz)
		//fmt.Println(oCamera.GetPosition())
		//oCamera.Rotate(rotation)
		fmt.Println(dx, dy, dz)
		renderer.Renderer2D.StartScene(oCamera)
		renderer.Renderer2D.DrawRect(math32.Vec3{dx, dy, dz}, math32.Vec2{150, 150}, math32.Vec3{1, 0.5, 0.5}, math32.Vec3{0,0,1}, rotation)
		renderer.Renderer2D.EndScene()

		window.OnUpdate()
		glfw.PollEvents()
		fmt.Println(time.Since(start).Milliseconds())
	}
}

func OurInterface() {
	window, err := renderer.CreateWindow(600, 800, "WW")
	if err != nil {
		panic(err)
	}
	//window.SetVSync(true)
	//ds := &[]time.Duration{}
	for window.IsOpen() {
		s := time.Now()
		// Do openGL stuff.
		window.OnUpdate()
		glfw.PollEvents()
		//*ds = append(*ds, time.Since(s))
		fmt.Println("--", time.Since(s).String())
	}
	window.Close()
	//fmt.Println(ds)

	renderer.RenderCommand.Init()
}

func native() {
	/* Initialize the library */
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	/* Create a windowed mode window and its OpenGL context */
	window, err := glfw.CreateWindow(600, 800, "WW", nil, nil)
	if err != nil {
		glfw.Terminate()
		panic(err)
	}

	/* Make the window's context current */
	//window.MakeContextCurrent()
	window.MakeContextCurrent()
	glfw.SwapInterval(0)

	/*Viewports for scaling and callback*/
	cb := func(w *glfw.Window, width int, height int) {
		gl.Viewport(0, 0, int32(width), int32(height))
	}
	window.SetFramebufferSizeCallback(cb)

	for !window.ShouldClose() {
		s := time.Now()

		window.SwapBuffers()
		glfw.PollEvents()

		fmt.Println(time.Since(s))

	}
	glfw.Terminate()
}
