package main

import "fmt"

type Parser interface {
	Parse() string
}

type xmlParser struct {
}
type textParser struct {
}

func (x xmlParser) Parse() string {
	return "done by xml"
}
func (x textParser) Parse() string {
	return "done by text"
}
func main() {
	xml := xmlParser{}
	text := textParser{}
	parsers := []Parser{xml, text}
	for _, parser := range parsers {
		fmt.Println(parser.Parse())
	}
}
