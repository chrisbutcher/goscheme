package types

import (
  "fmt"
  "strings"
)

type Expression interface{}

type Symbol string
type Number float64
type Boolean bool

type Lambda struct {
  Args Expression
  Body Expression
  Env  *Environment
}

type EnvVars map[Symbol]Expression
type Environment struct {
  EnvVars
  Parent *Environment
}

func (b Boolean) String() string {
  switch b {
  case true:
    return "#t"
  default:
    return "#f"
  }
}

func String(exp Expression) string {
  switch exp := exp.(type) {
  case []Expression:
    output := make([]string, len(exp))
    for i, atom := range exp {
      output[i] = String(atom) + " "
    }
    return strings.Join(output, " ")
  default:
    return fmt.Sprint(exp)
  }
}
