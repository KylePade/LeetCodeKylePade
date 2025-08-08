package golang

import (
	problem "leetCode/problemsKylePade/problemsKylePade_231"
	"testing"
)

func TestSolution(t *testing.T) {
	TestEach(t, "231", "problemsKylePade", problem.Solve)
}
