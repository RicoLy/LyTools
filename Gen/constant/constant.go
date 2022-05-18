package constant

const (
	RegExpMessage = "\\/\\/\\s*([\u4e00-\u9fa5\\w\\s|]*)message\\s*(\\w*)\\s*{([\\s\\|\\*\\/@\\w:\",=;]*)}"
	RegExpElementInfo = "\\/\\/\\s*@\\w*\\:\\s*([\\s\\w@:\",|=*]*)\\n([\\s\\w]*)=\\s*\\d;"
	//RegExpService = "\\/\\/\\s*([\u4e00-\u9fa5\\w\\s@:\"\\/]*)service\\s*(\\w*)\\s*{([\\s/\u4e00-\u9fa5\\w@:(){}|]*)}\n}"
	RegExpService = "\\/\\/\\s*([\u4e00-\u9fa5\\w\\s@:\"\\/]*)service\\s*(\\w*)\\s*{([\\s/\u4e00-\u9fa5\\w@:(){}|]*)}\\s*}"
)

var (
	//protobuf类型 <=> golang类型
	ProtoTypeToGoType = map[string]string{
		"double":   "float64",
		"float":    "float32",
		"int32":    "int32",
		"int64":    "int64",
		"uint32":   "uint32",
		"uint64":   "uint64",
		"sint32":   "int32",
		"sint64":   "int64",
		"fixed32":  "uint32",
		"fixed64":  "uint64",
		"sfixed32": "int32",
		"sfixed64": "int64",
		"bool":     "bool",
		"string":   "string",
		"bytes":    "[]byte",
	}
)
