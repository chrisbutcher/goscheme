package main

import "testing"
import "reflect"

func TestTokenizeSimpleSexpr(t *testing.T) {
	const in = "(+ 1 1)"
	var out = Sexpr{"(", "+", "1", "1", ")"}
	actual := Tokenize(in)

	if !reflect.DeepEqual(out, actual) {
		t.Error("Expected ", out, " got ", actual)
	}
}

func TestTokenizeNestedSexpr(t *testing.T) {
	const in = "(define timestwo (lambda (x) (* x 2)))"
	var out = Sexpr{"(", "define", "timestwo", "(", "lambda", "(", "x", ")", "(", "*", "x", "2", ")", ")", ")"}

	actual := Tokenize(in)

	if !reflect.DeepEqual(out, actual) {
		t.Error("Expected ", out, " got ", actual)
	}
}
