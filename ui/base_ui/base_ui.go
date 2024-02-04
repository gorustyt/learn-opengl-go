package base_ui

import "github.com/gorustyt/fyne/v2"

type IChapter interface {
	InitChapterContent(c *ChapterContent)
	InitParamsContent(c *ParamsContent)
}

type KeyHandel struct {
	OnKeyUp   func(ev *fyne.KeyEvent)
	OnKeyDown func(ev *fyne.KeyEvent)
	OnTyped   func(ev *fyne.KeyEvent)
}

func NewKeyHandel() *KeyHandel {
	return &KeyHandel{
		OnKeyDown: func(ev *fyne.KeyEvent) {},
		OnKeyUp:   func(ev *fyne.KeyEvent) {},
		OnTyped:   func(ev *fyne.KeyEvent) {},
	}
}
