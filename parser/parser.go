package parser

import (
  "errors"
  "strconv"
)

import "github.com/chrisbutcher/goscheme/types"

func Parse(tokens []string) (types.Expression, []string, error) {
  token := tokens[0]

  if token == "(" {
    list := make([]types.Expression, 0)

    for {
      tokens = tokens[1:]

      if tokens[0] == ")" {
        break
      }

      if parsed, new_tokens, err := Parse(tokens); err == nil {
        tokens = new_tokens
        list = append(list, parsed)
      } else {
        return nil, nil, err
      }
    }
    return list, tokens, nil
  } else if token == ")" {
    return nil, nil, errors.New("Unexpected )")
  } else {
    if token == "#t" {
      return types.Boolean(true), tokens, nil
    } else if token == "#f" {
      return types.Boolean(false), tokens, nil
    } else if number, err := strconv.ParseFloat(token, 64); err == nil {
      return types.Number(number), tokens, nil
    } else {
      return types.Symbol(token), tokens, nil
    }
  }
}
