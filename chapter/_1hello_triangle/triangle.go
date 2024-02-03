package _1hello_triangle

import (
	"github.com/gorustyt/fyne/v2/canvas3d"
	"github.com/gorustyt/learn-opengl-go/ui/base_ui"
)

var (
	vertices = []float32{
		-0.5, -0.5, 0.0, // Left
		0.5, -0.5, 0.0, // Right
		0.0, 0.5, 0.0, // Top
	}
	vertexShader = `
	#version 330 core
	layout (location = 0) in vec3 position;
	void main() {
	  gl_Position = vec4(position.x, position.y, position.z, 1.0);
}`

	fragShader = `
	#version 330 core
	out vec4 color;
	void main() {
		color = vec4(1.0f, 1.0f, 0.2f, 1.0f);
	}`
)

type Triangle struct {
	vert *canvas3d.VertexFloat32Array
}

func NewTriangle() base_ui.IChapter {
	t := &Triangle{vert: canvas3d.NewVertexFloat32Array()}
	t.vert.Arr = vertices
	t.vert.PositionSize = []int{3, 0}
	return t
}

func (t *Triangle) InitChapterContent(c *base_ui.ChapterContent) {
	c.SetShaderConfig(vertexShader, fragShader)
	c.AppendObj(t.vert)
}
func (t *Triangle) InitParamsContent(c *base_ui.ParamsContent) {

}
