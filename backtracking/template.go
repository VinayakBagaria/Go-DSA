package backtracking

type ElementType = string
type StateType = map[ElementType]struct{}
type SolutionType = []ElementType

func isValidState(state StateType) bool {
	// Check if its a valid solution
	return true
}

func getCandidates(state StateType) []ElementType {
	return []ElementType{}
}

func search(state StateType, solutions SolutionType) {
	if isValidState(state) {
		for key := range state {
			solutions = append(solutions, key)
		}
		// return
	}

	for _, candidate := range getCandidates(state) {
		state[candidate] = struct{}{}
		search(state, solutions)
		delete(state, candidate)
	}
}

func solve() SolutionType {
	state := make(map[ElementType]struct{})
	solutions := SolutionType{}
	search(state, solutions)
	return solutions
}
