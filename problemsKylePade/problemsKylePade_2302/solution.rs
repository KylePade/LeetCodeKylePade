#![allow(non_snake_case)]
use serde_json::{json, Value};

pub struct Solution;

impl Solution {
    pub fn count_subarrays(nums: Vec<i32>, k: i64) -> i64 {
        
    }
}

#[cfg(feature = "solution_2302")]
pub fn solve(input_string: String) -> Value {
	let input_values: Vec<String> = input_string.split('\n').map(|x| x.to_string()).collect();
	let nums: Vec<i32> = serde_json::from_str(&input_values[0]).expect("Failed to parse input");
	let k: i64 = serde_json::from_str(&input_values[1]).expect("Failed to parse input");
	json!(Solution::count_subarrays(nums, k))
}
