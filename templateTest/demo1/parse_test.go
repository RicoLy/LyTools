package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"text/template"
)

var method = &Method{
	Comment:     "commentAdd",
	Mata:        "fiberDemo",
	MethodType:  "post",
	Path:        "/ccddg",
	PathParams:  nil,
	MiddleWares: []string{"jwt, casBin"},
	Name:        "Add",
	Request:     "CommReq",
	Response:    "CommRsp",
	Group:       "user",
}

var methods = []*Method{
	&Method{
		Comment:     "commentAdd",
		Mata:        "fiberDemo",
		MethodType:  "get",
		Path:        "/ccddg",
		PathParams:  nil,
		MiddleWares: nil,
		Name:        "Add",
		Request:     "CommReq",
		Response:    "CommRsp",
		Group:       "user",
	},
	&Method{
		Comment:     "commentAdd",
		Mata:        "fiberDemo",
		MethodType:  "post",
		Path:        "/ccddg",
		PathParams:  nil,
		MiddleWares: []string{"jwt", "casBin"},
		Name:        "Add",
		Request:     "CommReq",
		Response:    "CommRsp",
		Group:       "user",
	},
}

var methods2 = []*Method{
	&Method{
		Comment:     "commentAdd",
		Mata:        "fiberDemo",
		MethodType:  "post",
		Path:        "/ccddg",
		PathParams:  nil,
		MiddleWares: []string{"jwt", "casBin"},
		Name:        "Add",
		Request:     "AddReq",
		Response:    "CommRsp",
		Group:       "order",
	},
	&Method{
		Comment:     "commentAdd",
		Mata:        "fiberDemo",
		MethodType:  "get",
		Path:        "/ccddg",
		PathParams:  nil,
		MiddleWares: nil,
		Name:        "Add",
		Request:     "CommReq",
		Response:    "CommRsp",
		Group:       "order",
	},
}

var service1 = &Service{
	Comment: "userServiceComment",
	Meta:    "fiberDemo",
	Name:    "userService",
	Group:   "user",
	Prefix:  "/user",
	Methods: methods,
}

var service2 = &Service{
	Comment: "orderServiceComment",
	Meta:    "fiberDemo",
	Name:    "orderService",
	Group:   "order",
	Prefix:  "/order",
	Methods: methods2,
}

var elements = []*ElementInfo{
	&ElementInfo{
		Meta: "test",
		Name: "Id",
		Type: "string",
		Tags: "json:\"id\" validate:\"required\"",
	},
	&ElementInfo{
		Meta: "test",
		Name: "Name",
		Type: "string",
		Tags: "json:\"name\" validate:\"required\"",
	},
	&ElementInfo{
		Meta: "test",
		Name: "Role",
		Type: "int",
		Tags: "json:\"role\" validate:\"required\"",
	},
}

var messages = []*Message{
	&Message{
		Comment:      "AddReq添加请求",
		Meta:         "添加请求",
		Name:         "AddReq",
		ElementInfos: elements,
	},
	&Message{
		Comment:      "AddReq添加请求",
		Meta:         "添加请求",
		Name:         "AddReq",
		ElementInfos: elements,
	},
	&Message{
		Comment:      "AddReq添加请求",
		Meta:         "添加请求",
		Name:         "AddReq",
		ElementInfos: elements,
	},
}


func TestLogic(t *testing.T) {
	contentStr := getFileContent("F:\\code\\goProject\\LyTools\\templateTest\\demo1\\logic.tpl")
	//params := []string{"id", "name"}

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

func TestRoutes(t *testing.T) {
	services := []*Service{service1, service2}
	project := &Project{
		Services: services,
		Name:     "fiberDemo",
	}

	contentStr := getFileContent("F:\\code\\goProject\\LyTools\\templateTest\\demo1\\routes.tpl")
	//params := []string{"id", "name"}

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
	err = tpl.Execute(content, project)
	if err != nil {
		panic(err)
	}
	fmt.Println(content)
}


func TestTypes(t *testing.T) {

	contentStr := getFileContent("F:\\code\\goProject\\LyTools\\templateTest\\demo1\\types.tpl")
	//params := []string{"id", "name"}

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
	err = tpl.Execute(content, messages)
	if err != nil {
		panic(err)
	}
	fmt.Println(content)
}