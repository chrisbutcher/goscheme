package main

import "testing"

func SetupEnvironment() Environment {
	env := Environment{}
	env.initialize()
	initializeBuiltins()
	return env
}

func TestEvalSimpleSexpr(t *testing.T) {
	in := "(* 2 2)"
	ast, _ := Parenthesize(Tokenize(in))
	env := SetupEnvironment()
	actual := Eval(ast, env)
	out := 4.0

	if actual.(Atom).valNum != out {
		t.Error("Expected ", out, " got ", actual)
	}
}

func TestEvalDefineLambda(t *testing.T) {
	in := "(define timestwo (lambda (x) (* x 2)))"
	ast, _ := Parenthesize(Tokenize(in))
	env := SetupEnvironment()
	actual := Eval(ast, env)
	outType := atomLambda

	if actual.(Atom).typ != outType {
		t.Error("Expected ", outType, " got ", actual)
	}
}

func TestEvalDefineAndUseLambda(t *testing.T) {
	define_in := "(define timestwo (lambda (x) (* x 2)))"

	ast, _ := Parenthesize(Tokenize(define_in))
	env := SetupEnvironment()
	actual_out := Eval(ast, env)
	expected_type := atomLambda

	if actual_out.(Atom).typ != expected_type {
		t.Error("Expected ", expected_type, " got ", actual_out.(Atom).typ)
	}

	in_lm := "(timestwo 50)"
	ast, _ = Parenthesize(Tokenize(in_lm))
	actual_out = Eval(ast, env)
	expected_valnum := 100.0

	if actual_out.(Atom).valNum != expected_valnum {
		t.Error("Expected ", expected_valnum, " got ", actual_out.(Atom).valNum)
	}
}

func TestCarCdr(t *testing.T) {
	in := "(car (cdr (cdr (quote 1 2 3))))"
	ast, _ := Parenthesize(Tokenize(in))
	env := SetupEnvironment()
	actual := Eval(ast, env)
	expected := 3.0

	if actual.(Atom).valNum != expected {
		t.Error("Expected ", expected, " got ", actual)
	}
}

func TestAnonymousLambdas(t *testing.T) {
	in := "((lambda (x) (+ x 5)) ((lambda (y) (+ y 1)) 1))"
	ast, _ := Parenthesize(Tokenize(in))
	env := SetupEnvironment()
	actual := Eval(ast, env)
	expected := 7.0

	if actual.(Atom).valNum != expected {
		t.Error("Expected ", expected, " got ", actual)
	}
}

func TestFindVariableInParentEnvironment(t *testing.T) {
	in := "(define y 41)"
	ast, _ := Parenthesize(Tokenize(in))
	env := SetupEnvironment()
	actual := Eval(ast, env)
	expected := 41.0

	if actual.(Atom).valNum != expected {
		t.Error("Expected ", expected, " got ", actual)
	}

	in = "((lambda (x) (+ x y)) 1)"
	ast, _ = Parenthesize(Tokenize(in))
	actual = Eval(ast, env)
	expected = 42.0

	if actual.(Atom).valNum != expected {
		t.Error("Expected ", expected, " got ", actual)
	}
}

func TestCondGreater(t *testing.T) {
	in := "(> 2 1)"
	ast, _ := Parenthesize(Tokenize(in))
	env := SetupEnvironment()
	actual := Eval(ast, env)
	expected := true

	if actual.(Atom).BooleanValue() != expected {
		t.Error("Expected ", expected, " got ", actual)
	}
}

func TestCondLessThan(t *testing.T) {
	in := "(< 2 1)"
	ast, _ := Parenthesize(Tokenize(in))
	env := SetupEnvironment()
	actual := Eval(ast, env)
	expected := false

	if actual.(Atom).BooleanValue() != expected {
		t.Error("Expected ", expected, " got ", actual)
	}
}
