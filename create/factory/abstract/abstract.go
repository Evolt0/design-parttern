package abstract

import "fmt"

type IConfigParser interface {
	Parse(data string)
}

type JsonConfigParser struct{}

func (i *JsonConfigParser) Parse(data string) {
	fmt.Printf("parse data for: %s", data)
}

type IConfigParserFactory interface {
	CreateParser() IConfigParser
}

type JsonConfigParserFactory struct{}

func (i *JsonConfigParserFactory) CreateParser() IConfigParser {
	return &JsonConfigParser{}
}
