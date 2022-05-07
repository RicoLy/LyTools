package main

import (
	"fmt"
	"github.com/mozillazg/go-pinyin"
	"strings"
)

func main() {
	a := pinyin.NewArgs()
	a.Fallback = func(r rune, a pinyin.Args) []string {
		return []string{string(r)}
	}
	pinyinName := pinyin.LazyPinyin("b罗依力", a)
	fmt.Println(strings.Split(pinyinName[0], "")[0])

	phone := "15625205465"

	fmt.Println(strings.Split(pinyinName[0], "")[0] + phone[len(phone)-6:])
}
