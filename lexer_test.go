package main

import "testing"
import "reflect"

func TestTokenizeSimpleExpr(t *testing.T) {
	const in = "(+ 1 1)"
	var out = Expr{"(", "+", "1", "1", ")"}
	actual := Tokenize(in)

	if !reflect.DeepEqual(out, actual) {
		t.Error("Expected ", out, " got ", actual)
	}
}

func TestTokenizeNestedExpr(t *testing.T) {
	const in = "(define timestwo (lambda (x) (* x 2)))"
	var out = Expr{"(", "define", "timestwo", "(", "lambda", "(", "x", ")", "(", "*", "x", "2", ")", ")", ")"}

	actual := Tokenize(in)

	if !reflect.DeepEqual(out, actual) {
		t.Error("Expected ", out, " got ", actual)
	}
}
