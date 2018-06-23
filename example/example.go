// Copyright (c) 2018 Beta Kuang
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"github.com/beta/glfw"
)

// This example is translated from the example code in GLFW's documentation
// (http://www.glfw.org/documentation.html).
func main() {
	// Initialize the library.
	ctx := glfw.Init()
	if ctx == nil {
		panic("failed to initialize GLFW")
	}
	defer ctx.Terminate()

	// Create a windowed mode window and its OpenGL context.
	win := ctx.CreateWindow(640, 480, "Hello World", nil, nil)
	if win == nil {
		panic("failed to create window")
	}
	defer win.Destroy()

	// Make the window's context current.
	ctx.MakeContextCurrent(win)

	// Loop until the user closes the window.
	for !win.ShouldClose() {
		// Render here.

		// Swap front and back buffers.
		win.SwapBuffers()

		// Poll for and process events.
		ctx.PollEvents()
	}
}
