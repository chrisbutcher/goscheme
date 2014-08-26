package main

var builtins map[string]func([]Atom, Environment) Atom

func isBuiltIn(s string) bool {
	_, result := builtins[s]
	return result
}

func initializeBuiltins() {
	builtins = make(map[string]func([]Atom, Environment) Atom)
	builtins["+"] = Add
	builtins["-"] = Subtract
	builtins["*"] = Multiply
	builtins["/"] = Divide
	builtins["define"] = Define
	builtins["quote"] = Quote
	builtins["car"] = Car
	builtins["cdr"] = Cdr
	builtins["if"] = CondIf
	builtins[">"] = CondGreater
	builtins["<"] = CondLessThan
}

func CondIf(input []Atom, env Environment) Atom {
	if input[0].typ == atomBoolean && input[0].val.(bool) == true {
		return input[1]
	} else {
		return input[2]
	}
}

func CondGreater(input []Atom, env Environment) Atom {
	boolean := input[0].valNum > input[1].valNum
	return Atom{typ: atomBoolean, val: boolean}
}

func CondLessThan(input []Atom, env Environment) Atom {
	boolean := input[0].valNum < input[1].valNum
	return Atom{typ: atomBoolean, val: boolean}
}

func Quote(input []Atom, env Environment) Atom {
	return Atom{typ: atomQuote, val: input}
}

func Car(input []Atom, env Environment) Atom {
	return createAtom(input[0].val.([]Atom)[0].val)
}

func Cdr(input []Atom, env Environment) Atom {
	return Quote(input[0].val.([]Atom)[1:], env)
}

func Add(input []Atom, env Environment) Atom {
	var sum float64
	for _, n := range input {
		sum += n.valNum
	}
	return Atom{typ: atomFloat, val: floatToString(sum), valNum: sum}
}

func Subtract(input []Atom, env Environment) Atom {
	diff := input[0].valNum
	for _, n := range input[1:] {
		diff -= n.valNum
	}
	return Atom{typ: atomFloat, val: floatToString(diff), valNum: diff}
}

func Multiply(input []Atom, env Environment) Atom {
	product := input[0].valNum
	for _, n := range input[1:] {
		product *= n.valNum
	}
	return Atom{typ: atomFloat, val: floatToString(product), valNum: product}
}

func Divide(input []Atom, env Environment) Atom {
	quotient := input[0].valNum
	for _, n := range input[1:] {
		quotient /= n.valNum
	}
	return Atom{typ: atomFloat, val: floatToString(quotient), valNum: quotient}
}

func Define(input []Atom, env Environment) Atom {
	env.set(input[0].val.(string), input[1])
	return input[1]
}
