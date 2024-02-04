package main

import (
	"github.com/gorustyt/fyne/v2"
	"github.com/gorustyt/fyne/v2/app"
	"github.com/gorustyt/fyne/v2/container"
	"github.com/gorustyt/learn-opengl-go/chapter"
	"github.com/gorustyt/learn-opengl-go/ui"
)

func main() {
	width, height := 1200, 900
	a := app.NewWithID("learn-go-opengl")
	w := a.NewWindow("learn go opengl")
	chap := chapter.NewChapter(w)
	content := container.NewStack()
	setView := func(uid string) {
		chap.ChangeChapter(uid)
		content.Objects = []fyne.CanvasObject{chap.ChapterContent}
		content.Refresh()
	}
	split := container.NewHSplit(chap.ParamsContent, content)
	split.Offset = 0.2
	split1 := container.NewHSplit(ui.CreateChapterTree(setView), split)
	split1.Offset = 0.2
	w.SetContent(split1)
	w.Resize(fyne.NewSize(float32(width), float32(height)))
	w.ShowAndRun()
}
