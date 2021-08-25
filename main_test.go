package main

import (
	"Gozel/renderer"
	"github.com/go-gl/glfw/v3.3/glfw"
	"testing"
)

func BenchmarkOurInterface(b *testing.B) {
	window, err := renderer.CreateWindow(600, 800, "WW")
	if err != nil {
		panic(err)
	}
	for i := 0; i < b.N; i++ {
		window.OnUpdate()
		glfw.PollEvents()
	}
}
