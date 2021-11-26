package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func choose(places, entries int) []int {
	places = min(places, entries)
	result := make([]int, 0, places)
	for i := 0; i < places; {
		r := rand.Intn(entries)
		if !contains(r, result) {
			result = append(result, r)
			i++
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func contains(v int, s []int) bool {
	for i := range s {
		if s[i] == v {
			return true
		}
	}
	return false
}
