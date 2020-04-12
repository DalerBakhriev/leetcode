from typing import List, Optional
import logging
import copy

class Solution:

    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:

        def backtrack(candidates: List[int], solution: List[int], start: int, remains: int) -> None:

            if remains < 0:
                return
            if remains == 0:
                solutions.append(copy.deepcopy(solution))
                return
            
            for ind in range(start, len(candidates)):
                solution.append(candidates[ind])
                backtrack(candidates, solution, ind, remains-candidates[ind])
                solution.pop(-1)

        solutions = list()
        if target:
            backtrack(candidates=candidates, solution=[], remains=target, start=0)
        
        return solutions



if __name__ == "__main__":
    logging.getLogger().setLevel(logging.INFO)
    print(Solution().combinationSum(candidates=[2, 3, 6, 7], target=7))
    print(Solution().combinationSum(candidates=[2, 3, 5], target=8))
