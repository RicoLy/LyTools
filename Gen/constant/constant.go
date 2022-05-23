package constant

const (
	RegExpMessage     = "\\/\\/\\s*([\u4e00-\u9fa5\\w\\s|]*)message\\s*(\\w*)\\s*{([\\s\\|\\*\\/@\\w:\",=;]*)}"
	RegExpElementInfo = "\\/\\/\\s*@\\w*\\:\\s*([\\s\\w@:\",|=*]*)\\n([\\s\\w]*)=\\s*\\d;"
	RegExpService     = "\\/\\/\\s*([\u4e00-\u9fa5\\w\\s@:\"\\/]*)service\\s*(\\w*)\\s*{([\\s/\u4e00-\u9fa5\\w@:(){}|]*)}\\s*}"
	RegExpAnoInfo     = "@(\\w*):\\s*([/:|\\w]*)"
	RegExpMethod      = "\\/\\/\\s*([\u4e00-\u9fa5\\w\\s@:|/]*)rpc\\s*(\\w*)\\s*\\((\\w*)\\)\\s*returns\\s*\\((\\w*)\\)\\s*{"
)

const (
	TplPostWithReq     = 0
	TplGetWithQueryReq = 1
	TplGetWithPathReq  = 2
	TplGetNoQueryReq   = 3
	TplGetNoPathReq    = 2
)

const (
	NullReqType = "CommReq"
	NullRspType = "CommRsp"
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
