package pkg

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func GetRootPath() string {
	dir := getRootpathOSExecutable()
	tmpEnv, _ := filepath.EvalSymlinks(os.TempDir())
	if strings.Contains(dir, tmpEnv) {
		return getRootpathSystemCaller()
	}
	return dir
}

func getRootpathOSExecutable() string {
	osexePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(osexePath, "../../")
}

func getRootpathSystemCaller() string {
	_, current, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("获取根目录失败")
	}
	return filepath.Join(current, "../../")
}
