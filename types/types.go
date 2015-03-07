package types

import (
  "fmt"
  "strings"
)

type Expression interface{}

type Symbol string
type Number float64

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
