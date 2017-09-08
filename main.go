package main

import (
	"fmt"
	"os"
	"app/container"
	"app/config"
)

func main() {
	container.Init()
	di := container.GetInstance()

	var params []string
	if len(os.Args) <= 1 {
		fmt.Println("请输入方法名")
		return
	}

	TaskMap := config.GetTaskMap(di)

	for i := 1; i < len(os.Args); i++ {
		params = append(params, os.Args[i])
	}

	if TaskMap[params[0]] == nil {
		fmt.Println("找不到此命令")
		return
	}

	err := TaskMap[params[0]].Handle()

	if err != nil {
		fmt.Println(err.Error())
	}
}