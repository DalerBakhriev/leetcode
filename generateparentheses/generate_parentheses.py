from typing import List

class Solution:

    def _insert_inside(self, string: str, left_index: int) -> str:

        left_part = string[:left_index + 1]
        right_part = string[left_index + 1:]

        return f"{left_part}(){right_part}"

    def generateParenthesis(self, n: int) -> List[str]:

        solutions = set()
        if not n:
            solutions.add("")
        else:
            previous_solution = self.generateParenthesis(n-1)
            for string in previous_solution:
                for ind, _ in enumerate(string):
                    if string[ind] == "(":
                        solution = self._insert_inside(string, ind)
                        solutions.add(solution)
                solutions.add("()" + string)
        
        return list(solutions)


if __name__ == "__main__":
    print(Solution().generateParenthesis(3))