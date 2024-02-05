package gl

import "github.com/gorustyt/fyne/v2"

type Painter3D struct {
	prog Program //sharder
	context
}

func NewPainter3D(ctx context) *Painter3D {
	return &Painter3D{context: ctx}
}

func (p *Painter3D) Program() Program {
	return p.prog
}

func (p *Painter3D) HasShader() bool {
	return p.prog == 0
}

func (p *Painter3D) DrawTrianglesByElement(index []uint32) {
	p.context.DrawElementsArrays(triangles, index)
}

func (p *Painter3D) DrawTriangles(count int) {
	p.context.DrawArrays(triangles, 0, count)
}

func (p *Painter3D) DefineVertexArray(name string, size, stride, offset int) {
	vertAttrib := p.GetAttribLocation(p.prog, name)
	p.context.EnableVertexAttribArray(vertAttrib)
	p.VertexAttribPointerWithOffset(vertAttrib, size, float, false, stride*floatSize, offset*floatSize)
}

func (p *Painter3D) BindTexture(texture Texture) {
	p.context.BindTexture(texture2D, texture)
}

type Canvas3D interface {
	InitOnce(p *Painter3D)
	Init(p *Painter3D)
	After(p *Painter3D)
}

type Canvas3DBeforePainter interface {
	BeforeDraw(p *Painter3D, pos fyne.Position, Frame fyne.Size)
}

type Canvas3DPainter interface {
	Draw(p *Painter3D, pos fyne.Position, Frame fyne.Size)
}

type Canvas3dObj struct {
	Painter          *Painter3D
	Objs             []Canvas3D
	VertStr, FragStr string
}

func (c *Canvas3dObj) InitOnce() {
	for _, v := range c.Objs {
		v.InitOnce(c.Painter)
	}
}

func (c *Canvas3dObj) Init() {
	c.Painter.EnableDepthTest()
	for _, v := range c.Objs {
		v.Init(c.Painter)
	}
}

func (c *Canvas3dObj) BeforeDraw(pos fyne.Position, frame fyne.Size) {
	for _, v := range c.Objs {
		if cc, ok := v.(Canvas3DBeforePainter); ok {
			cc.BeforeDraw(c.Painter, pos, frame)
		}
	}
}

func (c *Canvas3dObj) Draw(pos fyne.Position, frame fyne.Size) {
	for _, v := range c.Objs {
		if cc, ok := v.(Canvas3DPainter); ok {
			cc.Draw(c.Painter, pos, frame)
		}
	}
}

func (c *Canvas3dObj) After() {
	for i := len(c.Objs) - 1; i >= 0; i-- {
		c.Objs[i].After(c.Painter)
	}
	c.Painter.DisableDepthTest()
}

func (c *Canvas3dObj) Dragged(ev *fyne.DragEvent) {
	for _, v := range c.Objs {
		if p, ok := v.(fyne.Draggable); ok {
			p.Dragged(ev)
		}
	}
}
func (c *Canvas3dObj) DragEnd() {
	for _, v := range c.Objs {
		if p, ok := v.(fyne.Draggable); ok {
			p.DragEnd()
		}
	}
}
func (c *Canvas3dObj) Scrolled(ev *fyne.ScrollEvent) {
	for _, v := range c.Objs {
		if p, ok := v.(fyne.Scrollable); ok {
			p.Scrolled(ev)
		}
	}
}
func (c *Canvas3dObj) Move(position fyne.Position) {

}

func (c *Canvas3dObj) Position() fyne.Position {
	return fyne.Position{}
}

func (c *Canvas3dObj) Hide() {

}

func (c *Canvas3dObj) Visible() bool {
	return true
}

func (c *Canvas3dObj) Show() {

}

func (c *Canvas3dObj) MinSize() fyne.Size {
	return fyne.Size{Width: 600, Height: 600}
}

func (c *Canvas3dObj) Resize(size fyne.Size) {

}

func (c *Canvas3dObj) Size() fyne.Size {
	return fyne.Size{Width: 600, Height: 600}
}

func (c *Canvas3dObj) Refresh() {

}

func NewCustomObj() *Canvas3dObj {
	return &Canvas3dObj{}
}
