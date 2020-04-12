from typing import List
import copy

class Solution:

    def permute(self, nums: List[int]) -> List[List[int]]:
        
        def backtrack(candidates: List[int], solution: List[int]):

            if not candidates:
                solutions.append(copy.copy(solution))
                return

            for ind, candidate in enumerate(candidates):
                solution.append(candidate)
                candidates_copy = copy.copy(candidates)
                candidates_copy.pop(ind)
                backtrack(candidates_copy, solution)
                solution.pop(-1)
        
        solutions = list()
        if nums:
            backtrack(candidates=nums, solution=[])

        return solutions


if __name__ == "__main__":
    print(Solution().permute([1, 2, 3]))
    print(Solution().permute([5, 4, 2, 6]))