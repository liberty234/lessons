package main

import "fmt"

func GetMonthSales() {

	sales := map[string]int{

		"Jan": 1000,
		"Feb": 5000,
	}
	fmt.Printf("value = %v, Type = %T\n", sales["Jan"], sales["Jan"])

	ma1 := make(map[string]int)
	fmt.Printf("value = %v, Type = %T\n", ma1, ma1)

	sales["Mar"] = 10000

	fmt.Printf("value = %v, Type = %T\n", sales["Mar"], sales["Mar"])

	delete(sales, "Mar")
	fmt.Printf("value = %v, Type = %T\n", sales["Mar"], sales["Mar"])

	for k, v := range sales {
		fmt.Println(k, v)

	}

}
