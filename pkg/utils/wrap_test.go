package utils

import (
	"fmt"
	"testing"
)

func Test_wrapSensitiveStr(t *testing.T)  {
	fmt.Println(WrapSensitiveStr(""))
	fmt.Println(WrapSensitiveStr("asdacasc321@qq.com"))
	fmt.Println(WrapSensitiveStr("1892345"))
	fmt.Println(WrapSensitiveStr("18923453211"))
	fmt.Println(WrapSensitiveStr("18923453224323232"))
	fmt.Println(WrapSensitiveStr("鸡太美"))
	fmt.Println(WrapSensitiveStr("战gerger"))
	fmt.Println(WrapSensitiveStr("凡凡大战峰峰"))
}
