func partitionLabels(s string) []int {
    // https://leetcode.com/problems/partition-labels/description/
	// beats 100% 
	charPositionMap := make(map[rune]int)
    for i, char := range s {
        charPositionMap[char] = i
    }
    leftPtr := 0
    rightPtr := 0
    partitions := []int{}
    for i, char := range s {
        rightMost := charPositionMap[char]
        if rightPtr < rightMost {
            rightPtr = rightMost
        }
        if i == rightPtr {
            partitions = append(partitions, rightPtr - leftPtr + 1)
            leftPtr = rightPtr + 1
        }
    }
    return partitions
}