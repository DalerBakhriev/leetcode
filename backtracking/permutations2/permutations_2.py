from typing import List
import copy

class Solution:

    def permuteUnique(self, nums: List[int]) -> List[List[int]]:
        
        def backtrack(candidates: List[int], solution: List[int]):

            if not candidates:
                solutions.add(tuple(copy.copy(solution)))
                return

            for ind, candidate in enumerate(candidates):
                solution.append(candidate)
                candidates_copy = copy.copy(candidates)
                candidates_copy.pop(ind)
                backtrack(candidates_copy, solution)
                solution.pop(-1)
        
        solutions = set()
        if nums:
            backtrack(candidates=nums, solution=[])
        
        final_solutions = [list(solution) for solution in solutions]

        return final_solutions


if __name__ == "__main__":
    print(Solution().permuteUnique([1, 1, 2]))