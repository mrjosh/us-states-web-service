package main

type States map[string] []string

func (s States) Get(state string) []string {
	return s[state]
}

func (s States) Has(state string) bool {
	return s.Get(state) != nil
}
