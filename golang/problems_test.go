package golang

import (
	"leetCode/problemsKylePade/problemsKylePade_283"
	"testing"
)

func TestSolutions(t *testing.T) {
	TestEach(t, "283", "problemsKylePade", problem283.Solve)
}
