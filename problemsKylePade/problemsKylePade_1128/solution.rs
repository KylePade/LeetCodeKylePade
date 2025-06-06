#![allow(non_snake_case)]
use serde_json::{json, Value};

pub struct Solution;

impl Solution {
    pub fn num_equiv_domino_pairs(dominoes: Vec<Vec<i32>>) -> i32 {
        
    }
}

#[cfg(feature = "solution_1128")]
pub fn solve(input_string: String) -> Value {
	let input_values: Vec<String> = input_string.split('\n').map(|x| x.to_string()).collect();
	let dominoes: Vec<Vec<i32>> = serde_json::from_str(&input_values[0]).expect("Failed to parse input");
	json!(Solution::num_equiv_domino_pairs(dominoes))
}
