package ui

import (
	"github.com/gorustyt/fyne/v2"
	"github.com/gorustyt/fyne/v2/container"
	"github.com/gorustyt/fyne/v2/widget"
	"github.com/gorustyt/learn-opengl-go/chapter/desc"
)

const preferenceCurrentTutorial = "currentTutorial"

var (
	treeData = map[string][]string{
		"":                        {desc.ChapterHelloTriangle, desc.ChapterShader},
		desc.ChapterHelloTriangle: {desc.ChapterHelloTriangleSub1, desc.ChapterHelloTriangleSub2},
		desc.ChapterShader:        {},
	}
)

func CreateChapterTree(view func(uid string)) fyne.CanvasObject {
	a := fyne.CurrentApp()

	tree := &widget.Tree{
		ChildUIDs: func(uid string) []string {
			return treeData[uid]
		},
		IsBranch: func(uid string) bool {
			children, ok := treeData[uid]

			return ok && len(children) > 0
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("Collection Widgets")
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(uid)
		},
		OnSelected: func(uid string) {
			view(uid)
		},
	}
	currentPref := a.Preferences().StringWithFallback(preferenceCurrentTutorial, "welcome")
	tree.Select(currentPref)

	return container.NewBorder(nil, nil, nil, nil, tree)
}
