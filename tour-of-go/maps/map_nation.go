package main

import (
	"fmt"
	"sort"
)

func GetNation() {

	nations := make(map[string]int)

	nations["China"] = 62535537726
	nations["India"] = 54355636626
	nations["United State"] = 54663666272
	nations["Indonatia"] = 7566378338
	nations["Pakistan"] = 73666463387
	nations["Brasil"] = 6443256366263
	nations["Nigeria"] = 32443554762
	nations["Bangladesh"] = 534565263562
	nations["Russia"] = 73666525553
	nations["Mexico"] = 2435563662
	nations["Japan"] = 55366352772

	fmt.Println(nations)

	//sortByKeys(nations)
	//sortByKeysInReverse(nations)
	//return stetus

}

func sortByKeys(records map[string]int64) {
	var keys []string

	for k := range records {
		keys = append(keys, k)

	}

	sort.Strings(keys)

	fmt.Printf("-----------Sorting By keys-------\n")

	for _, key := range keys {
		value := records[key]
		fmt.Printf("%s has a population of %d", key, value)

	}

	fmt.Printf("\n\n")

}
