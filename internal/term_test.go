package internal_test

import (
	"strconv"
	"testing"

	"github.com/arconomy/go-math-expression-parser/funcs/userfunc"
	"github.com/arconomy/go-math-expression-parser/interfaces"
	"github.com/arconomy/go-math-expression-parser/internal"
	"github.com/arconomy/go-math-expression-parser/parser"
	"github.com/shopspring/decimal"
)

func TestTermGetVarList(t *testing.T) {
	t1, t2, t3 := "", "1.55", "c"
	term1 := internal.Term{Val: t1}
	term2 := internal.Term{Val: t2}
	term3 := internal.Term{Val: t3}

	var vars = map[string]interface{}{}
	term1.GetVarList(vars)

	if len(vars) != 0 {
		t.Error("incorrect map keys count = " + strconv.Itoa(len(vars)))
	}

	vars = map[string]interface{}{}
	term2.GetVarList(vars)

	if len(vars) != 0 {
		t.Error("incorrect map keys count = " + strconv.Itoa(len(vars)))
	}

	vars = map[string]interface{}{}
	term3.GetVarList(vars)

	if len(vars) != 1 {
		t.Error("incorrect map keys count = " + strconv.Itoa(len(vars)))
	}
}

func TestTermEvaluate(t *testing.T) {
	p := parser.NewParser()

	term1 := internal.Term{Val: ""}
	term2 := internal.Term{Val: "4"}
	term3 := internal.Term{Val: "a"}
	term4 := internal.Term{Val: "var3000"}
	// term5 := internal.Term{Val: "R"}

	var vars = map[string]decimal.Decimal{"a": decimal.NewFromFloat(17.7)}
	res, err := term1.Evaluate(vars, p)
	if !res.Equal(decimal.Zero) || err != nil {
		t.Error("incorrect result = " + res.String())
	}

	res, err = term2.Evaluate(vars, p)
	if err != nil {
		t.Error(err)
	}
	if !fuzzyEqual(res, decimal.NewFromFloat(4.0)) {
		t.Error("incorrect result = " + res.String())
	}

	res, err = term3.Evaluate(vars, p)
	if err != nil {
		t.Error(err)
	}
	if !fuzzyEqual(res, decimal.NewFromFloat(17.7)) {
		t.Error("incorrect result = " + res.String())
	}

	res, err = term4.Evaluate(vars, p)
	if !res.Equal(decimal.Zero) || err == nil {
		t.Error("incorrect error handling!")
	}
}

func TestTermString(t *testing.T) {
	term1 := internal.Term{Val: "2"}
	term2 := internal.Term{Val: "4"}
	term3 := internal.Term{Val: "9"}
	term4 := internal.Term{Val: "100"}
	f1 := userfunc.Func{Op: "average", Args: []interfaces.Expression{&term1, &term2, &term3}}
	f2 := userfunc.Func{Op: "foo", Args: []interfaces.Expression{&f1, &term4}}

	if f1.String() != "( average ( 2,4,9 ) )" {
		t.Error("incorrect string conversion = " + f1.String())
	}
	if f2.String() != "( foo ( ( average ( 2,4,9 ) ),100 ) )" {
		t.Error("incorrect string conversion = " + f2.String())
	}
}
