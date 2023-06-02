package pkg

import (
	"path/filepath"
	"runtime"
)

func GetRootPath() string {
	_, current, _, ok := runtime.Caller(0)
	if !ok {
		panic("获取根目录失败")
	}
	return filepath.Join(current, "../../")
}
