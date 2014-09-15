package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3}
	fmt.Println(len(arr))

	sli := arr[0:2]

	fmt.Println("sli:", sli)
	fmt.Println("sli len:", len(sli))

	sli1 := []int{1, 2, 3}

	sli2 := append(sli1, 5, 6)

	fmt.Println("sli2: ", sli2)

	sli3 := make([]int, 3)

	copy(sli3, sli1)

	fmt.Println("sli3 after copy: ", sli3)

	copy(sli2, sli1)

	fmt.Println("sli2 after copy", sli2)

	fmt.Println("sli2[2]", sli2[2])

	map1 := make(map[string]int)

	map1["key"] = 1

	fmt.Println("map1[key]: ", map1["key"])
	fmt.Println("map1[key1]: ", map1["key1"])

	map2 := make(map[string]string)
	map2["key"] = "value"

	fmt.Println("map2[key]", map2["key"])
	fmt.Println("map2[key1]", map2["key1"])

	value, ok := map1["key1"]

	fmt.Println(value, ok)

	if value, ok := map2["key3"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("map2 does not have key3")
	}

}