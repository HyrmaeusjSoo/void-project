package pkg

import (
	"errors"
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
	exePath, err := osExecutable()
	if err != nil {
		log.Fatal(err)
	}

	// 获取运行时当前目录
	currentPath, err := runtimeCaller()
	if err != nil {
		log.Fatal(err)
	}

	// 判断是否为开发"go run *"执行
	//获取系统环境变量%TEMP%
	tmpEnv, _ := filepath.EvalSymlinks(os.TempDir())
	//判断‘执行目录’是否在环境变量%TEMP%中
	if strings.Contains(exePath, tmpEnv) {
		return currentPath //返回运行时路径
	}

	//获取go env GOCACHE
	cmd := exec.Command("go", "env", "GOCACHE")
	if out, err := cmd.Output(); err == nil {
		GOCACHE := strings.TrimSpace(string(out)) //去掉结尾换行
		// 判断‘执行目录’是否在go env GOCACHE
		if strings.Contains(exePath, GOCACHE) {
			return currentPath //返回运行时路径
		}
	}

	//获取go env GOTMPDIR
	cmd = exec.Command("go", "env", "GOTMPDIR")
	if out, err := cmd.Output(); err == nil {
		GOTMPDIR := strings.TrimSpace(string(out)) //去掉结尾换行
		// 判断‘执行目录’是否在go env GOTMPDIR
		if strings.Contains(exePath, GOTMPDIR) {
			return currentPath //返回运行时路径
		}
	}

	return exePath
}

// 执行目录
func osExecutable() (string, error) {
	OSExePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Join(filepath.Dir(OSExePath), "../.."), nil
}

// 运行时目录
func runtimeCaller() (string, error) {
	_, current, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("获取根目录失败")
	}
	return filepath.Join(filepath.Dir(current), ".."), nil
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
