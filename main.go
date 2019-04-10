package main

import (
	"EasyADB/context"
	"EasyADB/module"
	"EasyADB/ui"
	"github.com/ying32/govcl/vcl"
)

func setupUI() {
	vcl.Application.Initialize()
	mainForm := vcl.Application.CreateForm()
	mainForm.SetCaption(context.WindowTitle)
	mainForm.EnabledMaximize(false)
	mainForm.ScreenCenter()
	mainForm.SetOnCloseQuery(func(Sender vcl.IObject, CanClose *bool) {
		*CanClose = module.MessageDlg()
	})
	//菜单区域
	ui.FileMenu(mainForm)

	vcl.Application.Run()
}

func main() {
	setupUI()
}
