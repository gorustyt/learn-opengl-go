package chapter

import (
	"github.com/gorustyt/fyne/v2"
	"github.com/gorustyt/fyne/v2/driver/desktop"
	"github.com/gorustyt/learn-opengl-go/chapter/_1hello_triangle"
	"github.com/gorustyt/learn-opengl-go/chapter/desc"
	"github.com/gorustyt/learn-opengl-go/ui/base_ui"
)

type Chapter struct {
	w              fyne.Window
	ParamsContent  *base_ui.ParamsContent
	ChapterContent *base_ui.ChapterContent
	Chapters       map[string]base_ui.IChapter
}

func (c *Chapter) ChangeChapter(uid string) {
	v, ok := c.Chapters[uid]
	if !ok {
		cns := chapterCns[uid]
		if cns != nil {
			v = cns()
		}
		if v != nil {
			c.Chapters[uid] = v
		}
	}
	if v != nil {
		c.ChapterContent.Reset()
		v.InitChapterContent(c.ChapterContent)
		v.InitParamsContent(c.ParamsContent)
		c.ChapterContent.Refresh()
		c.ParamsContent.Refresh()
	}

}

func NewChapter(w fyne.Window, winSize fyne.Size) *Chapter {
	c := &Chapter{
		w:              w,
		Chapters:       map[string]base_ui.IChapter{},
		ParamsContent:  base_ui.NewParamsContent(),
		ChapterContent: base_ui.NewChapterContent(),
	}
	c.ChapterContent.WinSize = winSize
	c.RegisterKeyEvent()
	return c
}

func (c *Chapter) RegisterKeyEvent() {
	c.w.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
		if c.ParamsContent.OnTyped != nil {
			c.ParamsContent.OnTyped(event)
		}
		if c.ChapterContent.OnTyped != nil {
			c.ChapterContent.OnTyped(event)
		}

	})
	if deskCanvas, ok := c.w.Canvas().(desktop.Canvas); ok {
		deskCanvas.SetOnKeyDown(func(event *fyne.KeyEvent) {
			if c.ParamsContent.OnKeyDown != nil {
				c.ParamsContent.OnKeyDown(event)
			}
			if c.ChapterContent.OnKeyDown != nil {
				c.ChapterContent.OnKeyDown(event)
			}
		})
		deskCanvas.SetOnKeyUp(func(event *fyne.KeyEvent) {
			if c.ParamsContent.OnKeyUp != nil {
				c.ParamsContent.OnKeyUp(event)
			}
			if c.ChapterContent.OnKeyUp != nil {
				c.ChapterContent.OnKeyUp(event)
			}
		})
	}
}

var chapterCns = map[string]func() base_ui.IChapter{
	desc.ChapterHelloTriangleSub1: _1hello_triangle.NewTriangle,
	desc.ChapterHelloTriangleSub2: _1hello_triangle.NewTriangleIndex,
	desc.ChapterHelloTriangleSub3: _1hello_triangle.NewHelloCoordinates,
}
