//go:build ignore
#include "cpp/common/Solution.h"


using namespace std;
using json = nlohmann::json;

class Solution {
public:
    string smallestEquivalentString(string s1, string s2, string baseStr) {
        
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
	string s1 = json::parse(inputArray.at(0));
	string s2 = json::parse(inputArray.at(1));
	string baseStr = json::parse(inputArray.at(2));
	return solution.smallestEquivalentString(s1, s2, baseStr);
}
