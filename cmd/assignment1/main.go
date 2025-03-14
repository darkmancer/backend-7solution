package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
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

func loadTriangleFromFile(filename string) ([][]int, error) {
	workingDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get current directory: %w", err)
	}

	filePath := filepath.Join(workingDir, filename)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		filePath = filepath.Join(workingDir, "cmd", "assignment1", filename)
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %w", filePath, err)
	}

	var triangle [][]int
	err = json.Unmarshal(data, &triangle)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return triangle, nil
}

func main() {
	triangle, err := loadTriangleFromFile("hard.json")
	if err != nil {
		log.Fatalf("Failed to load triangle data: %v", err)
	}

	result := maxPathSumInPlace(triangle)
	fmt.Println("Maximum Path Sum :", result)
}
