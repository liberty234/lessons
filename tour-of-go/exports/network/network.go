package network

import "fmt"

var network string = "MTN network is one of the best network in Nigeria"

func getNetworkName() {
	m := "MTN"
	g := "Glo"
	a := "Airtel"
	fmt.Println("networkName:", m, g, a)

}
