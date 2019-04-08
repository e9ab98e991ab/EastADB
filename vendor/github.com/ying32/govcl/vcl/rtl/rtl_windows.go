package rtl

import (
	"github.com/ying32/govcl/vcl/api"
	"github.com/ying32/govcl/vcl/win"
)

// MainInstance EXE自身的实例
func MainInstance() uintptr {
	return win.GetSelfModuleHandle()
}

// SetReportMemoryLeaksOnShutdown 程序结束时报告内存泄露，总有2-4字节的未知泄露位置
func SetReportMemoryLeaksOnShutdown(v bool) {
	api.DSetReportMemoryLeaksOnShutdown(v)
}

// Shortcut

//CreateURLShortCut
/*
   CN: 创建一个url的快捷方式
   EN: Create a shortcut to a URL

   rtl.CreateURLShortCut("C:\\aaa\\bbb\\", "govcl", "https://github.com/ying32/govcl")
*/
func CreateURLShortCut(aDestPath, aShortCutName, aURL string) {
	api.DCreateURLShortCut(aDestPath, aShortCutName, aURL)
}

// CreateShortCut
/*
   CN: 创建一个快捷方式
   EN: Create a shortcut

   1. rtl.CreateShortCut("C:\\Users\\administrator\\Desktop\\", "govcl", os.Args[0], "", "", "")
   2. rtl.CreateShortCut("C:\\Users\\administrator\\Desktop\\", "govcl", os.Args[0], "", "Description text", "-a -b")
*/
func CreateShortCut(aDestPath, aShortCutName, aSrcFileName, aIconFileName, aDescription, aCmdArgs string) bool {
	return api.DCreateShortCut(aDestPath, aShortCutName, aSrcFileName, aIconFileName, aDescription, aCmdArgs)
}
