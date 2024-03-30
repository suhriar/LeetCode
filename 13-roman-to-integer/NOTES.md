# romanToInt Function

This repository contains an implementation of the `romanToInt` function in Go. The `romanToInt` function converts a Roman numeral string into its corresponding integer value.

## Function Explanation

### Parameter
- `s string`: The Roman numeral string to be converted.

### Return Value
- `int`: Returns the integer value of the given Roman numeral string.

### Function Implementation

1. `func romanToInt(s string) int {`: This line declares the `romanToInt` function which takes a string `s` representing a Roman numeral as input and returns its corresponding integer value.

2. `romanValues := map[byte]int{ ... }`: This line initializes a map named `romanValues` where the keys are bytes representing Roman numerals ('I', 'V', 'X', etc.) and the values are their corresponding integer values (1, 5, 10, etc.).

3. `result := 0`: This line initializes a variable `result` to store the final integer value of the Roman numeral.

4. `for i := 0; i < len(s); i++ { ... }`: This line starts a loop that iterates over each character of the input string `s`.

5. `value := romanValues[s[i]]`: This line retrieves the integer value of the current Roman numeral (`s[i]`) from the `romanValues` map.

6. `if i < len(s)-1 && romanValues[s[i+1]] > value { ... } else { ... }`: This line checks if the current numeral is the last one or if its value is less than the value of the next numeral. If true, it subtracts the current numeral's value from the result; otherwise, it adds the value to the result.

7. `result -= value`: This line subtracts the value of the current numeral from the result if it is preceded by a larger numeral.

8. `result += value`: This line adds the value of the current numeral to the result if it is not preceded by a larger numeral or if it is the last numeral.

9. `return result`: This line returns the final integer value of the Roman numeral.

This README.md file provides a detailed explanation of the `romanToInt` function, its parameter, return value, and implementation in Go.
