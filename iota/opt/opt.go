package main

// 比较运算符枚举定义
type OptType int

const (
	LT OptType = iota
	lE
	GT
	GE
	EQ
	NE
)

func (c OptType) String() string {
	switch c {
	case LT:
		return "<"
	case lE:
		return "<="
	case GT:
		return ">"
	case GE:
		return ">="
	case EQ:
		return "=="
	case NE:
		return "!="
	}
	return ""
}