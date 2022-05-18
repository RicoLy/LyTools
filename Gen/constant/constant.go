package constant

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
