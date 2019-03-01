package main

type RotorReflectorStack struct {
	rotorStack RotorStack
	reflector Reflector
}

func (rrs *RotorReflectorStack) initStack() {
	rrs.rotorStack.initStack([3]int{0,1,2}, [3]int{0,0,0})
	rrs.reflector = Reflector{ reflectorSequence }
}

func (rrs RotorReflectorStack) code( index int) int {
	first := rrs.rotorStack.encodeLeftIndex(index)
	second := rrs.reflector.outputForInput(first)
	third := rrs.rotorStack.encodeRightIndex(second)

	return third
}