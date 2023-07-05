package main

import (
	"fmt"
	"void-project/pkg"
)

func main() {
	// logger.LogDebug("aaaaaaaaaaaaaaaaa")
	// logger.LogInfo("bbbbbbbbbbbbbbbbbbbb")
	// logger.LogWarn("cccccccccccccccccccc")
	// logger.LogError("ddddddddddddddddd")
	//logger.ClearLog(logger.RemoveAll)
	a := pkg.GetRootPath()
	fmt.Println(a)
	b := ""
	fmt.Scanln(&b)
	fmt.Println(b)
}
