package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type TraitUsePrecedence struct {
	position  *node.Position
	Ref       node.Node
	Insteadof node.Node
}

func NewTraitUsePrecedence(Ref node.Node, Insteadof node.Node) *TraitUsePrecedence {
	return &TraitUsePrecedence{
		nil,
		Ref,
		Insteadof,
	}
}

func (n *TraitUsePrecedence) Attributes() map[string]interface{} {
	return nil
}

func (n *TraitUsePrecedence) Position() *node.Position {
	return n.position
}

func (n *TraitUsePrecedence) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *TraitUsePrecedence) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Ref != nil {
		vv := v.GetChildrenVisitor("Ref")
		n.Ref.Walk(vv)
	}

	if n.Insteadof != nil {
		vv := v.GetChildrenVisitor("Insteadof")
		n.Insteadof.Walk(vv)
	}

	v.LeaveNode(n)
}
