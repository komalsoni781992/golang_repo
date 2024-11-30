package compositedatatypes

import "fmt"

// TestMaps - tests map functionality
func TestMaps() {
	var nilMap map[string]int
	fmt.Println(nilMap == nil) // maps are not comparable, can only be compared to nil like slices
	//Attempting to read a nil map always returns the zero value for the map’s value type
	fmt.Println(nilMap["Hello"])
	//nilMap["Hello"] = 0 - compile time error to assign a value to nil map

	totalWins := map[string]int{} //empty map literal
	totalWins["Orcas"] = 1        //// automatically grows
	//Note that you cannot use := to assign a value to a map key
	totalWins["Lions"] = 2
	fmt.Println(totalWins["Orcas"])
	fmt.Println(totalWins["Kittens"])
	totalWins["Kittens"]++
	fmt.Println(totalWins["Kittens"])
	totalWins["Lions"] = 3
	fmt.Println(totalWins["Lions"])

	teams := map[string][]string{
		"Orcas":   []string{"Fred", "Ralph", "Bijou"},
		"Lions":   []string{"Sarah", "Peter", "Billie"},
		"Kittens": []string{"Waldo", "Raul", "Ze"},
	}
	for key, value := range teams {
		// Note:- ranging over a map doesn't return values in the same order
		fmt.Println(key, value)
	}
	fmt.Println(len(teams))

	//creates a default size map - Maps created with make still have a length of 0, and they can grow past the initially specified size.
	ages := make(map[int][]string, 10)
	fmt.Println(ages)
	//The key for a map can be any comparable type. This means you cannot use a slice or a map as the key for a map.
}

// CommaOkIdiomTest - tests presence of key in map
func CommaOkIdiomTest() {
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	v, ok := m["hello"]
	fmt.Println(v, ok)
	v, ok = m["world"]
	fmt.Println(v, ok)
	v, ok = m["goodbye"]
	fmt.Println(v, ok)
}

// DeleteFromMap - deletes from map
func DeleteFromMap() {
	m := map[string]int{
		"hello": 5,
		"world": 10,
	}
	delete(m, "hello")
	delete(m, "Bye")
	fmt.Println(m)
}

// MapAsSet - using map as set of integers
func MapAsSet() {
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if intSet[100] {
		fmt.Println("100 is in the set")
	}
}

// Some people prefer to use struct{} for the value when a map is being used to implement a set.
// The advantage is that an empty struct uses zero bytes, while a boolean uses one byte.
// The disadvantage is that using a struct{} makes your code more clumsy.
// You have a less obvious assignment and you need to use the comma ok idiom to check if a value is in the set:
// Unless you have very large sets, it is unlikely that the difference in memory usage is significant enough to outweigh the disadvantages.

// MapAsSetWithStruct - using map witgh struct as value to replicate set functionality
func MapAsSetWithStruct() {
	intSet := map[int]struct{}{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = struct{}{}
	}
	if _, ok := intSet[5]; ok {
		fmt.Println("5 is in the set")
	}
}

// MapOfMaps - function storing map of maps
func MapOfMaps() {
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		}}
	if val, ok := elements["H"]; ok {
		for k, v := range val {
			fmt.Println(k, v)
		}
	}
}
