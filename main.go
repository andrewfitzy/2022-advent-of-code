package main

import "fmt"
import "strings"
import "github.com/emirpasic/gods/sets/hashset"

func main() {

	set1 := hashset.New() // empty
	slice1 := strings.Split("abc", "")
	fmt.Println(slice1)

	for _, item := range slice1 {
		set1.Add(item)
	}

	set2 := hashset.New() // empty
	slice2 := strings.Split("bcd", "")
	fmt.Println(slice2)

	for _, item := range slice2 {
		set2.Add(item)
	}

	intersect := set1.Intersection(set2)
	fmt.Println(intersect)
}
