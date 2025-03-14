package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxPathSumInPlace(triangle [][]int) int {
	m := len(triangle)

	for i := m - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += max(triangle[i+1][j], triangle[i+1][j+1])
		}
	}

	return triangle[0][0]
}

func main() {
	data, err := os.ReadFile("hard.json")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var triangle [][]int
	err = json.Unmarshal(data, &triangle)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	result := maxPathSumInPlace(triangle)
	fmt.Println("Maximum Path Sum (In-Place):", result)
}
