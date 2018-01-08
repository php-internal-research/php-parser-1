package name

import (
	"github.com/z7zmey/php-parser/node"
)

type Name struct {
	position *node.Position
	Parts    []node.Node
}

func NewName(Parts []node.Node) *Name {
	return &Name{
		nil,
		Parts,
	}
}

func (n *Name) Attributes() map[string]interface{} {
	return nil
}

func (n *Name) Position() *node.Position {
	return n.position
}

func (n *Name) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *Name) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Parts != nil {
		vv := v.GetChildrenVisitor("Parts")
		for _, nn := range n.Parts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
