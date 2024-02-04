//go:build !hints
// +build !hints

package theme

import (
	"image/color"

	"github.com/gorustyt/fyne/v2"
)

var (
	fallbackColor = color.Transparent
	fallbackIcon  = &fyne.StaticResource{}
)
