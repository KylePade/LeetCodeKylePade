package problem2894

import (
	"encoding/json"
	"log"
	"strings"
)

func differenceOfSums(n int, m int) int {
    
}

func Solve(inputJsonValues string) interface{} {
	inputValues := strings.Split(inputJsonValues, "\n")
	var n int
	var m int

	if err := json.Unmarshal([]byte(inputValues[0]), &n); err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal([]byte(inputValues[1]), &m); err != nil {
		log.Fatal(err)
	}

	return differenceOfSums(n, m)
}
