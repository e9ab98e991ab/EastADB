package main

import (
	"EasyADB/context"
	"github.com/ying32/govcl/vcl"
)

func setupUI() {
	vcl.Application.Initialize()
	mainForm := vcl.Application.CreateForm()
	mainForm.SetCaption(context.WindowTitle)
	mainForm.EnabledMaximize(false)
	mainForm.ScreenCenter()

	vcl.Application.Run()
}

func main() {
	setupUI()
}