package main

import (
	"github.com/go-ole/go-ole"
	"os"
	"github.com/go-ole/go-ole/oleutil"
	"flag"
	"fmt"
)

var path string
var name string
var start bool
var autoStartPath string

func init() {
	pathAddr := flag.String("p", "./", "可执行文件位置")
	nameAddr := flag.String("n", "yc", "快捷方式名称")
	startAddr := flag.Bool("s", true, "取消开机自启动")
	flag.Parse()
	path = *pathAddr
	name = *nameAddr
	start = *startAddr
	autoStartPath = os.Getenv("appdata") + "/Microsoft/Windows/Start Menu/Programs/Startup/" + name + ".lnk"
}

func main() {
	if start {
		ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED|ole.COINIT_SPEED_OVER_MEMORY)
		oleShellObject, err := oleutil.CreateObject("WScript.Shell")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer oleShellObject.Release()
		wshell, err := oleShellObject.QueryInterface(ole.IID_IDispatch)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer wshell.Release()
		cs, err := oleutil.CallMethod(wshell, "CreateShortcut", autoStartPath)
		if err != nil {
			fmt.Println(err)
			return
		}
		idispatch := cs.ToIDispatch()

		fmt.Println(path)
		oleutil.PutProperty(idispatch, "TargetPath", path)
		_, err = oleutil.CallMethod(idispatch, "Save")
	} else {
		err := os.Remove(autoStartPath)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("success")
}
