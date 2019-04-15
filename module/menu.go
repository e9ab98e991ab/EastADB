package module

import (
	"github.com/ying32/govcl/vcl"
)

var (
	ab AB
)

/***
 * 菜单item封装
 */

func init() {

}
func MenuItem(mainForm *vcl.TForm, itemName string, position int) *vcl.TMenuItem {

	item := vcl.NewMenuItem(mainForm)
	item.SetCaption(itemName)
	mainForm.SetTag(position)
	item.SetOnClick(func(sender vcl.IObject) {
		ab.OnMenuItemClick(item, position, mainForm)

	})
	return item
}

type AB struct {
}

func (AB) OnMenuItemClick(sender vcl.IObject, position int, mainForm *vcl.TForm) {
	switch position {
	case 0:
		vcl.ShowMessage("Click")
		break
	case 1:
		mainForm.Close()
		break

	}
}
