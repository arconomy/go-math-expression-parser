package userfunc_test

import (
	"errors"
	"strconv"
	"testing"

	"github.com/arconomy/go-math-expression-parser/funcs/userfunc"
	"github.com/arconomy/go-math-expression-parser/interfaces"
	"github.com/arconomy/go-math-expression-parser/internal"
	"github.com/arconomy/go-math-expression-parser/parser"
	"github.com/shopspring/decimal"
)

func foo(args ...decimal.Decimal) (decimal.Decimal, error) {
	if len(args) != 2 {
		return decimal.Zero, errors.New("need 2 args")
	}
	return args[0].Add(args[1]), nil
}

func average(args ...decimal.Decimal) (decimal.Decimal, error) {
	if len(args) < 1 {
		return decimal.Zero, errors.New("need 1 or more args")
	}
	var sum decimal.Decimal = decimal.Zero
	for _, a := range args {
		sum = sum.Add(a)
	}

	return sum.Div(decimal.NewFromInt(int64(len(args)))), nil
}

const float64EqualityThreshold = 1e-9

func almostEqual(a, b decimal.Decimal) bool {
	return a.Sub(b).Abs().LessThanOrEqual(decimal.NewFromFloat(float64EqualityThreshold))
}

func TestGetVarList(t *testing.T) {
	t1, t2, t3 := "a", "b", "c"
	term1 := internal.Term{Val: t1}
	term2 := internal.Term{Val: t2}
	term3 := internal.Term{Val: t3}
	oper1 := internal.Node{Op: "+", LExp: &term1, RExp: &term2}
	f1 := userfunc.Func{"foo", []interfaces.Expression{&oper1, &term3}}

	var vars = map[string]interface{}{}
	f1.GetVarList(vars)

	if len(vars) != 3 {
		t.Error("incorrect map keys count = " + strconv.Itoa(len(vars)))
	}

	if _, ok := vars[t1]; !ok {
		t.Error("not found t1")
	}

	if _, ok := vars[t2]; !ok {
		t.Error("not found t2")
	}

	if _, ok := vars[t3]; !ok {
		t.Error("not found t3")
	}

	// test empty
	f2 := userfunc.Func{"foo", []interfaces.Expression{}}

	vars = map[string]interface{}{}
	f2.GetVarList(vars)

	if len(vars) != 0 {
		t.Error("incorrect map keys count = " + strconv.Itoa(len(vars)))
	}
}

func TestEvaluate(t *testing.T) {
	p := parser.NewParser()
	p.AddFunction(foo, "foo")
	p.AddFunction(average, "average")
	//exp1 := "foo(average(2, 4, 9), 100)"

	term1 := internal.Term{Val: "2"}
	term2 := internal.Term{Val: "4"}
	term3 := internal.Term{Val: "9"}
	term4 := internal.Term{Val: "100"}
	f1 := userfunc.Func{"average", []interfaces.Expression{&term1, &term2, &term3}}
	f2 := userfunc.Func{"foo", []interfaces.Expression{&f1, &term4}}
	f3 := userfunc.Func{"foo", []interfaces.Expression{&term1, &term2, &term3}} // foo with incorrect Args count
	f4 := userfunc.Func{"foo", []interfaces.Expression{&f3, &term2}}

	var vars = map[string]decimal.Decimal{}
	res, err := f1.Evaluate(vars, p)
	if err != nil {
		t.Error(err)
	}
	if !almostEqual(res, decimal.NewFromFloat(5.0)) {
		t.Error("incorrect result = " + res.String())
	}

	res, err = f2.Evaluate(vars, p)
	if err != nil {
		t.Error(err)
	}
	if !almostEqual(res, decimal.NewFromFloat(105.0)) {
		t.Error("incorrect result = " + res.String())
	}

	_, err = f3.Evaluate(vars, p)
	if err == nil {
		t.Error("foo error was not handled!")
	}

	_, err = f4.Evaluate(vars, p)
	if err == nil {
		t.Error("foo error was not handled!")
	}
}

func TestString(t *testing.T) {
	term1 := internal.Term{Val: "2"}
	term2 := internal.Term{Val: "4"}
	term3 := internal.Term{Val: "9"}
	term4 := internal.Term{Val: "100"}
	f1 := userfunc.Func{"average", []interfaces.Expression{&term1, &term2, &term3}}
	f2 := userfunc.Func{"foo", []interfaces.Expression{&f1, &term4}}

	if f1.String() != "( average ( 2,4,9 ) )" {
		t.Error("incorrect string conversion = " + f1.String())
	}
	if f2.String() != "( foo ( ( average ( 2,4,9 ) ),100 ) )" {
		t.Error("incorrect string conversion = " + f2.String())
	}
}

func TestSetOperation(t *testing.T) {
	op1, op2 := "op1", "op2"

	f1 := userfunc.Func{}
	f2 := userfunc.Func{}

	f1.SetOperation(op1)
	f2.SetOperation(op2)

	if f1.Op != op1 {
		t.Error("incorrect Op = " + f1.Op)
	}

	if f2.Op != op2 {
		t.Error("incorrect Op = " + f1.Op)
	}
}

func TestGetOperation(t *testing.T) {
	op1, op2 := "op1", "op2"

	f1 := userfunc.Func{}
	f2 := userfunc.Func{}

	f1.Op = op1
	f2.Op = op2

	if f1.GetOperation() != op1 {
		t.Error("incorrect Op = " + f1.GetOperation())
	}

	if f2.GetOperation() != op2 {
		t.Error("incorrect Op = " + f1.GetOperation())
	}
}

func argsAreEqual(args1, args2 []interfaces.Expression) bool {
	if len(args1) != len(args2) {
		return false
	}

	for i := 0; i < len(args1); i++ {
		if args1[i] != args2[i] {
			return false
		}
	}
	return true
}

func TestSetArgs(t *testing.T) {
	term1 := internal.Term{Val: "2"}
	term2 := internal.Term{Val: "4"}
	term3 := internal.Term{Val: "9"}
	term4 := internal.Term{Val: "100"}

	pack1 := []interfaces.Expression{&term1, &term2}
	pack2 := []interfaces.Expression{&term3, &term4}

	f1 := userfunc.Func{}
	f2 := userfunc.Func{}

	f1.SetArgs(pack1)
	f2.SetArgs(pack2)

	if !argsAreEqual(f1.Args, pack1) {
		t.Error("incorrect Args")
	}

	if !argsAreEqual(f2.Args, pack2) {
		t.Error("incorrect Args")
	}
}

func TestGetArgs(t *testing.T) {
	term1 := internal.Term{Val: "2"}
	term2 := internal.Term{Val: "4"}
	term3 := internal.Term{Val: "9"}
	term4 := internal.Term{Val: "100"}

	pack1 := []interfaces.Expression{&term1, &term2}
	pack2 := []interfaces.Expression{&term3, &term4}

	f1 := userfunc.Func{}
	f2 := userfunc.Func{}

	f1.Args = pack1
	f2.Args = pack2

	if !argsAreEqual(f1.GetArgs(), pack1) {
		t.Error("incorrect Args")
	}

	if !argsAreEqual(f2.GetArgs(), pack2) {
		t.Error("incorrect Args")
	}
}

func TestUserFunction(t *testing.T) {
	func1 := func(args ...decimal.Decimal) (decimal.Decimal, error) {
		return args[0].Mul(args[1]).Sub(args[2]), nil
	}
	func2 := func(args ...decimal.Decimal) (decimal.Decimal, error) {
		return args[0].Add(args[1]).Add(args[2]), nil
	}
	pars := parser.NewParser()
	pars.AddFunction(func1, "func1")
	pars.AddFunction(func2, "func2")

	type TestVars map[string]decimal.Decimal
	type TestData struct {
		input  string
		vars   TestVars
		output decimal.Decimal
	}

	data := []TestData{
		{"func1(2, 3, 1)", TestVars{}, decimal.NewFromInt(5)},
		{"func1(2^x, y, (x+y))", TestVars{"x": decimal.NewFromInt(3), "y": decimal.NewFromFloat(5.1)}, decimal.NewFromFloat(32.7)},
		{"func2(600, 60, 6)", TestVars{}, decimal.NewFromInt(666)},
		{"func2(func2(700,70,7), 222, -8)", TestVars{}, decimal.NewFromInt(991)},
	}

	for _, d := range data {
		_, err := pars.Parse(d.input)
		if err != nil {
			t.Error(err)
		}
		res, err := pars.Evaluate(d.vars)
		if err != nil {
			t.Error(err)
		}
		if !almostEqual(res, d.output) {
			t.Error("incorrect result, need: " + d.output.String() + ", but get: " + res.String())
		}
	}
}
