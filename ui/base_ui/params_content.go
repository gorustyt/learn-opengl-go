package base_ui

import (
	"github.com/gorustyt/fyne/v2"
	"github.com/gorustyt/fyne/v2/widget"
)

type ParamsContent struct {
	widget.BaseWidget
	Objs []fyne.CanvasObject
}

func NewParamsContent() *ParamsContent {
	b := &ParamsContent{}
	b.ExtendBaseWidget(b)
	return b
}

func (p *ParamsContent) MinSize() fyne.Size {
	return fyne.Size{Width: 700, Height: 600}
}

func (p *ParamsContent) Move(position fyne.Position) {

}

func (p *ParamsContent) Position() fyne.Position {
	return fyne.Position{}
}

func (p *ParamsContent) Resize(size fyne.Size) {

}

func (p *ParamsContent) Size() fyne.Size {
	return fyne.Size{}
}

func (p *ParamsContent) Hide() {

}

func (p *ParamsContent) Visible() bool {
	return true
}

func (p *ParamsContent) Show() {

}

func (p *ParamsContent) Refresh() {

}

func (p *ParamsContent) CreateRenderer() fyne.WidgetRenderer {
	return &ParamsContentRender{c: p}
}

type ParamsContentRender struct {
	c *ParamsContent
}

func (p ParamsContentRender) Destroy() {

}

func (p ParamsContentRender) Layout(size fyne.Size) {

}

func (p ParamsContentRender) MinSize() fyne.Size {
	return p.MinSize()
}

func (p ParamsContentRender) Objects() []fyne.CanvasObject {
	return p.c.Objs
}

func (p ParamsContentRender) Refresh() {

}
