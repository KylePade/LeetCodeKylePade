package problem2873

import (
	"encoding/json"
	"log"
	"strings"
)

func maximumTripletValue(nums []int) int64 {
    
}

func Solve(inputJsonValues string) interface{} {
	inputValues := strings.Split(inputJsonValues, "\n")
	var nums []int

	if err := json.Unmarshal([]byte(inputValues[0]), &nums); err != nil {
		log.Fatal(err)
	}

	return maximumTripletValue(nums)
}
