package problem3362

import (
	"encoding/json"
	"log"
	"strings"
)

func maxRemoval(nums []int, queries [][]int) int {
    
}

func Solve(inputJsonValues string) interface{} {
	inputValues := strings.Split(inputJsonValues, "\n")
	var nums []int
	var queries [][]int

	if err := json.Unmarshal([]byte(inputValues[0]), &nums); err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal([]byte(inputValues[1]), &queries); err != nil {
		log.Fatal(err)
	}

	return maxRemoval(nums, queries)
}
