package simple

import "fmt"

type IConfigParser interface {
	Parse(data string)
}

type JsonConfigParser struct{}

func (i *JsonConfigParser) Parse(data string) {
	fmt.Printf("parse data for: %s", data)
}

type YamlConfigParser struct{}

func (i *YamlConfigParser) Parse(data string) {
	fmt.Printf("parse data for: %s", data)
}

func NewIConfigParser(data string) IConfigParser {
	switch data {
	case "json":
		return &JsonConfigParser{}
	case "yaml":
		return &YamlConfigParser{}
	default:
		return nil
	}
}
