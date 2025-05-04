package golang

import (
	problem "leetCode/problemsKylePade/problemsKylePade_790"
	"testing"
)

func TestSolution(t *testing.T) {
	TestEach(t, "790", "problemsKylePade", problem.Solve)
}
