package regexengine

func NewRunner(head *State) *runner {

	if head == nil {
		return nil
	}

	return &runner{
		head:    head,
		current: head,
	}

}

func (runner *runner) Next(input rune) {

	if runner == nil {
		return
	}

	return &runner{
		head:    head,
		current: head,
	}

}
