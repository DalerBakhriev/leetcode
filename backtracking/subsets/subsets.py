from typing import List

class Solution:

    def subsets(self, nums: List[int]) -> List[List[int]]:

        def get_subsets(candidates: List[int], curr_index: int):

            if len(candidates) == curr_index:
                all_subsets.append([])
            else:
                get_subsets(candidates, curr_index + 1)
                curr_element = candidates[curr_index]
                new_subsets = []
                for subset in all_subsets:
                    new_subset = subset + [curr_element]
                    new_subsets.append(new_subset)
                all_subsets.extend(new_subsets)
        
        all_subsets = list()
        get_subsets(candidates=nums, curr_index=0)

        return all_subsets



if __name__ == "__main__":
    print(Solution().subsets([1, 2, 3]))