func longestCommonPrefix(strs []string) string {
    // Check if the input slice is empty
    if len(strs) == 0 {
        return ""
    }

    // Initialize the common prefix with the first string in the slice
    prefix := strs[0]

    // Iterate through the remaining strings in the slice
    for i := 1; i < len(strs); i++ {
        // Update the common prefix by comparing characters between strings
        for !strings.HasPrefix(strs[i], prefix) {
            prefix = prefix[:len(prefix)-1]
        }

        // If the common prefix becomes empty, return an empty string
        if prefix == "" {
            return ""
        }
    }

    // Return the final common prefix
    return prefix
}
