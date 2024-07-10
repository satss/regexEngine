package regexengine

type State struct {
	transitions []Transition
}

type Predicate func(input rune) bool

type Transition struct {
	from      *State
	to        *State
	predicate Predicate
}

type Status string

const (
	InProgress = "in_progress"
	Failed     = "failed"
	Success    = "success"
)

type runner struct {
	head    *State
	current *State
}
