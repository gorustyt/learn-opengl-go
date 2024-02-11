package base_ui

import (
	"github.com/gorustyt/fyne/v2"
	"github.com/gorustyt/fyne/v2/canvas3d/canvas3d_render"
	"github.com/gorustyt/fyne/v2/widget"
)

type ChapterContent struct {
	c canvas3d_render.ICanvas3d
	widget.BaseWidget
	*KeyHandel
	WinSize fyne.Size
}

func (c *ChapterContent) Canvas3d() canvas3d_render.ICanvas3d {
	return c.c
}

func (c *ChapterContent) CreateRenderer() fyne.WidgetRenderer {
	return &ChapterContentRender{c: c}
}

func (c *ChapterContent) MinSize() fyne.Size {
	return fyne.Size{Width: 700, Height: 600}
}

func (c *ChapterContent) Refresh() {
	c.c.Refresh()
}

func NewChapterContent() *ChapterContent {
	c := &ChapterContent{KeyHandel: NewKeyHandel(), c: canvas3d_render.NewCanvas3d(2)}
	c.ExtendBaseWidget(c)
	return c
}

type ChapterContentRender struct {
	c *ChapterContent
}

func (c *ChapterContentRender) Destroy() {

}

func (c *ChapterContentRender) Layout(size fyne.Size) {
	c.c.Resize(size)
}

func (c *ChapterContentRender) MinSize() fyne.Size {
	return c.c.MinSize()
}

func (c *ChapterContentRender) Objects() (res []fyne.CanvasObject) {
	return append(res, c.c.c.GetRenderObj())
}

func (c *ChapterContentRender) Refresh() {
	c.c.Refresh()
}
