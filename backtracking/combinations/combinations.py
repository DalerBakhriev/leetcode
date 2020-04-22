from typing import List
import copy
import itertools

class Solution:

    def combine(self, n: int, k: int) -> List[List[int]]:
        
        def backtrack(candidates: List[int], combination: List[int]):
            
            if len(combination) == k:
                combinations.add(tuple(sorted(combination)))
            
            for ind, candidate in enumerate(candidates):
                combination.append(candidate)
                candidate_copy = copy.copy(candidates)
                candidate_copy.pop(ind)
                backtrack(candidate_copy, combination)
                combination.pop(-1)

        candidates = list(range(1, n + 1))
        combinations = set()
        backtrack(candidates, [])
        return [list(combination) for combination in combinations]


if __name__ == "__main__":
    print(Solution().combine(4, 2))
    print(Solution().combine(10, 5))
    print(Solution().combine(1, 1))