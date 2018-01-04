package cast

import (
	"github.com/z7zmey/php-parser/node"
)

type CastInt struct {
	Cast
}

func NewCastInt(expr node.Node) node.Node {
	return &CastInt{
		Cast{
			map[string]interface{}{},
			nil,
			expr,
		},
	}
}

func (n CastInt) Attributes() map[string]interface{} {
	return n.attributes
}

func (n CastInt) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n CastInt) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n CastInt) Position() *node.Position {
	return n.position
}

func (n CastInt) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n CastInt) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
