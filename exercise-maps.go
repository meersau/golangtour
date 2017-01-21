package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	m := make(map[string]int)
	for _, w := range words {
		elem, ok := m[w]
		if ok {
		  m[w] = elem + 1
		 continue
		}
		m[w] = 1
	}
	return m
}


