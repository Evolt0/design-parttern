package factory

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

type IConfigParserFactory interface {
	CreateParser() IConfigParser
}

type JsonConfigParserFactory struct{}

func (i *JsonConfigParserFactory) CreateParser() IConfigParser {
	return &JsonConfigParser{}
}

type YamlConfigParserFactory struct{}

func (i *YamlConfigParserFactory) CreateParser() IConfigParser {
	return &YamlConfigParser{}
}

func NewIConfigParserFactory(t string) IConfigParserFactory {
	switch t {
	case "json":
		return &JsonConfigParserFactory{}
	case "yaml":
		return &YamlConfigParserFactory{}
	default:
		return nil
	}
}
