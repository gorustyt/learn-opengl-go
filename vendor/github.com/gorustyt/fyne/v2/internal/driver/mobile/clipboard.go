package mobile

import (
	"github.com/gorustyt/fyne/v2"
)

// Declare conformity with Clipboard interface
var _ fyne.Clipboard = (*mobileClipboard)(nil)

// mobileClipboard represents the system mobileClipboard
type mobileClipboard struct {
}
