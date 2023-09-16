package network

import "fmt"

func GetNetworkType(net1, net2, net3 string) (string, string, string) {
	m := fmt.Sprintf("Network1:%s", net1)
	g := fmt.Sprintf("Network2:%s", net2)
	a := fmt.Sprintf("Network3:%s", net3)
	return m, g, a
}
