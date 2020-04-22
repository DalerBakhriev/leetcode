from typing import List
import logging
import copy

class Solution:

    def combinationSum2(self, candidates: List[int], target: int) -> List[List[int]]:

        def backtrack(candidates: List[int], solution: List[int], start: int, remains: int) -> None:

            if remains < 0:
                return
            if remains == 0:
                solution_copy = tuple(sorted(copy.copy(solution)))
                solutions.add(solution_copy)
                return
            
            for ind in range(start, len(candidates)):
                solution.append(candidates[ind])
                candidates_rm = copy.copy(candidates)
                _ = candidates_rm.pop(ind)
                backtrack(candidates_rm, solution, ind, remains-candidates[ind])
                solution.pop(-1)

        solutions = set()
        if target:
            backtrack(candidates=candidates, solution=[], remains=target, start=0)
        
        final_solutions = [list(solution) for solution in solutions]
        return final_solutions



if __name__ == "__main__":
    logging.getLogger().setLevel(logging.INFO)
    print(Solution().combinationSum2(candidates=[10, 1, 2, 7, 6, 1, 5], target=8))
    print(Solution().combinationSum2(candidates=[2, 5, 2, 1, 2], target=5))
