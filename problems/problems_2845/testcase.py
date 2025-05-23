from collections import namedtuple
import testcase

case = namedtuple("Testcase", ["Input", "Output"])


class Testcase(testcase.Testcase):
	def __init__(self):
		self.testcases = []
		self.testcases.append(case(Input=[[3, 2, 4], 2, 1], Output=3))
		self.testcases.append(case(Input=[[3, 1, 9, 6], 3, 0], Output=2))

	def get_testcases(self):
		return self.testcases
