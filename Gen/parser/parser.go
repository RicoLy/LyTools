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
	messageReg := regexp.MustCompile(constant.RegExpMessage)
	results := messageReg.FindAllStringSubmatch(fileInfo, -1)
	for _, result := range results {
		if len(result) == 4 {
			message := new(models.Message)
			message.Meta = result[0]
			message.Comment = strings.Trim(result[1], " \n")
			message.Name = result[2]
			message.ElementInfos = ParseElementInfo(result[3])
			messages = append(messages, message)
		}
	}
	return messages
}

func ParseElementInfo(elementStr string) []*models.ElementInfo {
	elementInfos := make([]*models.ElementInfo, 0)
	elementReg := regexp.MustCompile(constant.RegExpElementInfo)
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

func ParseService(fileInfo string) []*models.Service {
	services := make([]*models.Service, 0)
	serviceReg := regexp.MustCompile(constant.RegExpService)
	results := serviceReg.FindAllStringSubmatch(fileInfo, -1)
	for _, result := range results {
		service := new(models.Service)
		service.Meta = result[0]

		service.Comment = strings.Trim(FindTopStr(result[1], " @"), "/ ")
		fmt.Println("results len", len(results))
		fmt.Printf("Comment: %+v\n", service.Comment)
	}

	return services
}

func FindTopStr(str string, sep string) string {
	index := strings.Index(str, sep)
	if index == -1 {
		return str
	} else {
		return str[:index]
	}
}
