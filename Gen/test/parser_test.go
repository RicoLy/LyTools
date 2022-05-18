package test

import (
	"LyTools/Gen/parser"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestParseMessage(t *testing.T) {
	contentStr := getFileContent("../fiber_demo.proto")
	//fmt.Println(contentStr)
	messages := parser.ParseMessage(contentStr)
	for _, message := range messages {
		fmt.Println("================================")
		fmt.Printf("Comment: %+v\n", message.Comment)
		fmt.Printf("Name: %+v\n", message.Name)
		for _, info := range message.ElementInfos {
			fmt.Printf("ElementInfos: %+v\n", info.Name)
		}
		fmt.Println("================================")
	}

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
