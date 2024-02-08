package base_ui

import (
	"github.com/gorustyt/fyne/v2"
	"github.com/gorustyt/fyne/v2/canvas"
	"github.com/gorustyt/fyne/v2/canvas/canvas3d"
	"github.com/gorustyt/fyne/v2/widget"
)

type ChapterContent struct {
	cs []canvas3d.ICanvas3d
	widget.BaseWidget
	*KeyHandel
	WinSize fyne.Size
}

func (c *ChapterContent) Reset() {
	for _, v := range c.cs {
		v.Reset()
	}
}
func (c *ChapterContent) Painter(index int) canvas3d.ICanvas3d {
	return c.cs[index]
}
func (c *ChapterContent) MinSize() fyne.Size {
	return fyne.Size{Width: 700, Height: 600}
}

func (c *ChapterContent) Move(position fyne.Position) {

}

func (c *ChapterContent) Position() fyne.Position {
	return fyne.Position{}
}

func (c *ChapterContent) Resize(size fyne.Size) {

}

func (c *ChapterContent) Size() fyne.Size {
	return fyne.Size{}
}

func (c *ChapterContent) Hide() {

}

func (c *ChapterContent) Visible() bool {
	return true
}

func (c *ChapterContent) Show() {

}

func (c *ChapterContent) Refresh() {
	for _, v := range c.cs {
		canvas.Refresh(v.GetRenderObj())
	}

}

func NewChapterContent() *ChapterContent {
	c := &ChapterContent{KeyHandel: NewKeyHandel(), cs: []canvas3d.ICanvas3d{
		canvas3d.NewCanvas3d(),
		canvas3d.NewCanvas3d(),
		canvas3d.NewCanvas3d(),
	}}
	c.ExtendBaseWidget(c)
	return c
}
func (c *ChapterContent) CreateRenderer() fyne.WidgetRenderer {
	return NewChapterContentRender(c)
}

type ChapterContentRender struct {
	c *ChapterContent
}

func (c ChapterContentRender) Destroy() {

}

func (c ChapterContentRender) Layout(size fyne.Size) {

}

func (c ChapterContentRender) MinSize() fyne.Size {
	return c.c.MinSize()
}

func (c ChapterContentRender) Objects() (res []fyne.CanvasObject) {
	for _, v := range c.c.cs {
		res = append(res, v.GetRenderObj())
	}
	return res
}

func (c ChapterContentRender) Refresh() {
	c.c.Refresh()
}

func NewChapterContentRender(c *ChapterContent) fyne.WidgetRenderer {
	return &ChapterContentRender{c: c}
}
