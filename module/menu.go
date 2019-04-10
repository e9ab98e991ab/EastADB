package module

import (
	"github.com/ying32/govcl/vcl"
)

var on OnMenuItemClickListener

/***
 * 菜单item封装
 */

func MenuItem(mainForm vcl.IComponent, itemName string, position int) *vcl.TMenuItem {

	item := vcl.NewMenuItem(mainForm)
	item.SetCaption(itemName)
	mainForm.SetTag(position)
	item.SetOnClick(func(sender vcl.IObject) {
		on.OnMenuItemClick(item, position)
	})
	return item
}

type OnMenuItemClickListener interface {
	OnMenuItemClick(sender *vcl.TMenuItem, position int)
}
