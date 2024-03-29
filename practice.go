/*
This is not intended to run out of the box
The individual funcs are tested in Leetcode
*/

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

// Begin: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/
type Key struct {
    X, Y int
}

func maxProfit(prices []int) int {
    cache := make(map[Key]int)
    return helper(prices, 0, 0, cache)
}

func helper(prices []int, i int, op int, cache map[Key]int) int {
    const BUY = 0
    const SELL = 1
    if i>= len(prices) {
        cache[Key{i, op}] = 0
        return 0
    }
    if _, ok := cache[Key{i, op}]; ok {
        return cache[Key{i, op}]
    }
    var p int
    if op == BUY{
        p = max(helper(prices, i+1, SELL, cache)- prices[i], helper(prices, i+1, BUY, cache))
    } else if op == SELL{
        p = max(helper(prices, i+2, BUY, cache) + prices[i],
        helper(prices, i+1, SELL, cache))
    }
    cache[Key{i, op}] = p 
    return cache[Key{i, op}]
}
// END: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/

// BEGIN: https://leetcode.com/problems/daily-temperatures
func dailyTemperatures(temperatures []int) []int {
    stack := make([][2]int, 0)
    result := make([]int, len(temperatures)) //default 0

    for i, temp := range temperatures {
        for len(stack) > 0 && stack[len(stack)-1][1] < temp {
            prev := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            result[prev[0]] = i - prev[0]
        }
        stack = append(stack, [2]int{i, temp})

    }
    return result

}

// END: https://leetcode.com/problems/daily-temperatures

// BEGIN: https://leetcode.com/problems/max-area-of-island/description/
func dfs(row, col, ROWS, COLS int, visited[][]bool, grid[][]int) int {
    if row < 0  || row >= ROWS || col < 0 || col >= COLS || visited[row][col] || grid[row][col] == 0 {
        return 0
    }
    visited[row][col] = true
    return 1 + dfs(row+1, col, ROWS, COLS, visited, grid) + dfs(row-1, col, ROWS, COLS, visited, grid) + dfs(row, col+1, ROWS, COLS, visited, grid) + dfs(row, col-1, ROWS, COLS, visited, grid)

}

func maxAreaOfIsland(grid [][]int) int {
    ROWS := len(grid)
    COLS := len(grid[0])
    visited := make([][]bool, ROWS)
    for i := range visited {
        visited[i] = make([]bool, COLS)
    }
    area := 0

    for row := 0; row < ROWS; row++ {
        for col := 0; col < COLS; col++ {
            curr_area := dfs(row, col, ROWS, COLS, visited, grid)
            if curr_area > area {
                area = curr_area
            }
        }
    }
    return area
} 
// END: https://leetcode.com/problems/max-area-of-island/description/