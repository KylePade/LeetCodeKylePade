package problem904

import (
	"encoding/json"
	"log"
	"strings"
)

func totalFruit(fruits []int) int {
    
}

func Solve(inputJsonValues string) interface{} {
	inputValues := strings.Split(inputJsonValues, "\n")
	var fruits []int

	if err := json.Unmarshal([]byte(inputValues[0]), &fruits); err != nil {
		log.Fatal(err)
	}

	return totalFruit(fruits)
}
