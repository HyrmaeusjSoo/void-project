package ansicli

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type (
	Style      int    // 样式
	Foreground int    // 前景色
	Background int    // 背景色
	Cursor     string // 光标控制
	Erase      string // 擦除
)

// 命令行界面类型
type CLI struct {
	Foreground Foreground // 前景色
	Background Background // 背景色
	Styles     []Style    // 样式
}

// 创建色彩
//
//	fg 前景色
//	bg 背景色
//	styles... 样式合集
func New(fg Foreground, bg Background, styles ...Style) *CLI {
	if fg == 0 {
		fg = FgDefault
	}
	if bg == 0 {
		bg = BgDefault
	}
	if len(styles) == 0 {
		styles = []Style{StyleReset}
	}

	return &CLI{fg, bg, styles}
}

// 创建默认样式和色彩
func NewDefault() *CLI {
	return &CLI{FgDefault, BgDefault, []Style{StyleReset}}
}

// 设置样式
//
//	styles... 样式合集
func (c *CLI) SetStyle(styles ...Style) *CLI {
	c.Styles = styles
	return c
}

// 设置前景色
//
//	fg 前景色
func (c *CLI) SetForeground(fg Foreground) *CLI {
	c.Foreground = fg
	return c
}

// 设置背景色
//
//	bg 背景色
func (c *CLI) SetBackground(bg Background) *CLI {
	c.Background = bg
	return c
}

// 追加样式
//
//	styles 样式
func (c *CLI) AppendStyle(styles ...Style) *CLI {
	c.Styles = slices.Concat(c.Styles, styles)
	return c
}

// 重置到默认样式和色彩
func (c *CLI) Reset() *CLI {
	c.Styles = []Style{StyleReset}
	c.Foreground = FgDefault
	c.Background = BgDefault
	return c
}

// 转换色彩码格式的string
func (c *CLI) String(a any) string {
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

// 打印
func (c *CLI) Print(a any) {
	fmt.Print(c.String(a))
}

// 打印一行
func (c *CLI) Println(a any) {
	fmt.Println(c.String(a))
}

// 格式化打印
func (c *CLI) Printf(format string, a ...any) {
	fmt.Printf(c.String(format), a...)
}

// 格式化打印一行
func (c *CLI) Printlnf(format string, a ...any) {
	c.Println(fmt.Sprintf(format, a...))
}

// 使用给定的前景色打印
//
//	fg 前景色
//	a 内容
func (c *CLI) PrintWithFg(fg Foreground, a any) {
	c.Foreground, fg = fg, c.Foreground
	c.Print(c.String(a))
	c.Foreground = fg
}

// 使用给定的前景色格式化打印
//
//	fg 前景色
//	format 格式字符串
//	a 内容
func (c *CLI) PrintfWithFg(fg Foreground, format string, a ...any) {
	c.Foreground, fg = fg, c.Foreground
	c.Print(fmt.Sprintf(format, a...))
	c.Foreground = fg
}

// 使用给定的前景色打印一行
//
//	fg 前景色
//	a 内容
func (c *CLI) PrintlnWithFg(fg Foreground, a any) {
	c.Foreground, fg = fg, c.Foreground
	c.Println(c.String(a))
	c.Foreground = fg
}

// 使用给定的前景色格式化打印一行
//
//	fg 前景色
//	format 格式字符串
//	a 内容
func (c *CLI) PrintlnfWithFg(fg Foreground, format string, a ...any) {
	c.Foreground, fg = fg, c.Foreground
	c.Println(fmt.Sprintf(format, a...))
	c.Foreground = fg
}

// 光标移动
//
//	cu 移动类型
//	n... 移动步数
func (c *CLI) Move(cu Cursor, n ...int) *CLI {
	if cu == CursorUP && len(n) > 1 {
		fmt.Printf("\x1b[%v;%v%s", n[0], n[1], cu)
	} else if len(n) > 0 {
		fmt.Printf("\x1b[%v%s", n[0], cu)
	} else {
		fmt.Printf("\x1b[%s", cu)
	}
	return c
}

// 光标移动m;n位置
func (c *CLI) CursorPosition(x, y int) *CLI {
	c.Move(CursorUP, x, y)
	return c
}

// 光标上移
func (c *CLI) CursorUp(step int) *CLI {
	c.Move(CursorUU, step)
	return c
}

// 光标下移
func (c *CLI) CursorDown(step int) *CLI {
	c.Move(CursorUD, step)
	return c
}

// 光标前/右移
func (c *CLI) CursorForward(step int) *CLI {
	c.Move(CursorUF, step)
	return c
}

// 光标后/左移
func (c *CLI) CursorBack(step int) *CLI {
	c.Move(CursorUB, step)
	return c
}

// 光标移到下一行(行首)
func (c *CLI) CursorNextLine(step int) *CLI {
	c.Move(CursorNL, step)
	return c
}

// 光标移到上一行(行首)
func (c *CLI) CursorPreviousLine(step int) *CLI {
	c.Move(CursorPL, step)
	return c
}

// 光标水平绝对移动(移动到N列)
func (c *CLI) CursorHorizontalAbsolute(step int) *CLI {
	c.Move(CursorHA, step)
	return c
}

// 保存光标
func (c *CLI) CursorSavePosition() *CLI {
	c.Move(CursorSCP)
	return c
}

// 恢复光标
func (c *CLI) CursorRestorePosition() *CLI {
	c.Move(CursorRCP)
	return c
}

// 隐藏光标
func (c *CLI) CursorHidden() *CLI {
	c.Move(CursorHidden)
	return c
}

// 显示光标
func (c *CLI) CursorDisplay() *CLI {
	c.Move(CursorDisplay)
	return c
}

// 擦除光标
//
//	e 擦除类型
func (c *CLI) Erase(e Erase) *CLI {
	fmt.Printf("\x1b[%s", e)
	return c
}

// 清屏
func (c *CLI) EraseDisplay() *CLI {
	c.Erase(EraseDisplay)
	return c
}

// 擦除屏幕开头到光标位置
func (c *CLI) EraseDisplayBeforeCursor() *CLI {
	c.Erase(EraseDBefore)
	return c
}

// 擦除光标位置到屏幕结尾
func (c *CLI) EraseDisplayAfterCursor() *CLI {
	c.Erase(EraseDAfter)
	return c
}

// 清屏并移动光标到左上角
func (c *CLI) EraseDisplayMoveCursor() *CLI {
	c.Erase(EraseDMoveCursor)
	return c
}

// 擦除整行
func (c *CLI) EraseLine() *CLI {
	c.Erase(EraseLine)
	return c
}

// 擦除光标位置到行尾
func (c *CLI) EraseLineAfterCursor() *CLI {
	c.Erase(EraseLAfter)
	return c
}

// 擦除光标位置到行首
func (c *CLI) EraseLineBeforeCursor() *CLI {
	c.Erase(EraseLBefore)
	return c
}
