package chapter

import (
	"github.com/gorustyt/learn-opengl-go/chapter/_1hello_triangle"
	"github.com/gorustyt/learn-opengl-go/chapter/desc"
	"github.com/gorustyt/learn-opengl-go/ui/base_ui"
)

type Chapter struct {
	ParamsContent  *base_ui.ParamsContent
	ChapterContent *base_ui.ChapterContent
	Chapters       map[string]base_ui.IChapter
}

func (c *Chapter) ChangeChapter(uid string) {
	v, ok := c.Chapters[uid]
	if !ok {
		switch uid {
		case desc.ChapterHelloTriangleSub1:
			v = _1hello_triangle.NewTriangle()
		case desc.ChapterHelloTriangleSub2:
			v = _1hello_triangle.NewTriangleIndex()
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

func NewChapter() *Chapter {
	return &Chapter{
		Chapters:       map[string]base_ui.IChapter{},
		ParamsContent:  base_ui.NewParamsContent(),
		ChapterContent: base_ui.NewChapterContent(),
	}
}
