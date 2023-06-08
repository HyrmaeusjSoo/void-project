package main

import (
	"chat/pkg"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("Which one?  [g]=Generate Readme  [c]=Insert omment")
	var input string
	fmt.Scanln(&input)
	input = strings.ToLower(input)
	if input == "g" {
		GenReadme()
	} else if input == "c" {
		fmt.Println("Which one?  [a]=insert comment  [d]=delete comment")
		fmt.Scanln(&input)
		input = strings.ToLower(input)
		if input != "a" && input != "d" {
			fmt.Println("nothing!")
			return
		}
		InsertComment(pkg.IfElse(input == "d", false, true))
	} else {
		fmt.Println("nothing!")
		return
	}
}

// 插入文件头注释
func InsertComment(isSet bool) {
	rootpath := pkg.GetRootPath()

	// 读注释文件
	cMark, err := os.Open(rootpath + "/cmd/mark/c.mark")
	if err != nil {
		panic(err)
	}
	defer cMark.Close()
	normal, err := io.ReadAll(cMark)
	if err != nil {
		panic(err)
	}
	htmlMark, err := os.Open(rootpath + "/cmd/mark/html.mark")
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
	err = filepath.Walk(rootpath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if ext := filepath.Ext(path); ext == ".go" || ext == ".js" {
			writeFile(path, normal, isSet)
		} else if ext == ".html" || ext == ".htm" || ext == ".tmpl" {
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

// 生成README 目录结构
func GenReadme() {
	dirTree := ""
	err := filepath.Walk(pkg.GetRootPath(), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && !strings.Contains(path, ".git") && !strings.Contains(path, ".vscode") {
			depth := strings.Count(path, string(os.PathSeparator)) - strings.Count(pkg.GetRootPath(), string(os.PathSeparator))
			if depth == 0 {
				dirTree += info.Name() + "\r\n"
			} else {
				dirTree += strings.Repeat("    │", depth-1) + "    ├── " + info.Name() + "\r\n"
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	readme := append([]byte("# Void_Project\r\n```\r\n"), []byte(dirTree+"\r\n```")...)
	file, err := os.OpenFile(pkg.GetRootPath()+"/README.md", os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = file.Truncate(0)
	if err != nil {
		panic(err)
	}
	_, err = file.WriteAt(readme, 0)
	if err != nil {
		panic(err)
	}

	fmt.Println(dirTree)
}
