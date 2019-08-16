// Go maps, equivalent to well-known hash table
package main

import "fmt"

func main() {
	aMap := make(map[string]int)
	aMap["key1"] = 11
	aMap["key2"] = 22
	fmt.Println("aMap:", aMap)

	anotherMap := map[string]int{
		"key1": 11,
		"key2": 22,
	}
	fmt.Println("anotherMap:", anotherMap)

	// deleting the map element based on key using delete() function
	delete(anotherMap, "key1")

	// the delete function will not return an error or warning message although we try to delete a non existing element
	delete(anotherMap, "key1")

	fmt.Println("anotherMap:", anotherMap)

	// a technique to determine whether a given key is in the map or not
	_, ok := anotherMap["doesItExist"]
	if ok {
		fmt.Println("Exist!")
	} else {
		fmt.Println("Does NOT exist")
	}

	// using range keyword on a map
	for key, value := range aMap {
		fmt.Println(key, value)
	}
	// we cannot and should not make any assumption about the order of map pairs, because it is random

	thisMap := map[string]int{} // initiating an empty map

	//thisMap = nil
	//when you uncomment the line above, the code will be fail, we cant insert data into nil map,
	fmt.Println(thisMap)
	thisMap["test"] = 1
	fmt.Println(thisMap)

}
