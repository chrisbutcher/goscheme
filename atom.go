package main

import (
	"fmt"
	"strconv"
)

type atomEnumType int

const (
	atomSymbol atomEnumType = iota
	atomBuiltin
	atomFloat
	atomQuote
	atomLambda
)

type Atom struct {
	typ           atomEnumType
	val           T
	valNum        float64
	valLambdaArgs []Atom
	valLambdaFn   []Atom
}

func atomTypeToString(t atomEnumType) string {
	mapping := map[atomEnumType]string{atomSymbol: "s", atomBuiltin: "fn", atomFloat: "fl", atomQuote: "qt", atomLambda: "lm"}

	return mapping[t]
}

func (a Atom) String() string {
	if a.typ == atomLambda {
		return fmt.Sprintf("<%s:args(%s):fn(%s)>", atomTypeToString(a.typ), a.valLambdaArgs, a.valLambdaFn)
	} else {
		return fmt.Sprintf("<%s:%s>", atomTypeToString(a.typ), a.val)
	}
}

func floatToString(f float64) string {
	return fmt.Sprintf("%f", f)
}

func createAtom(val T) Atom {
	switch val.(type) {
	case Atom:
		return val.(Atom)
	default:
		value := val.(string)
		new_atom := Atom{}

		float_value, err := strconv.ParseFloat(value, 64)

		new_atom.val = value
		if err != nil {
			if isBuiltIn(value) {
				new_atom.typ = atomBuiltin
			} else if value == "lambda" {
				new_atom.typ = atomLambda
			} else {
				new_atom.typ = atomSymbol
			}
		} else {
			new_atom.typ = atomFloat
			new_atom.valNum = float_value
		}

		return new_atom
	}
}

func genericToAtomSlice(input T) []Atom {
	result := make([]Atom, 0)

	switch input.(type) {
	case Sexpr:
		slice := input.(Sexpr)

		for _, item := range slice {
			result = append(result, item.(Atom))
		}
	case Atom:
		result = append(result, input.(Atom))
	}

	return result
}

func populateLambda(input Sexpr) Atom {
	return Atom{typ: atomLambda, val: "lambda", valLambdaArgs: genericToAtomSlice(input[1]), valLambdaFn: genericToAtomSlice(input[2])}
}

func exprIsLambda(input T) bool {
	switch input.(type) {
	case Atom:
		return input.(Atom).typ == atomLambda
	default:
		return false
	}
}
