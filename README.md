# goLearning

this is the self-learning project for Go

## windows下安装和环境配置
1. 从这里下载Windows版本 https://golang.org/dl/
2. 直接装到默认的路径 C:/Go

## VS code 
1. 配置：https://code.visualstudio.com/docs/languages/go
2. Debug: https://github.com/Microsoft/vscode-go/wiki/Debugging-Go-code-using-VS-Code

## Go 语言指南
https://tour.go-zh.org/flowcontrol/6   


## 常见问题
1. 安装插件时提示没有权限
   1. Ctl + Shift + P打开command Palette
   2. 输入"User settings"
   3. 输入 "gopath", 找到 go:gopath， 点击edit in setting.json
   4. 添加一项 "go.toolsGopath": "C:\\Go\\vstools"， 其中的目录可以放在任何其它有权限的地方
2. debug 一直打不开
   1. vscode 以管理员模式打开，重启一次
   2. cmd以管理员模式打开，重启一次