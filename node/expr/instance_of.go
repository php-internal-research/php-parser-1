package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type InstanceOf struct {
	position *node.Position
	Expr     node.Node
	Class    node.Node
}

func NewInstanceOf(Expr node.Node, Class node.Node) *InstanceOf {
	return &InstanceOf{
		nil,
		Expr,
		Class,
	}
}

func (n *InstanceOf) Attributes() map[string]interface{} {
	return nil
}

func (n *InstanceOf) Position() *node.Position {
	return n.position
}

func (n *InstanceOf) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *InstanceOf) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	if n.Class != nil {
		vv := v.GetChildrenVisitor("Class")
		n.Class.Walk(vv)
	}

	v.LeaveNode(n)
}
