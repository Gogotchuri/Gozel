package main

import (
	"Gozel/renderer"
	"fmt"
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"runtime"
	"time"
)

// func init() {
// 	// This is needed to arrange that main() runs on main thread.
// 	// See documentation for functions that are only allowed to be called from the main thread.
// 	runtime.LockOSThread()
// }

func ballast(size int) func() {
	blst := make([]byte, size)
	return func() {
		runtime.KeepAlive(blst)
	}
}

func main() {
	OurInterface()
	//native()
}

func OurInterface()  {
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
		fmt.Println(time.Since(s))
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