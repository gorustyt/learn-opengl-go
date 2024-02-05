package canvas3d

import (
	"github.com/gorustyt/fyne/v2"
	"github.com/gorustyt/fyne/v2/internal/painter/gl"
)

var (
	_ gl.Canvas3D        = (*VertexFloat32Array)(nil)
	_ gl.Canvas3DPainter = (*VertexFloat32Array)(nil)
)

type VertexFloat32Array struct {
	Arr          []float32
	Index        []uint32
	PositionSize []int //大小和偏移
	NormalSize   []int
	TexCoordSize []int
	ColorSize    []int
	vbo, veo     gl.Buffer
}

func (v *VertexFloat32Array) InitOnce(p *gl.Painter3D) {

}

func NewVertexFloat32Array() *VertexFloat32Array {
	return &VertexFloat32Array{}
}

func (v *VertexFloat32Array) Init(p *gl.Painter3D) {

}

func (v *VertexFloat32Array) Draw(p *gl.Painter3D, pos fyne.Position, frame fyne.Size) {
	points := v.AddOffset(v.Arr, pos, frame)
	if len(v.Index) > 0 {
		v.vbo, v.veo = p.MakeVaoWithEbo(points, v.Index)
	} else {
		v.vbo = p.MakeVao(points)
	}
	stride := v.getStride()
	if len(v.PositionSize) == 2 {
		p.DefineVertexArray("position", v.PositionSize[0], stride, v.PositionSize[1])
	}
	if len(v.NormalSize) == 2 {
		p.DefineVertexArray("normal", v.NormalSize[0], stride, v.NormalSize[1])
	}
	if len(v.TexCoordSize) == 2 {
		p.DefineVertexArray("texCoord", v.TexCoordSize[0], stride, v.TexCoordSize[1])
	}
	if len(v.ColorSize) == 2 {
		p.DefineVertexArray("color", v.ColorSize[0], stride, v.ColorSize[1])
	}
	if len(v.Index) != 0 {
		p.DrawTrianglesByElement(v.Index)
	} else {
		p.DrawTriangles(len(points) / stride)
	}

}

func (v *VertexFloat32Array) getStride() (stride int) {
	if len(v.PositionSize) == 2 {
		stride += v.PositionSize[0]
	}
	if len(v.NormalSize) == 2 {
		stride += v.NormalSize[0]
	}
	if len(v.TexCoordSize) == 2 {
		stride += v.TexCoordSize[0]
	}
	if len(v.ColorSize) == 2 {
		stride += v.ColorSize[0]
	}
	return
}

func (v *VertexFloat32Array) AddOffset(points []float32, pos fyne.Position, size fyne.Size) (res []float32) {
	stride := v.getStride()
	offsetX := pos.X / size.Width
	offsetY := pos.Y / size.Height
	res = make([]float32, len(points))
	copy(res, points)
	for i := 0; i < len(points); i += stride {
		begin := i + v.PositionSize[1]
		res[begin] += offsetX
		res[begin+1] += offsetY
	}

	for i := 0; i < len(points); i += stride {
		begin := i + v.TexCoordSize[1]
		res[begin] += offsetX
		res[begin+1] += offsetY
	}
	return res
}
func (v *VertexFloat32Array) After(p *gl.Painter3D) {
	if v.vbo != 0 {
		p.DeleteBuffer(v.vbo)
	}
	if v.veo != 0 {
		p.DeleteBuffer(v.veo)
	}
}
