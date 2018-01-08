package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Yield struct {
	position *node.Position
	Key      node.Node
	Value    node.Node
}

func NewYield(Key node.Node, Value node.Node) *Yield {
	return &Yield{
		nil,
		Key,
		Value,
	}
}

func (n *Yield) Attributes() map[string]interface{} {
	return nil
}

func (n *Yield) Position() *node.Position {
	return n.position
}

func (n *Yield) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *Yield) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Key != nil {
		vv := v.GetChildrenVisitor("Key")
		n.Key.Walk(vv)
	}

	if n.Value != nil {
		vv := v.GetChildrenVisitor("Value")
		n.Value.Walk(vv)
	}

	v.LeaveNode(n)
}
