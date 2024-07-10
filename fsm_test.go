package regexengine

import "testing"

type test struct {
	name           string
	input          string
	expectedStatus Status
}

func TestFSMTT(t *testing.T) {
	tests := []test{
		{"empty string", "", InProgress},
		{"non matching string", "x", Failed},
		{"matching string", "abc", Success},
		{"partial matching string", "ab", InProgress},
	}

	var startState State

	var stateA State

	var stateB State

	var stateC State
	startState.transitions = append(startState.transitions, Transition{to: &stateA, predicate: func(input rune) bool { return input == 'a' }})
	stateA.transitions = append(stateA.transitions, Transition{to: &stateB, predicate: func(input rune) bool { return input == 'b' }})

	stateB.transitions = append(stateA.transitions, Transition{to: &stateC, predicate: func(input rune) bool { return input == 'c' }})

	for _, test := range tests {

	}

}
