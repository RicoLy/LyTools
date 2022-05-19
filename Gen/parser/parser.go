package parser

import (
	"LyTools/Gen/constant"
	"LyTools/Gen/models"
	"fmt"
	"regexp"
	"strings"
)

func ParseMessages(fileInfo string) []*models.Message {
	messages := make([]*models.Message, 0)
	messageReg := regexp.MustCompile(constant.RegExpMessage)
	results := messageReg.FindAllStringSubmatch(fileInfo, -1)
	for _, result := range results {
		if len(result) == 4 {
			message := new(models.Message)
			message.Meta = result[0]
			message.Comment = strings.Trim(result[1], " \n")
			message.Name = result[2]
			message.ElementInfos = ParseElementInfos(result[3])
			messages = append(messages, message)
		}
	}
	return messages
}

func ParseElementInfos(elementStr string) []*models.ElementInfo {
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

func ParseServices(fileInfo string) []*models.Service {
	services := make([]*models.Service, 0)
	serviceReg := regexp.MustCompile(constant.RegExpService)
	groupAndPrefixReg := regexp.MustCompile(constant.RegExpAnoInfo)
	results := serviceReg.FindAllStringSubmatch(fileInfo, -1)
	for _, result := range results {
		if len(result) == 4 {
			service := new(models.Service)
			service.Meta = result[0]
			service.Comment = strings.Trim(FindTopStr(result[1], " @"), "/ ")
			service.Name = result[2]
			gpResults := groupAndPrefixReg.FindAllStringSubmatch(result[1], -1)
			for _, gpResult := range gpResults {
				if len(gpResult) == 3 {
					if gpResult[1] == "group" {
						service.Group = gpResult[2]
					}
					if gpResult[1] == "prefix" {
						service.Prefix = gpResult[2]
					}
				}
			}
			service.Methods = ParseMethods(result[3])
			services = append(services, service)
		}
	}

	return services
}

func ParseMethods(methodStr string) []*models.Method {
	methods := make([]*models.Method, 0)
	methodReg := regexp.MustCompile(constant.RegExpMethod)
	extraInfoReg := regexp.MustCompile(constant.RegExpAnoInfo)
	results := methodReg.FindAllStringSubmatch(methodStr, -1)
	for _, result := range results {
		if len(result) == 5 {
			method := new(models.Method)
			method.Mata = result[0]
			method.Comment = strings.Trim(FindTopStr(result[1], " @"), "/ ")
			extraResults := extraInfoReg.FindAllStringSubmatch(result[1], -1)
			for _, extraResult := range extraResults {
				if extraResult[1] == "method" {
					method.MethodType = extraResult[2]
				}
				if extraResult[1] == "middleware" {
					method.MiddleWares = extraResult[2]
				}
			}
			method.Name = result[2]
			method.Request = result[3]
			method.Response = result[4]
			methods = append(methods, method)
		}
	}
	//fmt.Println("================================")
	//fmt.Println("methodStr:", methodStr)
	return methods
}

func FindTopStr(str string, sep string) string {
	index := strings.Index(str, sep)
	if index == -1 {
		return str
	} else {
		return str[:index]
	}
}
