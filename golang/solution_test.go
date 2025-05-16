package golang

import (
	problem "leetCode/problemsKylePade/problemsKylePade_75"
	"testing"
)

func TestSolution(t *testing.T) {
	TestEach(t, "75", "problemsKylePade", problem.Solve)
}
