func isPalindrome(x int) bool {
    // Handle negative numbers and numbers ending with 0
    if x < 0 || (x%10 == 0 && x != 0) {
        return false
    }

    reversed := 0
    original := x

    // Reversing the number
    for x > 0 {
        digit := x % 10
        reversed = reversed*10 + digit
        x /= 10
    }

    // Checking if the reversed number equals the original number
    return original == reversed
}
