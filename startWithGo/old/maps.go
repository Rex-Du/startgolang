package main

import "fmt"

func main() {
	m := map[string]string{
		"name": "duqing",
		"age":  "30",
	}
	fmt.Println(m)	// map是无序的，hashmap

	m2 := make(map[string]int)
	var m3 map[string]int

	fmt.Println(m2, m3)
	for k, v := range m {
		fmt.Println(k, v)
	}
	delete(m, "name")
	for _, v :=range m{
		fmt.Println(v)
	}
}
