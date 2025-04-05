package main

import (
	"log/slog"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func render() {
	myApp := app.New()

	myWindow := myApp.NewWindow("GUI Mockup")

	tree := widget.NewTree(firstTree, secondTree, thirdTree, fourthTree)

	rect := canvas.NewRectangle(color.RGBA{100, 100, 100, 255})

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			slog.Debug("New document")
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			slog.Debug("Display help")
		}),
	)
	content := container.NewBorder(toolbar, nil, tree, nil, rect)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func firstTree(id widget.TreeNodeID) []widget.TreeNodeID {
	switch id {
	case "":
		return []widget.TreeNodeID{"a", "b", "c"}
	case "a":
		return []widget.TreeNodeID{"a1", "a2"}
	}
	return []string{}
}
func secondTree(id widget.TreeNodeID) bool {
	return id == "" || id == "a"
}
func thirdTree(branch bool) fyne.CanvasObject {
	if branch {
		return widget.NewLabel("Branch template")
	}
	return widget.NewLabel("Leaf template")
}
func fourthTree(id widget.TreeNodeID, branch bool, o fyne.CanvasObject) {
	text := id
	if branch {
		text += " (branch)"
	}
	o.(*widget.Label).SetText(text)
}
