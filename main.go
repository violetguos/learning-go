package main

import (
    "strings"
    "fmt"
)

func letterCombinations(digits string) []string {
    numToChar := map[string]string{
        "2": "abc",
        "3": "def",
        "4": "ghi",
        "5": "jkl",
        "6": "mno",
        "7": "pqrs",
        "8": "tuv",
        "9": "wxyz",
    }
    result := []string{}
    if len(digits) == 0 {
        return result
    }
    for i := range digits {
        if len(result) == 0 {
            fmt.Println(digits[i])
            chars := strings.Split(numToChar[digits[i]], "")
            for _, c := range chars {
                result = append(result, c)
            }
        } else {
            curr_result := []string{}
            for j := range result {
                chars := strings.Split(numToChar[digits[i]], "")
                for _, c := range chars {
                    curr_result = append( append(result[j], c))
                }
            }
            result = curr_result
        }
    }
    return result

}

func main() {

	fmt.Println(letterCombinations("23"))
}