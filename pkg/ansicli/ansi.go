package ansicli

const (
	StyleReset      Style = iota // 默认样式
	StyleBold                    // 粗体/高亮
	StyleDark                    // 暗/虚化
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
	FgLightGrey                           // 浅灰/白色前景
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

const (
	CursorUP      Cursor = "H"    // 光标移动x;y位置
	CursorUU      Cursor = "A"    // 光标上移
	CursorUD      Cursor = "B"    // 光标下移
	CursorUF      Cursor = "C"    // 光标前/右移
	CursorUB      Cursor = "D"    // 光标后/左移
	CursorNL      Cursor = "E"    // 光标移到下一行
	CursorPL      Cursor = "F"    // 光标移到上一行
	CursorHA      Cursor = "G"    // 光标水平移动
	CursorSCP     Cursor = "s"    // 光标保存
	CursorRCP     Cursor = "u"    // 恢复光标的位置
	CursorHidden  Cursor = "?25l" // 隐藏光标
	CursorDisplay Cursor = "?25h" // 显示光标
)

const (
	EraseDisplay     Erase = "J"  // 清屏
	EraseDBefore     Erase = "0J" // 擦除屏幕开头到光标位置
	EraseDAfter      Erase = "1J" // 擦除光标位置到屏幕结尾
	EraseDMoveCursor Erase = "2J" // 清屏并移动光标到左上角
	EraseLine        Erase = "2K" // 擦除整行
	EraseLAfter      Erase = "0K" // 擦除光标位置到行尾
	EraseLBefore     Erase = "1K" // 擦除光标位置到行首
)
