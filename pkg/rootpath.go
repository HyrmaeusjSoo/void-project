package pkg

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// 获取根目录
func GetRootPath() string {
	// 获取执行目录
	dir := osExecutable()

	// 判断是否为开发"go run *"执行
	//获取环境变量%TEMP%
	tmpEnv, _ := filepath.EvalSymlinks(os.TempDir())
	//判断‘执行目录’是否在环境变量%TEMP%中
	if strings.Contains(dir, tmpEnv) {
		return runtimeCaller() //返回运行时路径
	}

	//获取go env GOCACHE
	cmd := exec.Command("go", "env", "GOCACHE")
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	GOCACHE := string(out)
	GOCACHE = GOCACHE[:len(GOCACHE)-1] //去掉结尾换行
	// 判断‘执行目录’是否在go env GOCACHE
	if strings.Contains(dir, GOCACHE) {
		return runtimeCaller() //返回运行时路径
	}
	return dir
}

// 执行目录
func osExecutable() string {
	OSExePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(OSExePath, "../../../")
}

// 运行时目录
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
