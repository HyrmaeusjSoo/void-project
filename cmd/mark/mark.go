package main

import (
	"chat/pkg"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	fmt.Println("Which one?  [a]=set comment  [d]=delete comment:")
	var input string
	fmt.Scanln(&input)
	input = strings.ToLower(input)
	if input != "a" && input != "d" {
		fmt.Println("nothing!")
		return
	}
	SetComment(pkg.IfElse(input == "d", false, true))
}

func SetComment(isSet bool) {
	_, currentFile, _, _ := runtime.Caller(0)

	// 读注释文件
	cMark, err := os.Open(filepath.Dir(currentFile) + "/c.mark")
	if err != nil {
		panic(err)
	}
	defer cMark.Close()
	normal, err := io.ReadAll(cMark)
	if err != nil {
		panic(err)
	}
	htmlMark, err := os.Open(filepath.Dir(currentFile) + "/html.mark")
	if err != nil {
		panic(err)
	}
	defer htmlMark.Close()
	html, err := io.ReadAll(htmlMark)
	if err != nil {
		panic(err)
	}
	if len(normal) == 0 || len(html) == 0 {
		panic("无注释范本")
	}

	// 遍历路径
	err = filepath.Walk(filepath.Join(currentFile, "../../../"), func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".go" || filepath.Ext(path) == ".js" {
			writeFile(path, normal, isSet)
		} else if filepath.Ext(path) == ".html" || filepath.Ext(path) == ".htm" {
			writeFile(path, html, isSet)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func writeFile(path string, mark []byte, isSet bool) {
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// 读旧内容
	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	// 删旧注释
	if pkg.CompareSlice(content[:6], mark[:6]) {
		if pos := pkg.FindSliceInSlice(content, mark[len(mark)-8:]); pos > 0 {
			content = content[pos+1:]
		}
	}

	// 加新注释
	if isSet {
		content = append(mark, content...)
	}
	// 覆写到文件
	err = file.Truncate(0)
	if err != nil {
		panic(err)
	}
	_, err = file.WriteAt(content, 0)
	if err != nil {
		panic(err)
	}
}
