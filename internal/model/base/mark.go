package base

import (
	"fmt"
)

func NewMark() []string {
	return []string{
		string([]rune{9484, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9488}),
		string([]rune{9474, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 118, 111, 105, 100, 45, 112, 114, 111, 106, 101, 99, 116}),
		string([]rune{9500, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9508}),
		string([]rune{9474, 32, 83, 363, 32, 83, 104, 275, 110, 103, 120, 476, 39, 115, 32, 118, 111, 105, 100, 45, 112, 114, 111, 106, 101, 99, 116, 32, 105, 115, 32, 97, 32, 119, 101, 98, 32, 97, 112, 112, 108, 105, 99, 97, 116, 105, 111, 110, 32, 97, 114, 99, 104, 105, 116, 101, 99, 116, 117, 114, 101, 32, 100, 101, 118, 101, 108, 111, 112, 101, 100, 32, 105, 110, 32, 71, 111, 46}),
		string([]rune{9500, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9508}),
		string([]rune{9474, 32, 80, 114, 111, 106, 101, 99, 116, 32, 114, 101, 112, 111, 115, 105, 116, 111, 114, 121, 32, 108, 105, 110, 107, 58, 32, 104, 116, 116, 112, 115, 58, 47, 47, 103, 105, 116, 104, 117, 98, 46, 99, 111, 109, 47, 72, 121, 114, 109, 97, 101, 117, 115, 106, 83, 111, 111, 47, 118, 111, 105, 100, 45, 112, 114, 111, 106, 101, 99, 116}),
		string([]rune{9500, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9508}),
		string([]rune{9474, 32, 38134, 27827, 31995, 32, 127756, 9883, 65039, 129516, 129482, 128302, 128481, 65039, 10017, 65039, 127966, 65039, 127752, 127918, 129694, 129767, 32, 82, 101, 113, 117, 101, 115, 116, 115, 46, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 50, 48, 50, 51}),
		string([]rune{9500, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9508}),
		string([]rune{9474, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 8212, 8212, 8212, 8212, 8212, 8212, 32, 72, 121, 114, 109, 97, 101, 117, 115, 106, 32, 33487}),
		string([]rune{9492, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9472, 9496}),
	}
}

func Echo() {
	title := `
    ┌───────────────────────────────────────────────────────────────────────────────────────────┐
    │                                       void-project
    ├───────────────────────────────────────────────────────────────────────────────────────────┤
    │ Sū Shēngxǜ's void-project is a web application architecture developed in Go.
    ├───────────────────────────────────────────────────────────────────────────────────────────┤
    │ Project repository link: https://github.com/HyrmaeusjSoo/void-project
    ├───────────────────────────────────────────────────────────────────────────────────────────┤
    │ 银河系 🌌⚛️🧬🧊🔮🗡️✡️🏞️🌈🎮🪞🫧 Requests.                                      2023
    ├───────────────────────────────────────────────────────────────────────────────────────────┤
    │                                                                   —————— Hyrmaeusj 苏
    └───────────────────────────────────────────────────────────────────────────────────────────┘
    `
	fmt.Println(title)
	fmt.Println([]rune("┌───────────────────────────────────────────────────────────────────────────────────────────┐"))
	fmt.Println([]rune("│                                       void-project"))
	fmt.Println([]rune("├───────────────────────────────────────────────────────────────────────────────────────────┤"))
	fmt.Println([]rune("│ Sū Shēngxǜ's void-project is a web application architecture developed in Go."))
	fmt.Println([]rune("├───────────────────────────────────────────────────────────────────────────────────────────┤"))
	fmt.Println([]rune("│ Project repository link: https://github.com/HyrmaeusjSoo/void-project"))
	fmt.Println([]rune("├───────────────────────────────────────────────────────────────────────────────────────────┤"))
	fmt.Println([]rune("│ 银河系 🌌⚛️🧬🧊🔮🗡️✡️🏞️🌈🎮🪞🫧 Requests.                                      2023"))
	fmt.Println([]rune("├───────────────────────────────────────────────────────────────────────────────────────────┤"))
	fmt.Println([]rune("│                                                                   —————— Hyrmaeusj 苏"))
	fmt.Println([]rune("└───────────────────────────────────────────────────────────────────────────────────────────┘"))
}