from typing import List

class Solution:

    def subsetsWithDup(self, nums: List[int]) -> List[List[int]]:

        def get_subsets(candidates: List[int], curr_index: int):

            if len(candidates) == curr_index:
                all_subsets.add(())
            else:
                get_subsets(candidates, curr_index + 1)
                curr_element = candidates[curr_index]
                new_subsets = list()
                for subset in all_subsets:
                    new_subset = list(subset) + [curr_element]
                    new_subsets.append(new_subset)
                
                for new_subset in new_subsets:
                    all_subsets.add(tuple(sorted(new_subset)))
        
        all_subsets = set()
        get_subsets(candidates=nums, curr_index=0)

        return [list(subset) for subset in all_subsets]


if __name__ == "__main__":
    print(Solution().subsetsWithDup([1, 2, 2]))
    print(Solution().subsetsWithDup([4, 4, 4, 1, 4]))