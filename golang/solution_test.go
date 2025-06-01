package golang

import (
	problem "leetCode/problemsKylePade/problemsKylePade_135"
	"testing"
)

func TestSolution(t *testing.T) {
	TestEach(t, "135", "problemsKylePade", problem.Solve)
}
