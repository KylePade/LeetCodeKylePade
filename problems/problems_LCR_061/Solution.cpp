//go:build ignore
#include "cpp/common/Solution.h"


using namespace std;
using json = nlohmann::json;

class Solution {
public:
    vector<vector<int>> kSmallestPairs(vector<int>& nums1, vector<int>& nums2, int k) {

    }
};

json leetcode::qubh::Solve(string input_json_values) {
	vector<string> inputArray;
	size_t pos = input_json_values.find('\n');
	while (pos != string::npos) {
		inputArray.push_back(input_json_values.substr(0, pos));
		input_json_values = input_json_values.substr(pos + 1);
		pos = input_json_values.find('\n');
	}
	inputArray.push_back(input_json_values);

	Solution solution;
	vector<int> nums1 = json::parse(inputArray.at(0));
	vector<int> nums2 = json::parse(inputArray.at(1));
	int k = json::parse(inputArray.at(2));
	return solution.kSmallestPairs(nums1, nums2, k);
}
