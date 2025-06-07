package golang

import (
	problem "leetCode/problemsKylePade/problemsKylePade_386"
	"testing"
)

func TestSolution(t *testing.T) {
	TestEach(t, "386", "problemsKylePade", problem.Solve)
}
