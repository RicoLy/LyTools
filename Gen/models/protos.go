package models

// proto信息
// 结构体信息 message\s*(\w*)\s*{([\s\|\*\/@\w:",=;]*)}
// \/\/\s*([\u4e00-\u9fa5\w\s|])*message\s*(\w*)\s*{([\s\|\*\/@\w:",=;]*)}
type Message struct {
	Comment      string        // 注释
	Meta         string        // 元数据
	Name         string        // 消息名
	ElementInfos []*ElementInfo // proto字段
}

// 字段信息 \/\/\s*@\w*\:\s*([\s\w@:",|=*]*)\n([\s\w]*)=\s*\d;;
// \s*(\w*)\s*(\w*)\s*=\s*(\w*)\s*;
// \w*:"[\w\|\*=]*"
type ElementInfo struct {
	Meta string // 元数据
	Name string // 名称
	Type string // golang 数据类型
	Tags string // 标签信息 tag | value  json:"pid" form:"pid"
}

// 方法信息 rpc[\s]*([\w]*)\((\w*)\)\s*\w*\s*\((\w*)\)\s*{}
type Method struct {
	Comment     string // 注释
	Mata        string // 源信息
	MethodType  string // 方法类型
	MiddleWares string // 中间件
	Name        string // 方法名
	Request     string // 参数
	Response    string // 返回值
}

// 服务信息  service\s*(\w*)\s*{[\s\w(){};]*}
// \/\/\s*([\u4e00-\u9fa5\w\s@:"\/]*)service\s*(\w*)\s*{([\s/\u4e00-\u9fa5\w@:(){}|]*)}\s*}
type Service struct {
	Comment string   // 注释
	Meta    string   // 元数据
	Name    string   // 服务名
	Group   string   // 分组信息
	Prefix  string   // 路由前缀
	Methods []*Method // 方法
}
