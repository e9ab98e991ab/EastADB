package module

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type MsgType struct {
	Msg string
}

//退出提示
func MessageDlg() bool {
	return vcl.MessageDlg("是否退出？", types.MtConfirmation, types.MbYes, types.MbNo) == types.IdYes
}

func (ms MsgType) MessageDlg() bool {
	return vcl.MessageDlg(ms.Msg, types.MtConfirmation, types.MbYes, types.MbNo) == types.IdYes
}
