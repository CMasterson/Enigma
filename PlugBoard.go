package main

type PlugBoard struct {
	pairings map[string]string
}

func (p *PlugBoard) pair(first string, second string) {
	p.clearValue(first)
	p.clearValue(second)

	p.pairings[first] = second
	p.pairings[second] = first
}

func (p *PlugBoard) clearValue(value string) {
	var oldValue = p.pairings[value]

	p.pairings[value] = ""

	if oldValue == "" {
		p.pairings[oldValue] = ""
	}
}

func (p PlugBoard) valueForInput(input string) string {
	output := p.pairings[input]

	if output == "" {
		return input
	}

	return output
}