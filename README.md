# windows-autostart
windows版自启动你的app程序

---
[![stable](https://img.shields.io/badge/stable-stable-green.svg)](https://github.com/snail007/goproxy/)

### 使用前须知
 - [作用](#作用)
 - [下载](#下载)
 - [参数](#参数)
 - [依赖包](#依赖包)

### 作用
1、自启动你的window应用
 
### 下载
[下载地址](https://github.com/yincongcyincong/windows-autostart/releases)  

### 参数
-p：可执行文件位置，一定要是绝对路径，不然快捷方式找不到路径    
-n：快捷方式名称  
-s：是否开启自启动自启动  
开启自启动例子：./windows-autostart.exe -n=xxx -p=/your/app/path/xxx.exe  
关闭自启动例子：./windows-autostart.exe -n=xxx -s=false

### 依赖包
[github.com/go-ole/go-ole](https://github.com/go-ole/go-ole)，go包生成快捷方式    
