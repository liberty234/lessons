package parsers

import "fmt"

func Goods() {
	Bread := "10k"
	Beans := "15k"
	Rice := "12k"
	fmt.Printf("This are the Goods.Bread:%s Beans: %s,Rice:%s ", Bread, Beans, Rice)

}

func Total() {
	alto := total(10, 15, 12)
	fmt.Println("\nThe total of bread, bean and rice is", alto)
}
