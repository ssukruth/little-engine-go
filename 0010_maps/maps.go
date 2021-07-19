package main

import "fmt"

func main() {
	// Maps store key:value pairs
	// Constant get, set & deletion time as they're backed by hash tables
	// Keys & Values are statically typed. All keys must be of same type.
	// All values must be of the same type.
	// All keys must be unique and must be of a comparable type. float
	// types have comparable issues and avoid using them as keys.
	// Maps cannot be compared to each other and can only be compared to nil.
	// Maps are unordered data structures
	var myMap map[string]string // map with string keys and values

	// Default value of map is nil and default number of keys is 0
	fmt.Printf("myMap is of type %T and value %#v\n", myMap, myMap) // nil value
	fmt.Println("myMap has", len(myMap), "keys")

	// If an element doesn't exist in the map, it returns 0 value of the value type
	fmt.Printf("myMap[\"nonexistentkey\"] is [%q]\n", myMap["nonexistentkey"])

	// Cannot use slice as key since slices are not comparable
	// The following line is invalid
	// var sliceMap map[[]int]int
	var arrMap map[[2]int]int
	fmt.Printf("arrMap is of type %T\n", arrMap)

	// Maps need to be initialized using make function or a map literal
	// before assignment
	var newMap map[int]int
	// Following line is invalid
	// newMap[2] = 4

	// initialize map using make
	newMap = make(map[int]int)
	newMap[2] = 4
	newMap[3] = 9
	fmt.Printf("newMap's value is %#v\n", newMap)

	// initialize map with empty map
	newMap1 := map[int]int{}
	newMap1[2] = 4
	newMap1[3] = 6
	fmt.Printf("newMap1's value is %#v\n", newMap1)

	// initialize with map literal
	newMap2 := map[int]int{
		1: 1,
		2: 2,
		3: 3,
	}
	fmt.Printf("newMap2's value is %#v\n", newMap2)

	// As mentioned before, go returns 0 value of the value type
	// if a key doesn't exist. inorder to distinguish between an actual 0
	// value and a missing key, we use the "comma, ok" idiom
	newMap2[0] = 0
	v1, ok := newMap2[0]
	if !ok {
		fmt.Println("0 is not present in newMap2")
	} else {
		fmt.Println("newMap2[0] is", v1)
	}
	v2, ok := newMap2[4]
	if !ok {
		fmt.Println("4 is not present in newMap2")
	} else {
		fmt.Println("newMap2[4] is", v2)
	}
	// Iterating over maps using range
	for k, v := range newMap2 {
		fmt.Printf("key: %#v, val: %#v\n", k, v)
	}

	// To delete a pair use delete operation
	delete(newMap2, 0)
	fmt.Printf("newMap2's value is %#v\n", newMap2)
	// delete can work with non-existent keys too
	delete(newMap2, -1)

	// Maps can only be compared to nil. However if you need to
	// compare 2 maps, we can compare their string representations
	map1 := map[string]int{"hi": 2, "bye": 3}
	map2 := map[string]int{"bye": 3, "hi": 2}
	s1 := fmt.Sprintf("%s", map1)
	s2 := fmt.Sprintf("%s", map2)
	fmt.Println(s1)
	fmt.Println(s2)
	if s1 == s2 {
		fmt.Println("maps are equal")
	} else {
		fmt.Println("maps are not equal")
	}

	// When declaring a map variable go creates a pointer to map header
	// in memory. The map references this internal data structure, map
	// header. The map contains only the memory address if the map header.
	// The key:value pairs are not stored directly into the map.
	// They're stored in memory at the address referenced by map header.
	// Therefore when you copy, the internal data structure is not copied
	// but just referenced
	m1 := map[string]int{"math": 95, "science": 96}
	m2 := m1
	m2["english"] = 90
	fmt.Println(m1) //map[english:90 math:95 science:96]
	fmt.Println(m2) //map[english:90 math:95 science:96]

	// To create a clone, we initialize new map and then iterate over
	// the first map to copy each element into the second
	m3 := make(map[string]int)
	for k, v := range m1 {
		m3[k] = v
	}
	m3["social"] = 91
	fmt.Println(m1) //map[english:90 math:95 science:96]
	fmt.Println(m3) //map[english:90 math:95 science:96 social:91]

}
