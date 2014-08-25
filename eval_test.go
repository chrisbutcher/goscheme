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
	parseTree, _ := Parenthesize(Tokenize(in))
	env := SetupEnvironment()
	actual := Eval(parseTree, env)
	out := Sexpr{}

	if actual.(Atom).valNum != 4 {
		t.Error("Expected ", out, " got ", actual)
	}
}
