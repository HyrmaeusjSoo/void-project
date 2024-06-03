/*
反射包

	实在不愿意用反射！它影响性能，不能直观的从代码中看到数据的流动过程，反射滥用了会导致项目难以维护。
	还有使用unsafe包也一样让其他同志难以阅读。
	Rust的官网有一句话大概意思是：这就是黑魔法也被成为死灵魔法。
	尽量都封装到这里，项目中有用到的地方调用此包，之后有更好的方式不用这儿也容易根除。
*/
package necromancy

import (
	"errors"
	"reflect"
)

// 萃取 在interface中对应struct的字段值
func Extraction(entity any, name string) (val any, err error) {
	chaos := reflect.ValueOf(entity)
	if chaos.Kind() != reflect.Struct {
		return nil, errors.New("reflect:元素非结构体")
	}
	field := chaos.FieldByName(name)
	if field.IsValid() {
		val = field.Interface()
	} else {
		err = errors.New("reflect:未找到指定的字段")
	}

	return
}
