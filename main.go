package main

import (
	"github.com/go-ole/go-ole"
	"os"
	"github.com/go-ole/go-ole/oleutil"
	"flag"
	"fmt"
	"strings"
)

var path string
var name string

func init() {
	pathAddr := flag.String("p", "./", "可执行文件位置")
	nameAddr := flag.String("n", "yc", "软件名")
	flag.Parse()
	path = *pathAddr
	name = *nameAddr
}

func main() {
	ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED|ole.COINIT_SPEED_OVER_MEMORY)
	oleShellObject, err := oleutil.CreateObject("WScript.Shell")
	if err != nil {
		fmt.Println(err)
	}
	defer oleShellObject.Release()
	wshell, err := oleShellObject.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		fmt.Println(err)
	}
	defer wshell.Release()

	autoStartPath := os.Getenv("appdata") + "/Microsoft/Windows/Start Menu/Programs/Startup/" + name + ".lnk"
	cs, err := oleutil.CallMethod(wshell, "CreateShortcut", autoStartPath)
	if err != nil {
		fmt.Println(err)
	}
	idispatch := cs.ToIDispatch()

	if err != nil {
		fmt.Println(err)
	}

	pathLen := len(path)
	if strings.LastIndex(path, "/") != pathLen-1 {
		path = path + "/"
	}

	oleutil.PutProperty(idispatch, "TargetPath", path+name)
	_, err = oleutil.CallMethod(idispatch, "Save")
	fmt.Println("success")
}
