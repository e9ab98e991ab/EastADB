package ui

import (
	"EasyADB/module"
	"github.com/ying32/govcl/vcl"
)

var (
	MainMenu *vcl.TMainMenu
	MF       *vcl.TForm
)

/***
 * 文件菜单模块
 */
func FileMenu(mainForm *vcl.TForm) {
	MF = mainForm
	MainMenu = vcl.NewMainMenu(mainForm)
	item := vcl.NewMenuItem(mainForm)
	item.SetCaption("File(&F)")
	MainMenu.Items().Add(item)

	item.Add(module.MenuItem(mainForm, "Setting", 0))
	item.Add(module.MenuItem(mainForm, "Exit", 1))

	//item.Items(0).SetOnClick(func(sender vcl.IObject) {
	//	vcl.ShowMessage("Click")
	//})
	//
	//item.Items(1).SetOnClick(func(sender vcl.IObject) {
	//	mainForm.Close()
	//})

}
