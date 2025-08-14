package golang

import (
	problem "leetCode/problemsKylePade/problemsKylePade_342"
	"testing"
)

func TestSolution(t *testing.T) {
	TestEach(t, "342", "problemsKylePade", problem.Solve)
}
