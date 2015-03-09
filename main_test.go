package main

import "testing"

import (
  "github.com/chrisbutcher/goscheme/env"
  "github.com/chrisbutcher/goscheme/types"
)

func setupEnvironment() types.Environment {
  globals := types.Environment{}
  env.Initialize(&globals)
  return globals
}

func TestMathEval(t *testing.T) {
  env := setupEnvironment()
  input := "(+ 1.8 (* 5 (/ (- 110 10) (+ 4 1))))"
  actual := ReadEval(input, env)
  expected := 101.8

  if types.Number(expected) != actual {
    t.Error("Expected ", expected, " got ", actual)
  }
}

func TestIfEval(t *testing.T) {
  env := setupEnvironment()
  input := "(if (> 5 1) 12 42)"
  actual := ReadEval(input, env)
  expected := 12

  if types.Number(expected) != actual {
    t.Error("Expected ", expected, " got ", actual)
  }

  input = "(if (< 5 1) 12 42)"
  actual = ReadEval(input, env)
  expected = 42

  if types.Number(expected) != actual {
    t.Error("Expected ", expected, " got ", actual)
  }

  input = "(if (= 5 5) 12 42)"
  actual = ReadEval(input, env)
  expected = 12

  if types.Number(expected) != actual {
    t.Error("Expected ", expected, " got ", actual)
  }
}

func TestConsCar(t *testing.T) {
  env := setupEnvironment()
  input := "(car (cons 1 2))"
  actual := ReadEval(input, env)
  expected := 1

  if types.Number(expected) != actual {
    t.Error("Expected ", expected, " got ", actual)
  }
}

func TestConsCdr(t *testing.T) {
  env := setupEnvironment()
  input := "(cdr (cons 1 2))"
  evaluated := ReadEval(input, env)
  actual := evaluated.([]types.Expression)
  expected := 2

  if types.Number(expected) != actual[0] {
    t.Error("Expected ", expected, " got ", actual)
  }
}

func TestListCar(t *testing.T) {
  env := setupEnvironment()
  input := "(car (list 1 2))"
  actual := ReadEval(input, env)
  expected := 1

  if types.Number(expected) != actual {
    t.Error("Expected ", expected, " got ", actual)
  }
}

func TestListCdr(t *testing.T) {
  env := setupEnvironment()
  input := "(cdr (list 1 2))"
  evaluated := ReadEval(input, env)
  actual := evaluated.([]types.Expression)
  expected := 2

  if types.Number(expected) != actual[0] {
    t.Error("Expected ", expected, " got ", actual)
  }
}

func TestQuoteCar(t *testing.T) {
  env := setupEnvironment()
  input := "(car (quote (1 2)))"
  actual := ReadEval(input, env)
  expected := 1

  if types.Number(expected) != actual {
    t.Error("Expected ", expected, " got ", actual)
  }
}

func TestQuoteCdr(t *testing.T) {
  env := setupEnvironment()
  input := "(cdr (quote (1 2)))"
  evaluated := ReadEval(input, env)
  actual := evaluated.([]types.Expression)
  expected := 2

  if types.Number(expected) != actual[0] {
    t.Error("Expected ", expected, " got ", actual)
  }
}

func TestLambda(t *testing.T) {
  env := setupEnvironment()
  input := "((lambda (x y) (+ x y)) 1 2)"
  actual := ReadEval(input, env)
  expected := 3

  if types.Number(expected) != actual {
    t.Error("Expected ", expected, " got ", actual)
  }
}

func TestBegin(t *testing.T) {
  env := setupEnvironment()
  input := "(begin (define lessthanten (lambda (x) (if (< x 10) #t #f))) (lessthanten 9))"
  actual := ReadEval(input, env)
  expected := true

  if types.Boolean(expected) != actual {
    t.Error("Expected ", expected, " got ", actual)
  }
}

func TestDefineSet(t *testing.T) {
  env := setupEnvironment()
  input := "(define x 5)"
  ReadEval(input, env)

  input = "(set! x 6)"
  _ = ReadEval(input, env)

  input = "(+ x 1)"
  actual := ReadEval(input, env)
  expected := 7

  if types.Number(expected) != actual {
    t.Error("Expected ", expected, " got ", actual)
  }
}

func TestSet(t *testing.T) {
  env := setupEnvironment()
  input := "(set! x 5)"
  actual := ReadEval(input, env)

  if actual != nil {
    t.Error("Expected result to be nil")
  }
}
