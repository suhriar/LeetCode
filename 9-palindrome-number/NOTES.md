# isPalindrome Function

This repository contains an implementation of the `isPalindrome` function in Go. The `isPalindrome` function checks whether a given integer `x` is a palindrome or not.

## Function Explanation

### Parameter
- `x int`: The integer to be checked for palindrome.

### Return Value
- `bool`: Returns `true` if the integer `x` is a palindrome, otherwise `false`.

### Function Implementation

1. `func isPalindrome(x int) bool {`: This declares the `isPalindrome` function with one input parameter: `x`, an integer to be checked for palindrome. The function returns a boolean value indicating whether `x` is a palindrome or not.

2. `if x < 0 || (x%10 == 0 && x != 0) { return false }`: This is an initial check to handle cases where `x` is negative or ends with 0 (except for 0 itself), which cannot be palindromes. If this condition is met, the function immediately returns `false`.

3. `reversed := 0`: The `reversed` variable is used to store the reversed number.

4. `original := x`: The `original` variable is used to store the original value of `x` before modification.

5. `for x > 0 {`: This is a loop to reverse the number. This loop will continue as long as `x` is greater than 0.

6. `digit := x % 10`: This extracts the last digit of `x`.

7. `reversed = reversed*10 + digit`: This reverses `x` by multiplying `reversed` by 10 and adding the last digit of `x`.

8. `x /= 10`: This removes the last digit of `x`.

9. `return original == reversed`: After the loop is finished, the function compares the original value of `x` with the reversed value (`reversed`). If they are the same, then `x` is a palindrome, and the function returns `true`; otherwise, `x` is not a palindrome, and the function returns `false`.

This is a simple implementation of the `isPalindrome` function that efficiently reverses an integer mathematically. It does not convert the integer into a string or an array to perform the palindrome check.
