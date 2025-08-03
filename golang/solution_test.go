package golang

import (
	problem "leetCode/problemsKylePade/problemsKylePade_904"
	"testing"
)

func TestSolution(t *testing.T) {
	TestEach(t, "904", "problemsKylePade", problem.Solve)
}
