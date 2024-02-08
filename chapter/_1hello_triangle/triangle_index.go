package _1hello_triangle

import (
	"github.com/gorustyt/fyne/v2/canvas3d"
	"github.com/gorustyt/learn-opengl-go/ui/base_ui"
)

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
	c.Painter(0).SetShaderConfig(vertexShader, fragShader)
	c.Painter(0).AppendObj(t.vert)
}
func (t *TriangleIndex) InitParamsContent(c *base_ui.ParamsContent) {

}
