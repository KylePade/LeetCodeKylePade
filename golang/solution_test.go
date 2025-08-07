package golang

import (
	problem "leetCode/problemsKylePade/problemsKylePade_808"
	"testing"
)

func TestSolution(t *testing.T) {
	TestEach(t, "808", "problemsKylePade", problem.Solve)
}
