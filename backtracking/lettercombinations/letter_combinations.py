class Solution:

    numbers = {
        "2": "abc",
        "3": "def",
        "4": "ghi",
        "5": "jkl",
        "6": "mno",
        "7": "pqrs",
        "8": "tuv",
        "9": "wxyz",
    }

    def letterCombinations(self, digits: str) -> List[str]:

        def backtrack(combination: str, curr_digits: str) -> None:
            
            if not curr_digits:
                all_solutions.append(combination)
            else:
                curr_digit = curr_digits[0]
                candidates = self.numbers.get(curr_digit)
                for ind, _ in enumerate(candidates):
                    letter = candidates[ind]
                    backtrack(f"{combination}{candidates[ind]}", curr_digits[1:])
        
        all_solutions = list()
        if digits:
            backtrack("", digits)
        
        return all_solutions