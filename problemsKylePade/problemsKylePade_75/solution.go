package problem75

import (
	"encoding/json"
	"log"
	"strings"
)

func sortColors(nums []int)  {
    
}

func Solve(inputJsonValues string) interface{} {
	inputValues := strings.Split(inputJsonValues, "\n")
	var nums []int

	if err := json.Unmarshal([]byte(inputValues[0]), &nums); err != nil {
		log.Fatal(err)
	}

	sortColors(nums)
	return nums
}
