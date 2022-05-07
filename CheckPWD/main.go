package main

import (
	"regexp"
	"strings"
)

func main() {
	println(CheckPasswordLever(".................."))
	a := "123Ahabd"
	println(strings.ToLower(a))
	var TitleDic = map[string]bool{
		"其他":    true,
		"主任医师":  true,
		"副主任医师": true,
		"主治医师":  true,
		"医师":    true,
	}

	println("TitleDic", TitleDic["其他"])

}


func CheckPasswordLever(ps string) bool {
	if len(ps) < 6 {
		return false
	}
	num := `[0-9]{1}`
	word := `[a-zA-Z]{1}`
	symbol := `[!@#~$%^&*(),.+|_]{1}`
	if b, err := regexp.MatchString(num, ps); !b || err != nil {
		return false
	}
	if b, err := regexp.MatchString(word, ps); b || err != nil {
		return true
	}
	if b, err := regexp.MatchString(symbol, ps); b || err != nil {
		return true
	}
	return false
}