package node

type Nullable struct {
	position *Position
	Expr     Node
}

func NewNullable(Expression Node) *Nullable {
	return &Nullable{
		nil,
		Expression,
	}
}

func (n *Nullable) Attributes() map[string]interface{} {
	return nil
}

func (n *Nullable) Position() *Position {
	return n.position
}

func (n *Nullable) SetPosition(p *Position) Node {
	n.position = p
	return n
}

func (n *Nullable) Walk(v Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
