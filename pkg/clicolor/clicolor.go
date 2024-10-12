package clicolor

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type (
	Style      int // 终端样式
	Foreground int // 终端前景色
	Background int // 终端背景色
)

const (
	StyleReset      Style = iota // 默认样式
	StyleBold                    // 粗体/高亮
	StyleDark                    // 暗
	StyleItalic                  // 斜体
	StyleUnderline               // 下划线
	StyleBlinkSlow               // 慢闪烁
	StyleBlinkRapid              // 快闪烁
	StyleInvert                  // 反转
	StyleHidden                  // 隐藏
	StyleCross                   // 删除线
)

const (
	FgBlack        Foreground = iota + 30 // 黑色前景
	FgRed                                 // 红色前景
	FgGreen                               // 绿色前景
	FgYellow                              // 黄色前景
	FgBlue                                // 蓝色前景
	FgMagenta                             // 品红前景
	FgCyan                                // 青色前景
	FgLightGrey                           // 浅灰前景
	FgDefault      Foreground = iota + 31 // 默认前景
	FgDarkGrey     Foreground = iota + 81 // 深灰前景
	FgLightRed                            // 浅红前景
	FgLightGreen                          // 浅绿前景
	FgLightYellow                         // 浅黄前景
	FgLightBlue                           // 浅蓝前景
	FgLightMagenta                        // 浅品红前景
	FgLightCyan                           // 浅青色前景
	FgWhite                               // 白色前景
)

const (
	BgBlack        Background = iota + 40 // 黑色背景
	BgRed                                 // 红色背景
	BgGreen                               // 绿色背景
	BgYellow                              // 黄色背景
	BgBlue                                // 蓝色背景
	BgMagenta                             // 品红背景
	BgCyan                                // 青色背景
	BgLightGrey                           // 浅灰背景
	BgDefault      Background = iota + 41 // 默认背景
	BgDarkGrey     Background = iota + 91 // 深灰背景
	BgLightRed                            // 浅红背景
	BgLightGreen                          // 浅绿背景
	BgLightYellow                         // 浅黄背景
	BgLightBlue                           // 浅蓝背景
	BgLightMagenta                        // 浅品红背景
	BgLightCyan                           // 浅青色背景
	BgWhite                               // 白色背景
)

// 色彩终端类型
type Color struct {
	Foreground Foreground // 前景色
	Background Background // 背景色
	Styles     []Style    // 样式
}

// 创建色彩终端
//
//	fg 前景色
//	bg 背景色
//	styles... 样式合集
func New(fg Foreground, bg Background, styles ...Style) *Color {
	if fg == 0 {
		fg = FgDefault
	}
	if bg == 0 {
		bg = BgDefault
	}
	if len(styles) > 0 {
		styles = []Style{StyleReset}
	}

	return &Color{fg, bg, styles}
}

// 创建默认样式和色彩的终端
func NewDefault() *Color {
	return &Color{FgDefault, BgDefault, []Style{StyleReset}}
}

// 设置样式
//
//	styles... 样式合集
func (c *Color) SetStyle(styles ...Style) *Color {
	c.Styles = styles
	return c
}

// 设置终端前景色
//
//	fg 前景色
func (c *Color) SetForeground(fg Foreground) *Color {
	c.Foreground = fg
	return c
}

// 设置终端背景色
//
//	bg 背景色
func (c *Color) SetBackground(bg Background) *Color {
	c.Background = bg
	return c
}

// 追加终端样式
//
//	styles 样式
func (c *Color) AppendStyle(styles ...Style) *Color {
	c.Styles = slices.Concat(c.Styles, styles)
	return c
}

// 重置终端到默认样式和色彩
func (c *Color) Reset() *Color {
	c.Styles = []Style{StyleReset}
	c.Foreground = FgDefault
	c.Background = BgDefault
	return c
}

// 转换色彩码格式的string
func (c *Color) String(a any) string {
	var str = strings.Builder{}
	str.WriteString("\x1b[")
	for _, style := range c.Styles {
		str.WriteString(strconv.Itoa(int(style)))
		str.WriteString(";")
	}
	str.WriteString(strconv.Itoa(int(c.Foreground)))
	str.WriteString(";")
	str.WriteString(strconv.Itoa(int(c.Background)))
	str.WriteString("m")
	str.WriteString(fmt.Sprint(a))
	str.WriteString("\x1b[0m")
	return str.String()
}

// 终端打印一行
func (c *Color) Println(a any) {
	fmt.Println(c.String(a))
}

// 终端打印
func (c *Color) Print(a any) {
	fmt.Print(c.String(a))
}
