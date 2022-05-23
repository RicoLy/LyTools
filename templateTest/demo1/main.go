package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

type Method struct {
	Comment     string   // 注释
	Mata        string   // 源信息
	MethodType  string   // 方法类型
	Path        string   // 路径
	PathParams  []string // 路径参数
	MiddleWares string   // 中间件
	Name        string   // 方法名
	Request     string   // 参数
	Response    string   // 返回值
	Group       string   // 分组名
}

func main() {
	contentStr := getFileContent("F:\\code\\goProject\\LyTools\\templateTest\\demo1\\handler.tpl")
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
		Response:    "AddRsp",
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

func getFileContent(fileName string) (content string) {
	f, err := os.OpenFile(fileName, os.O_RDONLY, 0600)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()

	contentByte, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	content = string(contentByte)
	return
}
