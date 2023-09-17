package state

import "fmt"

func FetchStateName(state1, state2, state3, state4 string) (string, string, string, string) {
	e := fmt.Sprintf("State:%s", state1)
	a := fmt.Sprintf("State:%s", state2)
	d := fmt.Sprintf("State:%s", state3)
	i := fmt.Sprintf("State:%s", state4)
	return e, a, d, i

}
