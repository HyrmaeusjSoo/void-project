package pkg

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// 获取根目录
func GetRootPath() string {
	dir := osExecutable()
	tmpEnv, _ := filepath.EvalSymlinks(os.TempDir())
	if strings.Contains(dir, tmpEnv) {
		return runtimeCaller()
	}
	return dir
}

// 编译后运行
func osExecutable() string {
	osexePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(osexePath, "../../../")
}

// 开发时运行
func runtimeCaller() string {
	_, current, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("获取根目录失败")
	}
	return filepath.Join(current, "../../")
}

// 根目录下的相对路径
//
//	传入绝对路径（全路径），返回与根目录的相对路径
func SubPath(dir string) string {
	relPath, err := filepath.Rel(GetRootPath(), dir)
	if err != nil {
		return dir
	}
	return relPath
}
