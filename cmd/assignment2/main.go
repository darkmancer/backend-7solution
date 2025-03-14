package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter encoded string (using L, R, =): ")
	encoded, _ := reader.ReadString('\n')
	encoded = strings.TrimSpace(encoded)

	n := len(encoded)

	var segments []int
	var cons []byte

	i := 0
	for i < (n + 1) {
		count := 1
		j := i
		for j < n && encoded[j] == '=' {
			count++
			j++
		}
		segments = append(segments, count)
		if j < n {
			cons = append(cons, encoded[j])
		}
		i = j + 1
	}

	m := len(segments)
	v := make([]int, m)

	for i := 0; i < len(cons); i++ {
		if cons[i] == 'R' {
			if v[i+1] < v[i]+1 {
				v[i+1] = v[i] + 1
			}
		}
	}

	for i := len(cons) - 1; i >= 0; i-- {
		if cons[i] == 'L' {
			if v[i] < v[i+1]+1 {
				v[i] = v[i+1] + 1
			}
		}
	}

	var result strings.Builder
	for i, segLen := range segments {
		for j := 0; j < segLen; j++ {
			result.WriteByte(byte('0' + v[i]))
		}
	}

	fmt.Println("Decoded minimal sum digit sequence:", result.String())
}
