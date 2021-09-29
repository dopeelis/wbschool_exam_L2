package main

import (
	"sort"
	"strings"
)

func getAnagrams(arr []string) map[string][]string {

	res := make(map[string][]string)

	for _, i := range arr {
		i = strings.ToLower(i)
		res[i] = []string{}
	}

	for k := range res {
		for _, i := range arr {
			i = strings.ToLower(i)
			sortStr := SortString(i)
			sortKey := SortString(k)
			if i == k {
				break
			}
			if sortStr == sortKey {
				delete(res, k)
			}
		}
	}

	for _, i := range arr {
		i = strings.ToLower(i)
		for k := range res {
			sortStr := SortString(i)
			sortKey := SortString(k)
			if i == k {
				break
			}
			if sortStr == sortKey {
				res[k] = append(res[k], i)
			}
		}
	}

	for k, v := range res {
		if len(v) == 0 {
			delete(res, k)
		}
	}

	return res
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
