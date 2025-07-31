package golang

import (
	problem "leetCode/problemsKylePade/problemsKylePade_118"
	"testing"
)

func TestSolution(t *testing.T) {
	TestEach(t, "118", "problemsKylePade", problem.Solve)
}
