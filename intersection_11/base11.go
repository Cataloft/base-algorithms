package main

import "fmt"

//Реализовать пересечение двух неупорядоченных множеств.

func intersection(a []string, b []string) []string {
	m := map[string]int{}

	for _, v := range a {
		m[v] += 1
	}

	for _, v := range b {
		m[v] += 2
	}

	var intersect []string

	for k, v := range m {
		if v >= 3 {
			intersect = append(intersect, k)
		}
	}

	return intersect
}

func main() {
	a := []string{"cat", "dog", "tree", "sun", "wall"}
	b := []string{"elephant", "frog", "cat", "desk", "dog", "cat"}
	fmt.Println(intersection(a, b))
}
