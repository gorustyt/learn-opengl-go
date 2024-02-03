package _1hello_triangle

import (
	"github.com/gorustyt/fyne/v2/canvas3d"
	"github.com/gorustyt/learn-opengl-go/ui/base_ui"
)

var vertexShader1 = `#version 330 core
layout (location = 0) in vec3 position;
void main()
{
   gl_Position = vec4(position.x, position.y, position.z, 1.0);
}`

var fragmentShader1 = `#version 330 core
out vec4 FragColor;
void main()
{
  FragColor = vec4(1.0f, 0.5f, 0.2f, 1.0f);
}`

var vertices1 = []float32{
	0.5, 0.5, 0.0, // 右上角
	0.5, -0.5, 0.0, // 右下角
	-0.5, -0.5, 0.0, // 左下角
	-0.5, 0.5, 0.0, // 左上角
}

var indices = []uint32{ // note that we start from 0!
	0, 1, 3, // first Triangle
	1, 2, 3, // second Triangle
}

type TriangleIndex struct {
	vert *canvas3d.VertexFloat32Array
}

func NewTriangleIndex() base_ui.IChapter {
	t := &TriangleIndex{vert: canvas3d.NewVertexFloat32Array()}
	t.vert.Arr = vertices1
	t.vert.PositionSize = []int{3, 0}
	t.vert.Index = indices
	return t
}

func (t *TriangleIndex) InitChapterContent(c *base_ui.ChapterContent) {
	c.SetShaderConfig(vertexShader1, fragmentShader1)
	c.AppendObj(t.vert)
}
func (t *TriangleIndex) InitParamsContent(c *base_ui.ParamsContent) {

}
