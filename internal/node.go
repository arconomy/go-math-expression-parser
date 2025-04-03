package internal

import (
	"errors"

	"github.com/arconomy/go-math-expression-parser/interfaces"
	"github.com/shopspring/decimal"
)

// Node - the struct which contains two variables and a binary operation
type Node struct {
	Op   string
	LExp interfaces.Expression
	RExp interfaces.Expression
}

// Evaluate - execute expression tree
func (n *Node) Evaluate(vars map[string]decimal.Decimal, p interfaces.ExpParser) (decimal.Decimal, error) {
	left, err := n.LExp.Evaluate(vars, p)
	if err != nil {
		return decimal.Zero, err
	}
	right, err := n.RExp.Evaluate(vars, p)
	if err != nil {
		return decimal.Zero, err
	}
	indx, exist := BinaryOperatorExist(n.Op, p)
	if !exist {
		return decimal.Zero, errors.New("not supported binary operation: '" + string(n.Op) + "'")
	}
	result, err := p.GetFunctions()[indx][n.Op](left, right)
	return result, err
}

func (n *Node) GetVarList(vars map[string]interface{}) {
	n.LExp.GetVarList(vars)
	n.RExp.GetVarList(vars)
}

// toString conversation
func (n *Node) String() string {
	return "( " + string(n.Op) + " " + n.LExp.String() + " " + n.RExp.String() + " )"
}
