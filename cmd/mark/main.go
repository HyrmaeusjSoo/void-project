package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"void-project/pkg"
	"void-project/pkg/types/composite"
)

const (
	elegansSpace  = "     "
	elegansLine   = "    │"
	elegansLast   = "    └── "
	elegansMiddle = "    ├── "
)

var (
	levels      []bool
	rootpath    = pkg.GetRootPath()
	osPathSept  = string(os.PathSeparator)
	excludeDirs = []string{
		".git",
		".vscode",
		"web" + osPathSept + "upload" + osPathSept,
	}
)

// 添加标记
// g => 生成和替换README文档中目录树部分、version版本号部分。
// g => 往源代码里添加头部标记注释。注释内容读取的是同目录c.mark和html.mark插入到对应的Go，JS，HTML代码文件头部。
func main() {
	fmt.Println("g=Generate README, c=Insert comment")
	var input string
	fmt.Scanln(&input)
	input = strings.ToLower(input)
	if input == "g" {
		GenReadme()
	} else if input == "c" {
		fmt.Println("a=insert, d=delete")
		fmt.Scanln(&input)
		input = strings.ToLower(input)
		if input != "a" && input != "d" {
			fmt.Println("nothing!")
			return
		}
		total, modified := InsertComment(pkg.IfElse(input == "d", false, true))
		fmt.Println(modified, "/", total)
	} else {
		fmt.Println("nothing!")
		return
	}
}

// 插入文件头注释
func InsertComment(isSet bool) (total, modified int) {
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
		if ext := filepath.Ext(path); ext == ".go" || ext == ".js" || ext == ".ts" {
			writeFile(path, normal, isSet)
			modified++
		} else if ext == ".html" || ext == ".htm" || ext == ".tmpl" || ext == ".vue" {
			writeFile(path, html, isSet)
			modified++
		}
		total++
		return nil
	})
	if err != nil {
		panic(err)
	}
	return
}

// 写入文件
func writeFile(path string, mark []byte, isSet bool) {
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 读内容
	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	// 删旧注释
	if slices.Equal(content[:6], mark[:6]) {
		if pos := composite.SearchSubSlice(content, mark[len(mark)-8:]); pos > 0 {
			content = content[pos+1:]
		}
	}
	// 加新注释
	if isSet {
		content = slices.Concat(mark, content)
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

// 生成README 目录结构、版本号
func GenReadme() {
	exclCfg, err := os.ReadFile(rootpath + "/.git/info/exclude")
	if err != nil {
		panic(err)
	}
	exclNames := strings.SplitSeq(string(exclCfg), "\n")
	for v := range exclNames {
		if v == "" || strings.Contains(v, "#") {
			continue
		}
		path := strings.Replace(v, "\\", osPathSept, -1)
		path = strings.Replace(v, "/", osPathSept, -1)
		excludeDirs = append(excludeDirs, path)
	}

	levels = make([]bool, 72)
	tree := GenDirectoryTree(rootpath, 0)
	// tree := GenDirectoryTree2()
	structureTree := []byte(tree)

	file, err := os.OpenFile(rootpath+"/README.md", os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 原文件内容
	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	readme := make([]byte, 0, len(content))
	lastPoint := 0
	// 版本号部分
	if begin := composite.SearchSubSlice(content, []byte(`<img src="https://img.shields.io/badge/version-`)); begin > 0 {
		endVersion := []byte(`-05e5a5">`)
		if end := composite.SearchSubSlice(content, endVersion); end > 0 {
			version, err := os.ReadFile(rootpath + "/VERSION")
			if err != nil {
				panic(err)
			}
			readme = slices.Concat(content[:begin+1:begin+1], version)
			lastPoint = end - len(endVersion) + 1
		}
	}
	// 目录结构部分
	var (
		beginTree = []byte("```\r\n──────────────────begin──────────────────\r\nvoid-project\r\n")
		endTree   = []byte("\r\n───────────────────end───────────────────\r\n```\r\n")
	)
	if pos := composite.SearchSubSlice(content, beginTree); pos > 0 {
		readme = slices.Concat(readme, content[lastPoint:pos+1])
		readme = slices.Concat(readme, structureTree)
	}
	if pos := composite.SearchSubSlice(content, endTree); pos > 0 {
		lastPoint = pos - len(endTree) + 2
	}
	// 拼接剩下部分
	readme = slices.Concat(readme, content[lastPoint+1:])

	err = file.Truncate(0)
	if err != nil {
		panic(err)
	}
	_, err = file.WriteAt(readme, 0)
	if err != nil {
		panic(err)
	}

	fmt.Println(tree)
}

// 读取项目文件夹，生成目录树
func GenDirectoryTree(dir string, lv int) string {
	levels[lv] = true
	entries, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	directory := make([]string, 0, 72)
	for _, entry := range entries {
		current := filepath.Join(dir, entry.Name())
		if !entry.IsDir() || isExcludeDir(current) {
			continue
		}

		directory = append(directory, entry.Name())
	}
	var tree strings.Builder
	for i, v := range directory {
		current := filepath.Join(dir, v)
		isLast := i == len(directory)-1
		levels[lv] = !isLast
		if strings.Count(current, osPathSept)-strings.Count(rootpath, osPathSept) > 0 {
			for j := range lv {
				tree.WriteString(pkg.IfElse(levels[j], elegansLine, elegansSpace))
			}
			tree.WriteString(pkg.IfElse(isLast, elegansLast, elegansMiddle))
		}
		tree.WriteString(v)
		tree.WriteString("\r\n")

		tree.WriteString(GenDirectoryTree(current, lv+1))
	}
	return tree.String()
}

func isExcludeDir(dirName string) bool {
	for _, exclude := range excludeDirs {
		if strings.Contains(dirName+osPathSept, "void-project"+osPathSept+exclude) {
			return true
		}
	}
	return false
}

func GenDirectoryTree2() string {
	dirTree := ""
	err := filepath.Walk(rootpath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && !strings.Contains(path, ".git") && !strings.Contains(path, ".vscode") {
			depth := strings.Count(path, osPathSept) - strings.Count(rootpath, osPathSept)
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
	return dirTree
}
