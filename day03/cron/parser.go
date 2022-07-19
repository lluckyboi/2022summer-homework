package cron

type Parser struct {
	options ParseOpt
}

// Parse 解析 返回工作周期和err
func (p Parser) Parse(spec string) (Cycle, error) {}

// NewParser parse options自定义parse
func NewParser(options ParseOpt) Parser {
	return Parser{options}
}

//默认解析器
var standardParser = NewParser(
//..
)
