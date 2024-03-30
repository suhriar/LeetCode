func romanToInt(s string) int {
    // Define a map to store the integer values corresponding to each Roman numeral
    romanValues := map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

    // Initialize the result variable
    result := 0

    // Loop through the input string
    for i := 0; i < len(s); i++ {
        // Get the integer value of the current Roman numeral
        value := romanValues[s[i]]

        // Check if the current numeral is the last one or its value is less than the next numeral
        if i < len(s)-1 && romanValues[s[i+1]] > value {
            // Subtract the value of the current numeral from the result
            result -= value
        } else {
            // Add the value of the current numeral to the result
            result += value
        }
    }

    // Return the final result
    return result
}
