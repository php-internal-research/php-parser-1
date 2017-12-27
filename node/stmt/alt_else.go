package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func(n AltElse) Name() string {
	return "AltElse"
}

type AltElse struct {
	name  string
	token token.Token
	stmt  node.Node
}

func NewAltElse(token token.Token, stmt node.Node) node.Node {
	return AltElse{
		"AltElse",
		token,
		stmt,
	}
}

func (n AltElse) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.name, n.token.StartLine, n.token.EndLine, n.token.Value)

	if n.stmt != nil {
		fmt.Fprintf(out, "\n%vstmt:", indent+"  ")
		n.stmt.Print(out, indent+"    ")
	}
}