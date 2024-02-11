package base_ui

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/gorustyt/fyne/v2"
	"github.com/gorustyt/fyne/v2/container"
	"github.com/gorustyt/fyne/v2/data/binding"
	"github.com/gorustyt/fyne/v2/widget"
)

type Vec3Slider struct {
	name string
	sx   *widget.Slider
	sy   *widget.Slider
	sz   *widget.Slider

	pos mgl32.Vec3
}

func NewVec3Slider(label string, initPos mgl32.Vec3) *Vec3Slider {
	v := &Vec3Slider{name: label}
	v.sx = v.GetSlider(initPos.X(), func(f float64) {
		v.pos[0] = float32(f)
		chapterRefresh()
	})
	v.sy = v.GetSlider(initPos.Y(), func(f float64) {
		v.pos[1] = float32(f)
		chapterRefresh()

	})
	v.sz = v.GetSlider(initPos.Z(), func(f float64) {
		v.pos[2] = float32(f)
		chapterRefresh()
	})
	v.pos = initPos
	return v
}

func (s *Vec3Slider) GetSlider(value float32, fn func(f float64)) *widget.Slider {
	f := float64(value)
	sl := widget.NewSliderWithData(0, 1, binding.BindFloat(&f))
	sl.Step = 0.01
	sl.OnChanged = func(f float64) {
		fn(f)
	}
	return sl
}
func (s *Vec3Slider) GetPos() mgl32.Vec3 {
	return s.pos
}
func (s *Vec3Slider) GetRenderObj() fyne.CanvasObject {
	return container.NewVBox(widget.NewLabel(s.name),
		widget.NewLabel(fmt.Sprintf("%v.x", s.name)),
		s.sx,
		widget.NewLabel(fmt.Sprintf("%v.y", s.name)),
		s.sy,
		widget.NewLabel(fmt.Sprintf("%v.z", s.name)),
		s.sz)
}
