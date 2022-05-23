package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"text/template"
)

func TestLogic(t *testing.T) {
	contentStr := getFileContent("F:\\code\\goProject\\LyTools\\templateTest\\demo1\\logic.tpl")
	//params := []string{"id", "name"}
	method := &Method{
		Comment:     "commentAdd",
		Mata:        "fiberDemo",
		MethodType:  "post",
		Path:        "/ccddg",
		PathParams:  nil,
		MiddleWares: "jwt",
		Name:        "Add",
		Request:     "CommReq",
		Response:    "CommRsp",
		Group:       "user",
	}

	tpl, err := template.New("tpl").Funcs(map[string]interface{}{
		"Title": func(s string) string {
			return strings.Title(s)
		},
	}).Parse(contentStr)
	if err != nil {
		return
	}
	// 解析
	content := bytes.NewBuffer([]byte{})
	err = tpl.Execute(content, method)
	if err != nil {
		panic(err)
	}
	fmt.Println(content)
}
