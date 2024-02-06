package base_ui

import (
	"github.com/gorustyt/fyne/v2"
	"github.com/gorustyt/fyne/v2/container"
)

type ParamsContent struct {
	*fyne.Container
	*KeyHandel
}

func NewParamsContent() *ParamsContent {
	b := &ParamsContent{KeyHandel: NewKeyHandel(), Container: container.NewStack()}
	b.Container.Resize(fyne.Size{Width: 100, Height: 800})
	return b
}
