from typing import List, Tuple
import copy

class Solution:

    @staticmethod
    def math_round(num: int):
        num = int(num + (0.5 if num > 0 else -0.5))
        return num

    def _factorial(self, n: int) -> int:

        if n <= 1:
            return 1
        return self._factorial(n-1)*n

    def _calculate_kth_perm_first_element(self, candidates: List[int], n: int, k: int) -> Tuple[str, int]:
        
        num_times_first_element = self._factorial(n - 1)

    
        start_element_index = self.math_round(k / num_times_first_element) - 1
        new_index = (k - 1) % num_times_first_element


        return str(candidates[start_element_index]), new_index

    def getPermutation(self, n: int, k: int) -> str:
        
        def backtrack(candidates: List[int], combination: str):

            if combination and combination[0] != first_element:
                return

            if len(solutions) == k:
                return

            if len(combination) == n:
                solutions.append(combination)
                return

            for ind, candidate in enumerate(candidates):
                combination = "".join((combination, str(candidate)))
                candidates_copy = copy.copy(candidates)
                candidates_copy.pop(ind)
                backtrack(candidates_copy, combination)
                combination = combination[:-1]
        
        solutions = list()
        candidates = list(range(1, n + 1))
        first_element, new_index = self._calculate_kth_perm_first_element(candidates, n, k)
        backtrack(candidates=candidates, combination="")
        print(solutions, new_index, first_element)

        return solutions[new_index]


if __name__ == "__main__":
    print(Solution().getPermutation(n=3, k=3))
    # print(Solution().getPermutation(n=4, k=9))
    # print(Solution().getPermutation(n=2, k=1))
    # print(Solution().getPermutation(n=9, k=161191))
    # print(Solution().getPermutation(n=1, k=1))
    # print(Solution().getPermutation(n=2, k=2))
    # print(Solution().getPermutation(n=3, k=2))