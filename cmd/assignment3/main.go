package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const baconIpsumURL = "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"

func countMeats(text string) map[string]int {
	re := regexp.MustCompile(`[a-zA-Z0-9-]+`)
	words := re.FindAllString(text, -1)
	counts := make(map[string]int)
	for _, word := range words {
		lowerWord := strings.ToLower(word)
		counts[lowerWord]++
	}
	return counts
}

func beefSummaryHandler(w http.ResponseWriter, r *http.Request) {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(baconIpsumURL)
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read data", http.StatusInternalServerError)
		return
	}
	text := string(body)

	counts := countMeats(text)

	result := map[string]interface{}{
		"beef": counts,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/beef/summary", beefSummaryHandler)
	log.Println("Server is listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
