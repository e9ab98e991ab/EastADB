package ui

import (
	"EasyADB/module"
	"fmt"
	"github.com/ying32/govcl/vcl"
)

var (
	mainMenu *vcl.TMainMenu
)

/***
 * 文件菜单模块
 */
func FileMenu(mainForm *vcl.TForm) {

	mainMenu = vcl.NewMainMenu(mainForm)
	item := vcl.NewMenuItem(mainForm)
	item.SetCaption("File(&F)")
	mainMenu.Items().Add(item)

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

func (i ItemClick) OnMenuItemClick(sender *vcl.TMenuItem, position int) {
	switch position {
	case 0:
		vcl.ShowMessage("Click")
		break
	case 1:
		vcl.ShowMessage("Click1")
		break

	}
}

type ItemClick struct {
	module.OnMenuItemClickListener
}

type Rect struct {
	width,
	length float64
}

func main() {
	var rect Rect
	rect.width = 100
	rect.length = 200
	fmt.Println(rect.width * rect.length)
}
