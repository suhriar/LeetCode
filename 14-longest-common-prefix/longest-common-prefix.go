func longestCommonPrefix(strs []string) string {
    // Check if the input slice is empty
    if len(strs) == 0 {
        return ""
    }

    // Initialize the common prefix with the first string in the slice
    prefix := strs[0]

    // Iterate through the remaining strings in the slice
    for i := 1; i < len(strs); i++ {
        // Find the common prefix between the current prefix and the current string
        j := 0
        for j < len(prefix) && j < len(strs[i]) && prefix[j] == strs[i][j] {
            j++
        }
        prefix = prefix[:j]

        // If the common prefix becomes empty, no need to check further
        if prefix == "" {
            break
        }
    }

    // Return the final common prefix
    return prefix
}
