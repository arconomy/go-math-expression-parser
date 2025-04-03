package internal

import (
	"errors"

	"github.com/arconomy/go-math-expression-parser/interfaces"
	"github.com/shopspring/decimal"
)

// Unary - the struct which contains a variable and a unary operation
type Unary struct {
	Op  string
	Exp interfaces.Expression
}

func (u *Unary) GetVarList(vars map[string]interface{}) {
	u.Exp.GetVarList(vars)
}

// Evaluate - execute unary operator
func (u *Unary) Evaluate(vars map[string]decimal.Decimal, p interfaces.ExpParser) (decimal.Decimal, error) {
	val, err := u.Exp.Evaluate(vars, p)
	if err != nil {
		return decimal.Zero, err
	}
	indx, exist := UnaryOperatorExist(u.Op, p)
	if !exist {
		return decimal.Zero, errors.New("not supported unary operation: '" + u.Op + "'")
	}
	result, err := p.GetFunctions()[indx][u.Op](val)
	return result, err
}

// toString conversation
func (u *Unary) String() string {
	return "( " + string(u.Op) + " " + u.Exp.String() + " )"
}
