# GLFW

Go binding for [GLFW](http://www.glfw.org/), an open-source and multi-platform library for OpenGL.

GLFW is at a very early stage and is still working in progress. Changes to the APIs are expected before the first release.

## Versioning

Versioning of GLFW is done with Git tags following [semantic versioning](https://semver.org/) (`major.minor.patch`). The major and minor version numbers are same to those of the targeted GLFW library version.

The patch version number, however, has no connection with the GLFW library. GLFW will use the latest revision of the GLFW library on each update.

## Example Code

This example is translated from the example code in GLFW's [documentation](http://www.glfw.org/documentation.html).

```go
package main

import (
	"github.com/paperui/glfw"
)

func main() {
	// Initialize the library.
	ctx := glfw3.Init()
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

```

## Credits

Thanks to [go-gl/glfw](https://github.com/go-gl/glfw) for reference on the Cgo build tags.

## License

MIT