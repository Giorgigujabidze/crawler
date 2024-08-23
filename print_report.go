package main

import (
	"fmt"
	"sort"
)

func printReport(pages map[string]int, basUrl string) {
	fmt.Println("REPORT for", basUrl)
	keys := sortPages(pages)
	for _, k := range keys {
		fmt.Printf("Found %v internal links to %v\n", pages[k], k)
	}
}

func sortPages(pages map[string]int) []string {

	m := pages
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	sort.SliceStable(keys, func(i, j int) bool {

		return m[keys[i]] > m[keys[j]]
	})

	return keys
}
