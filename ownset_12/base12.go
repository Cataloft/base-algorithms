package main

import "fmt"

//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

func ownSet(a []string) map[string][]string {
	m := map[string]int{}
	set := map[string][]string{}

	for _, v := range a {
		m[v] += 1
	}

	for k, v := range m {
		if v > 1 {
			for i := 0; i < v; i++ {
				set[k] = append(set[k], k)
			}
			continue
		}
		set[k] = append(set[k], k)
	}

	return set
}

func main() {
	a := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(ownSet(a))
}
