# twoSum Function

This repository contains an implementation of the `twoSum` function in Go. The `twoSum` function is a solution to the Two Sum problem, where given an array of integers `nums` and an integer `target`, it returns indices of the two numbers such that they add up to `target`.

## Function Explanation

### Parameters
- `nums []int`: An array of integers.
- `target int`: The target sum.

### Return Value
- `[]int`: A slice containing the indices of the two numbers in `nums` that add up to `target`. If no such pair is found, it returns `nil`.

### Function Implementation

1. `func twoSum(nums []int, target int) []int {`: This declares the `twoSum` function with two input parameters: `nums`, an array of integers, and `target`, the desired sum. The function returns a slice of integers, which contains the indices of two elements in `nums` that sum up to `target`.

2. `numMap := make(map[int]int)`: This creates a map to store previously seen elements along with their indices in `nums`. This map will be used to match elements with their complements.

3. `for i, num := range nums {`: This loop iterates through each element in `nums`. `i` is the index of the current element, and `num` is its value.

4. `complement := target - num`: This calculates the complement of the current value `num` needed to reach the `target`.

5. `if _, found := numMap[complement]; found { return []int{numMap[complement], i} }`: This checks if the complement of the current value (`complement`) has been seen before in `numMap`. If so, it means there exists a pair of values whose sum equals `target`, and the function returns a slice containing the indices of these two elements. The first index is the previously found complement, and the second index is the index of the current element.

6. `numMap[num] = i`: If the complement is not found in `numMap`, then the current element is added to the map along with its index.

7. `return nil`: If no pair of values is found whose sum equals `target`, the function returns `nil` to indicate that no solution was found.
