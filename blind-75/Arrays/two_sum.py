"""Given an array of integers nums and an integer target, 
return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, 
and you may not use the same element twice."""

# nums = [2,7,11,15], target = 9 => [0,1]

def solve(nums: list, target: int) -> list:
    checked = {}
    

    for i in range(len(nums)):
        diff = target - nums[i]
        if diff in checked:
            return [i, checked[diff]]
        checked[nums[i]] = i


nums = [2,7,11,15]
target = 9

print(solve(nums,target))