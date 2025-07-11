package golang

import (
	problem "leetCode/problemsKylePade/problemsKylePade_1900"
	"testing"
)

func TestSolution(t *testing.T) {
	TestEach(t, "1900", "problemsKylePade", problem.Solve)
}
