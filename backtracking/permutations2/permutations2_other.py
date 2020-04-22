import collections
from typing import DefaultDict, List

class Solution:
    def permuteUniqueHelper(self, freqMap: DefaultDict[int, int], prefix: List[int], n: int, permutations: List[int]):
        if len(prefix) == n:
            permutations.append(prefix)
            return
        
        for num in freqMap:
            if freqMap[num] > 0:
                freqMap[num] -= 1
                self.permuteUniqueHelper(freqMap, prefix + [num], n, permutations)
                freqMap[num] += 1
    
    def permuteUnique(self, nums: List[int]) -> List[List[int]]:
        freqMap = collections.defaultdict(int)
        for num in nums:
            freqMap[num] += 1
            
        permutations = []
        self.permuteUniqueHelper(freqMap, [], len(nums), permutations)
        return permutations


if __name__ == "__main__":
    print(Solution().permuteUnique([1, 1, 2]))