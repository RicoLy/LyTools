package parser

import (
	"LyTools/Gen/constant"
	"LyTools/Gen/models"
	"fmt"
	"regexp"
	"strings"
)

func ParseMessage(fileInfo string) []*models.Message {
	messages := make([]*models.Message, 0)
	messageReg := regexp.MustCompile("\\/\\/\\s[\u4e00-\u9fa5\\w\\s|]*message\\s*(\\w*)\\s*{([\\s\\|\\*\\/@\\w:\",=;]*)}")
	results := messageReg.FindAllStringSubmatch(fileInfo, -1)
	commentReg := regexp.MustCompile("\\/\\/\\s*([\u4e00-\u9fa5\\w\\s]*)message")
	for _, result := range results {
		if len(result) == 3 {
			message := new(models.Message)
			message.Meta = result[0]
			message.Name = result[1]
			commentRes := commentReg.FindAllStringSubmatch(message.Meta, -1)
			if len(commentRes) == 1 {
				if len(commentRes[0]) == 2 {
					message.Comment = strings.Trim(commentRes[0][1], "\n")
				}
			}
			// fmt.Println("Comment: ", message.Comment)
			message.ElementInfos = ParseElementInfo(result[2])
			// fmt.Println("-----------------**********------------------", len(result))
			// fmt.Println("Comment: ", message.Comment)
			messages = append(messages, message)
		}
	}
	return messages
}

func ParseElementInfo(elementStr string) []*models.ElementInfo {
	elementInfos := make([]*models.ElementInfo, 0)
	elementReg := regexp.MustCompile("\\/\\/\\s*@\\w*\\:\\s*([\\s\\w@:\",|=*]*)\\n([\\s\\w]*)=\\s*\\d;")
	results := elementReg.FindAllStringSubmatch(elementStr, -1)
	//typeNameReg := regexp.MustCompile("")
	for _, result := range results {
		element := new(models.ElementInfo)
		element.Meta = result[0]
		element.Tags = result[1]
		typeNames := strings.Split(strings.Trim(result[2], " "), " ")
		if len(typeNames) == 3 && typeNames[0] == "repeated" {
			element.Type = fmt.Sprintf("[]*%s", typeNames[1])
			element.Name = typeNames[2]
		} else {
			element.Type = typeNames[0]
			element.Name = typeNames[1]
		}
		if golangType := constant.ProtoTypeToGoType[element.Type]; golangType != "" {
			element.Type = constant.ProtoTypeToGoType[element.Type]
		}

		elementInfos = append(elementInfos, element)
	}

	return elementInfos
}