#![allow(non_snake_case)]
use serde_json::{json, Value};

pub struct Solution;

impl Solution {
    pub fn difference_of_sums(n: i32, m: i32) -> i32 {
        
    }
}

#[cfg(feature = "solution_2894")]
pub fn solve(input_string: String) -> Value {
	let input_values: Vec<String> = input_string.split('\n').map(|x| x.to_string()).collect();
	let n: i32 = serde_json::from_str(&input_values[0]).expect("Failed to parse input");
	let m: i32 = serde_json::from_str(&input_values[1]).expect("Failed to parse input");
	json!(Solution::difference_of_sums(n, m))
}
