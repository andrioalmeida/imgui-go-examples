// +build sdl

package main

import (
	"fmt"
	"os"

	"github.com/inkyblackness/imgui-go/v2"

	"github.com/inkyblackness/imgui-go-examples/internal/demo"
	"github.com/inkyblackness/imgui-go-examples/internal/platforms"
	"github.com/inkyblackness/imgui-go-examples/internal/renderers"
)

func main() {
	context := imgui.CreateContext(nil)
	defer context.Destroy()
	io := imgui.CurrentIO()

	platform, err := platforms.NewSDL(io, platforms.SDLClientAPIOpenGL2)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}
	defer platform.Dispose()

	renderer, err := renderers.NewOpenGL2(io)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}
	defer renderer.Dispose()

	demo.Run(platform, renderer)
}
