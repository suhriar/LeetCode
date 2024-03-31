# longestCommonPrefix Function

This repository contains an implementation of the `longestCommonPrefix` function in Go. The `longestCommonPrefix` function finds the longest common prefix among a slice of strings.

## Function Explanation

### Parameter
- `strs []string`: A slice of strings.

### Return Value
- `string`: Returns the longest common prefix among all strings in the input slice.

### Function Implementation

1. `func longestCommonPrefix(strs []string) string {`: This line declares the `longestCommonPrefix` function, which takes a slice of strings `strs` as input and returns the longest common prefix among them.

2. `if len(strs) == 0 { return "" }`: This line checks if the input slice is empty. If it is, the function immediately returns an empty string because there is no common prefix among zero strings.

3. `prefix := strs[0]`: This line initializes the common prefix with the first string in the slice.

4. `for i := 1; i < len(strs); i++ { ... }`: This line starts a loop that iterates through the remaining strings in the slice, starting from the second string.

5. `j := 0`: This line initializes a variable `j` to keep track of the index while finding the common prefix between the current prefix and the current string.

6. `for j < len(prefix) && j < len(strs[i]) && prefix[j] == strs[i][j] { j++ }`: This line iteratively compares characters at the same index in the current prefix and the current string until a mismatch is found or one of the strings is exhausted.

7. `prefix = prefix[:j]`: This line updates the common prefix to contain only the characters that are common among all strings so far.

8. `if prefix == "" { break }`: This line checks if the common prefix becomes empty after updating. If it does, it means there is no common prefix among the strings, so the loop can be terminated early.

9. `return prefix`: This line returns the final common prefix found among all strings in the input slice.

This function efficiently finds the longest common prefix among a slice of strings by iteratively comparing characters between strings and updating the common prefix accordingly.
