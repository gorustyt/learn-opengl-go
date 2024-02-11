package chapter

import (
	"github.com/gorustyt/fyne/v2"
	"github.com/gorustyt/fyne/v2/driver/desktop"
	"github.com/gorustyt/learn-opengl-go/chapter/_1hello_triangle"
	"github.com/gorustyt/learn-opengl-go/chapter/_2lighting"
	"github.com/gorustyt/learn-opengl-go/chapter/desc"
	"github.com/gorustyt/learn-opengl-go/ui/base_ui"
	"log/slog"
	"sync"
	"time"
)

type Chapter struct {
	w              fyne.Window
	ParamsContent  *base_ui.ParamsContent
	ChapterContent *base_ui.ChapterContent
	Chapters       map[string]base_ui.IChapter
	ticker         *time.Ticker
	exit           chan struct{}
	curChapter     base_ui.IChapter

	lock sync.Mutex
}

func (c *Chapter) TimerRefresh() {
	if c.curChapter == nil {
		return
	}
	if c.ticker != nil {
		c.ticker.Stop()
		c.exit <- struct{}{}
		c.ticker = nil
	}
	if v, ok := c.curChapter.(base_ui.IChapterRefresh); ok {
		c.ticker = time.NewTicker(v.RefreshInterVal())
		go func() {
			for {
				select {
				case <-c.ticker.C:
					c.refresh()
				case <-c.exit:
					slog.Info("timer exit ....")
					return
				}
			}
		}()
	}
}

func (c *Chapter) refresh() {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.curChapter == nil {
		return
	}
	c.ChapterContent.Canvas3d().Reset()
	c.ParamsContent.RemoveAll()
	c.curChapter.InitChapterContent(c.ChapterContent)
	c.curChapter.InitParamsContent(c.ParamsContent)
	c.ChapterContent.Refresh()
	c.ParamsContent.Refresh()

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
	c.curChapter = v
	c.refresh()
	c.TimerRefresh()
}

func NewChapter(w fyne.Window, winSize fyne.Size) *Chapter {
	c := &Chapter{
		w:              w,
		Chapters:       map[string]base_ui.IChapter{},
		ParamsContent:  base_ui.NewParamsContent(),
		ChapterContent: base_ui.NewChapterContent(),
		exit:           make(chan struct{}, 1),
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
	desc.ChapterLighting1:         _2lighting.NewLighting,
	desc.ChapterLighting2:         _2lighting.NewLight1,
	desc.ChapterLighting3:         _2lighting.NewLight2,
	desc.ChapterLighting7:         _2lighting.NewMaterial,
}
