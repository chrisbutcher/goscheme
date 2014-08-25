package main

type Environment struct {
	env    map[string]T
	parent *Environment
}

func (e *Environment) set(key string, i T) {
	e.env[key] = i
}

func (e *Environment) get(key string) (T, bool) {
	result, found := e.env[key]

	if !found && e.parent != nil {
		result, found = e.parent.get(key)
	}

	return result, found
}

func (e *Environment) initialize() {
	e.env = make(map[string]T)
}

func New(keys, values []Atom, parentEnvironment Environment) Environment {
	e := Environment{}
	e.env = make(map[string]T)
	e.parent = &parentEnvironment

	for i, key := range keys {
		e.set(key.val.(string), values[i])
	}

	return e
}
