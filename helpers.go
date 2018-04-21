package regexsolver

import "regexp"

func isSolved(p *RegexPuzzle) bool {
	for _, row := range p.CellRows {
		if !row.IsValidRow() || !row.IsFull() {
			return false
		}
	}
	return true
}

func testEq(a, b []*regexp.Regexp) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}