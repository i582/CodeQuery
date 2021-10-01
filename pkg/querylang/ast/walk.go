package ast

func (n *Root) Walk(v Visitor) {
	if !v.EnterNode(n) {
		return
	}

	for _, stmt := range n.Stmts {
		stmt.Walk(v)
	}

	v.LeaveNode(n)
}

func (n *CommaSeparatedList) Walk(v Visitor) {
	if !v.EnterNode(n) {
		return
	}

	for _, item := range n.Items {
		item.Walk(v)
	}

	v.LeaveNode(n)
}

func (n *Identifier) Walk(v Visitor) {
	if !v.EnterNode(n) {
		return
	}

	v.LeaveNode(n)
}

func (n *Variable) Walk(v Visitor) {
	if !v.EnterNode(n) {
		return
	}

	v.LeaveNode(n)
}

func (n *ComparisonExpr) Walk(v Visitor) {
	if !v.EnterNode(n) {
		return
	}

	n.Left.Walk(v)
	n.Right.Walk(v)

	v.LeaveNode(n)
}

func (n *BinaryExpr) Walk(v Visitor) {
	if !v.EnterNode(n) {
		return
	}

	n.Left.Walk(v)
	n.Right.Walk(v)

	v.LeaveNode(n)
}

func (n *NotExpr) Walk(v Visitor) {
	if !v.EnterNode(n) {
		return
	}

	n.Expr.Walk(v)

	v.LeaveNode(n)
}

func (n *MethodCallExpr) Walk(v Visitor) {
	if !v.EnterNode(n) {
		return
	}

	n.Variable.Walk(v)

	for _, arg := range n.Args {
		arg.Walk(v)
	}

	v.LeaveNode(n)
}

func (n *BasicLit) Walk(v Visitor) {
	if !v.EnterNode(n) {
		return
	}

	v.LeaveNode(n)
}

func (n *SelectExpr) Walk(v Visitor) {
	if !v.EnterNode(n) {
		return
	}

	if n.Select != nil {
		n.Select.Walk(v)
	}

	if n.FromExpr != nil {
		n.FromExpr.Walk(v)
	}

	if n.WhereExpr != nil {
		n.WhereExpr.Walk(v)
	}

	if n.WithExpr != nil {
		n.WithExpr.Walk(v)
	}

	if n.LimitExpr != nil {
		n.LimitExpr.Walk(v)
	}

	v.LeaveNode(n)
}

func (n *SelectSubjectExpr) Walk(v Visitor) {
	if !v.EnterNode(n) {
		return
	}

	if n.List != nil {
		n.List.Walk(v)
	}

	v.LeaveNode(n)
}

func (n *FromExpr) Walk(v Visitor) {
	if !v.EnterNode(n) {
		return
	}

	v.LeaveNode(n)
}

func (n *LimitExpr) Walk(v Visitor) {
	if !v.EnterNode(n) {
		return
	}

	n.Value.Walk(v)

	v.LeaveNode(n)
}

func (n *OrderByExpr) Walk(v Visitor) {
	if !v.EnterNode(n) {
		return
	}

	n.Field.Walk(v)

	v.LeaveNode(n)
}

func (n *WhereExpr) Walk(v Visitor) {
	if !v.EnterNode(n) {
		return
	}

	n.Expr.Walk(v)

	v.LeaveNode(n)
}

func (n *WithExpr) Walk(v Visitor) {
	if !v.EnterNode(n) {
		return
	}

	n.WithList.Walk(v)

	v.LeaveNode(n)
}
