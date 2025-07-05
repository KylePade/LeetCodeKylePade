package problem1865

import (
	"encoding/json"
	"log"
	"strings"
)

type FindSumPairs struct {
    
}


func Constructor(nums1 []int, nums2 []int) FindSumPairs {
    
}


func (this *FindSumPairs) Add(index int, val int)  {
    
}


func (this *FindSumPairs) Count(tot int) int {
    
}


/**
 * Your FindSumPairs object will be instantiated and called as such:
 * obj := Constructor(nums1, nums2);
 * obj.Add(index,val);
 * param_2 := obj.Count(tot);
 */

func Solve(inputJsonValues string) interface{} {
	inputValues := strings.Split(inputJsonValues, "\n")
	var operators []string
	var opValues [][]interface{}
	var ans []interface{}
	if err := json.Unmarshal([]byte(inputValues[0]), &operators); err != nil {
		log.Println(err)
		return nil
	}
	if err := json.Unmarshal([]byte(inputValues[1]), &opValues); err != nil {
		log.Println(err)
		return nil
	}
	var arr []int
	if v, ok := opValues[0][0].([]int); ok {
		arr = v
	} else {
		for _, vi := range opValues[0][0].([]interface{}) {
			arr = append(arr, int(vi.(float64)))
		}
	}
	obj := Constructor(arr, arr)
	ans = append(ans, nil)
	for i := 1; i < len(operators); i++ {
		var res interface{}
		switch operators[i] {
		case "add", "Add":
			res = nil
			obj.Add(int(opValues[i][0].(float64)), int(opValues[i][1].(float64)))
		case "count", "Count":
			res = obj.Count(int(opValues[i][0].(float64)))
		default:
			res = nil
		}
		ans = append(ans, res)
	}


	return ans
}
