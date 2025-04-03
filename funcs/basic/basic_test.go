package basic_test

import (
	"testing"

	dfuncs "github.com/arconomy/go-math-expression-parser/funcs/basic"
	"github.com/shopspring/decimal"
)

func TestDefaultOperators(t *testing.T) {
	// UnarySum
	res, err := dfuncs.UnarySum(decimal.NewFromFloat(5.4))
	if err != nil {
		t.Error(err)
	}
	if !res.Equal(decimal.NewFromFloat(5.4)) {
		t.Error("incorrect UnarySum result: " + res.String())
	}

	res, err = dfuncs.UnarySum(decimal.NewFromFloat(5.4), decimal.NewFromFloat(4.3))
	if !res.Equal(decimal.Zero) || err == nil {
		t.Error("incorrect UnarySum error handling")
	}

	res, err = dfuncs.UnarySum()
	if !res.Equal(decimal.Zero) || err == nil {
		t.Error("incorrect UnarySum error handling")
	}

	// UnarySub
	res, err = dfuncs.UnarySub(decimal.NewFromFloat(5.4))
	if err != nil {
		t.Error(err)
	}
	if !res.Equal(decimal.NewFromFloat(-5.4)) {
		t.Error("incorrect UnarySub result: " + res.String())
	}
	res, err = dfuncs.UnarySub(decimal.NewFromFloat(5.4), decimal.NewFromFloat(4.3))
	if !res.Equal(decimal.Zero) || err == nil {
		t.Error("incorrect UnarySub error handling")
	}

	res, err = dfuncs.UnarySub()
	if !res.Equal(decimal.Zero) || err == nil {
		t.Error("incorrect UnarySub error handling")
	}

	// Sqrt
	res, err = dfuncs.Sqrt(decimal.NewFromFloat(9.0))
	if err != nil {
		t.Error(err)
	}
	if !res.Equal(decimal.NewFromFloat(3.0)) {
		t.Error("incorrect Sqrt result: " + res.String())
	}
	res, err = dfuncs.Sqrt(decimal.NewFromFloat(5.4), decimal.NewFromFloat(4.3))
	if !res.Equal(decimal.Zero) || err == nil {
		t.Error("incorrect Sqrt error handling")
	}

	res, err = dfuncs.Sqrt()
	if !res.Equal(decimal.Zero) || err == nil {
		t.Error("incorrect Sqrt error handling")
	}

	res, err = dfuncs.Sqrt(decimal.NewFromFloat(-9.0))
	if !res.Equal(decimal.Zero) || err == nil {
		t.Error("incorrect Sqrt error handling")
	}

	// Abs
	res, err = dfuncs.Abs(decimal.NewFromFloat(-19.0))
	if err != nil {
		t.Error(err)
	}

	if !res.Equal(decimal.NewFromFloat(19.0)) {
		t.Error("incorrect Abs result: " + res.String())
	}

	res, err = dfuncs.Abs(decimal.NewFromFloat(19.0))
	if err != nil {
		t.Error(err)
	}

	if !res.Equal(decimal.NewFromFloat(19.0)) {
		t.Error("incorrect Abs result: " + res.String())
	}

	res, err = dfuncs.Abs(decimal.NewFromFloat(5.4), decimal.NewFromFloat(4.3))
	if !res.Equal(decimal.Zero) || err == nil {
		t.Error("incorrect Abs error handling")
	}

	res, err = dfuncs.Abs()
	if !res.Equal(decimal.Zero) || err == nil {
		t.Error("incorrect Abs error handling")
	}

	// Mult
	res, err = dfuncs.Mult(decimal.NewFromFloat(11.2), decimal.NewFromFloat(3.0))
	if err != nil {
		t.Error(err)
	}

	if !res.Equal(decimal.NewFromFloat(33.6)) {
		t.Error("incorrect Mult result: " + res.String())
	}

	res, err = dfuncs.Mult(decimal.NewFromFloat(11.2), decimal.NewFromFloat(-3.0))
	if err != nil {
		t.Error(err)
	}

	if !res.Equal(decimal.NewFromFloat(-33.6)) {
		t.Error("incorrect Mult result: " + res.String())
	}

	res, err = dfuncs.Mult()
	if !res.Equal(decimal.Zero) || err == nil {
		t.Error("incorrect Mult error handling")
	}

	res, err = dfuncs.Mult(decimal.NewFromFloat(1.0), decimal.NewFromFloat(2.0), decimal.NewFromFloat(3.0))
	if !res.Equal(decimal.Zero) || err == nil {
		t.Error("incorrect Mult error handling")
	}

	res, err = dfuncs.Mult(decimal.NewFromFloat(1.0), decimal.NewFromFloat(2.0), decimal.NewFromFloat(3.0))
	if !res.Equal(decimal.Zero) || err == nil {
		t.Error("incorrect Mult error handling")
	}

	// Div
	res, err = dfuncs.Div(decimal.NewFromFloat(15.0), decimal.NewFromFloat(2.0))
	if err != nil {
		t.Error(err)
	}

	if !res.Equal(decimal.NewFromFloat(7.5)) {
		t.Error("incorrect Div result: " + res.String())
	}

	res, err = dfuncs.Div(decimal.NewFromFloat(-44.0), decimal.NewFromFloat(11.0))
	if err != nil {
		t.Error(err)
	}

	if !res.Equal(decimal.NewFromFloat(-4.0)) {
		t.Error("incorrect Div result: " + res.String())
	}

	_, err = dfuncs.Div(decimal.NewFromFloat(-44.0), decimal.NewFromFloat(0.0))
	if err == nil {
		t.Error("incorrect Mult error handling")
	}

	res, err = dfuncs.Div()
	if !res.Equal(decimal.Zero) || err == nil {
		t.Error("incorrect Div error handling")
	}

	res, err = dfuncs.Div(decimal.NewFromFloat(1.0), decimal.NewFromFloat(2.0), decimal.NewFromFloat(3.0))
	if !res.Equal(decimal.Zero) || err == nil {
		t.Error("incorrect Div error handling")
	}

	res, err = dfuncs.Div(decimal.NewFromFloat(1.0), decimal.NewFromFloat(2.0), decimal.NewFromFloat(3.0))
	if !res.Equal(decimal.Zero) || err == nil {
		t.Error("incorrect Div error handling")
	}

	// Pow
	res, err = dfuncs.Pow(decimal.NewFromFloat(2.0), decimal.NewFromFloat(5.0))
	if err != nil {
		t.Error(err)
	}

	if !res.Equal(decimal.NewFromFloat(32.0)) {
		t.Error("incorrect Pow result: " + res.String())
	}

	res, err = dfuncs.Pow(decimal.NewFromFloat(-3.0), decimal.NewFromFloat(3.0))
	if err != nil {
		t.Error(err)
	}

	if !res.Equal(decimal.NewFromFloat(-27.0)) {
		t.Error("incorrect Pow result: " + res.String())
	}

	res, err = dfuncs.Pow()
	if !res.Equal(decimal.Zero) || err == nil {
		t.Error("incorrect Pow error handling")
	}

	res, err = dfuncs.Pow(decimal.NewFromFloat(1.0), decimal.NewFromFloat(2.0), decimal.NewFromFloat(3.0))
	if !res.Equal(decimal.Zero) || err == nil {
		t.Error("incorrect Pow error handling")
	}

	res, err = dfuncs.Pow(decimal.NewFromFloat(1.0), decimal.NewFromFloat(2.0), decimal.NewFromFloat(3.0))
	if !res.Equal(decimal.Zero) || err == nil {
		t.Error("incorrect Pow error handling")
	}

	// DivReminder
	res, err = dfuncs.DivReminder(decimal.NewFromFloat(17.0), decimal.NewFromFloat(5.0))
	if err != nil {
		t.Error(err)
	}

	if !res.Equal(decimal.NewFromFloat(2.0)) {
		t.Error("incorrect DivReminder result: " + res.String())
	}

	_, err = dfuncs.DivReminder(decimal.NewFromFloat(20.0), decimal.NewFromFloat(0.0))
	if err == nil {
		t.Error(err)
	}

	if err == nil {
		t.Error("incorrect DivReminder error handling")
	}

	res, err = dfuncs.DivReminder()
	if !res.Equal(decimal.Zero) || err == nil {
		t.Error("incorrect DivReminder error handling")
	}

	res, err = dfuncs.DivReminder(decimal.NewFromFloat(1.0), decimal.NewFromFloat(2.0), decimal.NewFromFloat(3.0))
	if !res.Equal(decimal.Zero) || err == nil {
		t.Error("incorrect DivReminder error handling")
	}

	res, err = dfuncs.DivReminder(decimal.NewFromFloat(1.0), decimal.NewFromFloat(2.0), decimal.NewFromFloat(3.0))
	if !res.Equal(decimal.Zero) || err == nil {
		t.Error("incorrect DivReminder error handling")
	}

	// Sum
	res, err = dfuncs.Sum(decimal.NewFromFloat(15.0), decimal.NewFromFloat(0.2))
	if err != nil {
		t.Error(err)
	}

	if !res.Equal(decimal.NewFromFloat(15.2)) {
		t.Error("incorrect Sum result: " + res.String())
	}

	res, err = dfuncs.Sum(decimal.NewFromFloat(-44.0), decimal.NewFromFloat(11.0))
	if err != nil {
		t.Error(err)
	}

	if !res.Equal(decimal.NewFromFloat(-33.0)) {
		t.Error("incorrect Sum result: " + res.String())
	}

	res, err = dfuncs.Sum()
	if !res.Equal(decimal.Zero) || err == nil {
		t.Error("incorrect Sum error handling")
	}

	res, err = dfuncs.Sum(decimal.NewFromFloat(1.0), decimal.NewFromFloat(2.0), decimal.NewFromFloat(3.0))
	if !res.Equal(decimal.Zero) || err == nil {
		t.Error("incorrect Sum error handling")
	}

	res, err = dfuncs.Sum(decimal.NewFromFloat(1.0), decimal.NewFromFloat(2.0), decimal.NewFromFloat(3.0))
	if !res.Equal(decimal.Zero) || err == nil {
		t.Error("incorrect Sum error handling")
	}

	// Sub
	res, err = dfuncs.Sub(decimal.NewFromFloat(15.0), decimal.NewFromFloat(0.2))
	if err != nil {
		t.Error(err)
	}

	if !res.Equal(decimal.NewFromFloat(14.8)) {
		t.Error("incorrect Sub result: " + res.String())
	}

	res, err = dfuncs.Sub(decimal.NewFromFloat(-44.0), decimal.NewFromFloat(11.0))
	if err != nil {
		t.Error(err)
	}

	if !res.Equal(decimal.NewFromFloat(-55.0)) {
		t.Error("incorrect Sub result: " + res.String())
	}

	res, err = dfuncs.Sub()
	if !res.Equal(decimal.Zero) || err == nil {
		t.Error("incorrect Sub error handling")
	}

	res, err = dfuncs.Sub(decimal.NewFromFloat(1.0), decimal.NewFromFloat(2.0), decimal.NewFromFloat(3.0))
	if !res.Equal(decimal.Zero) || err == nil {
		t.Error("incorrect Sub error handling")
	}

	res, err = dfuncs.Sub(decimal.NewFromFloat(1.0), decimal.NewFromFloat(2.0), decimal.NewFromFloat(3.0))
	if !res.Equal(decimal.Zero) || err == nil {
		t.Error("incorrect Sub error handling")
	}
}
