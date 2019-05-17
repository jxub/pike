package dom

import (
	"errors"
	"strings"
)

type AttrMap map[string]string
type NodeData string

type Node struct {
	Children []Node
	NodeType Attributer // string || ElementData
}

type Attributer interface {
	Attributes() AttrMap
}

type ElementData struct {
	TagName    string
	Attributes AttrMap
}

func (s NodeData) Attributes() AttrMap {
	return map[string]string{}
}

// pub fn text
func NewNodeText(data interface{}) Node {
	return Node{
		Children: make([]Node, 0),
		NodeType: data.(NodeData),
	}
}

// pub fn elem
func NewNodeElem(name string, attrs AttrMap, children []Node) Node {
	return Node{
		Children: children,
		NodeType: ElementData{
			TagName:    name,
			Attributes: attrs,
		},
	}
}

func (e ElementData) Attributes() AttrMap {
	return e.Attributes
}

func (e *ElementData) ID() (string, error) {
	val, ok := e.Attributes["id"]
	if !ok {
		return "", errors.New("ID not found")
	} else {
		return val, nil
	}
}

func (e *ElementData) Classes() []string {
	classList, ok := e.Attributes["class"]
	if !ok {
		return nil
	}

	return strings.Split(classList, " ")
}
