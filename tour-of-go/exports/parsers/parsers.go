package parsers

import "fmt"

var Goo int = 590

func GetGoods() {
	Bread := "10k"
	Beans := "15k"
	Rice := "12k"
	fmt.Printf("This are the Goods.Bread:%s Beans: %s,Rice:%s ", Bread, Beans, Rice)

}

func GetTotal() {
	alto := getTotal(10, 15, 12)
	fmt.Println("\nThe total of bread, bean and rice is", alto)
	fmt.Println("Total:", totas)
}
