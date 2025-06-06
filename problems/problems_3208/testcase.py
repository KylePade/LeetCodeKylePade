from collections import namedtuple
import testcase

case = namedtuple("Testcase", ["Input", "Output"])


class Testcase(testcase.Testcase):
	def __init__(self):
		self.testcases = []
		self.testcases.append(case(Input=[[0, 1, 0, 1, 0], 3], Output=3))
		self.testcases.append(case(Input=[[0, 1, 0, 0, 1, 0, 1], 6], Output=2))
		self.testcases.append(case(Input=[[1, 1, 0, 1], 4], Output=0))
		self.testcases.append(case(Input=[[0,1,1],3], Output=1))

	def get_testcases(self):
		return self.testcases
