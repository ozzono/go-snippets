/*
tamara
amarat
aramat
aaarmt
aaatmr
tmraaa
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(isAnagram("tamara", "tmraaa"))
}

func isAnagram(v1, v2 string) bool {
	if len(v1) != len(v2) {
		return false
	}
	m1 := toMap(v1)
	m2 := toMap(v2)

	for key := range m1 {
		if m1[key] != m2[key] {
			return false
		}
	}

	return true
}

func toMap(input string) map[string]int {
	m := map[string]int{}
	for _, v := range input { //texto
		_, found := m[string(v)]
		if found {
			m[string(v)]++
		} else {
			m[string(v)] = 1
		}
	}
	return m
}

// a:=map[chave]valor
// a,b:=map[chave]valor
