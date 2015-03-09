package lexer

import "strings"

func Lex(program string) ([]string, error) {
  p := strings.Replace(program, "(", " ( ", -1)
  p = strings.Replace(p, ")", " ) ", -1)

  tokens := strings.Split(p, " ")

  fn := func(s string) bool {
    return s != " " && s != "" && s != "\n"
  }

  return filterStrings(tokens, fn), nil
}

func filterStrings(s []string, fn func(string) bool) []string {
  var p []string // == nil
  for _, v := range s {
    if fn(v) {
      p = append(p, v)
    }
  }
  return p
}
