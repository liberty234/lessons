package network

import "fmt"

func GetNetworkType(net1, net2, net3 string) (string, string, string) {
	mtnNetwork := fmt.Sprintf("Network1:%s", net1)
	gloNetwork := fmt.Sprintf("Network2:%s", net2)
	airtelNetwork := fmt.Sprintf("Network3:%s", net3)
	return mtnNetwork, gloNetwork, airtelNetwork
}
